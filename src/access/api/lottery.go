package api

import (
	"context"
	"encoding/json"
	"errors"
	"lottery_backend/src/xorm"

	"github.com/gin-gonic/gin"

	"lottery_backend/src/access/model"
	cfg "lottery_backend/src/config"
	xredis "lottery_backend/src/redis"
	"lottery_backend/src/utils"
	"lottery_backend/src/xlog"
)

// Lottery: lottery
func Lottery(c *gin.Context, ctx context.Context) (int, interface{}) {
	sessionId := utils.GetSessionIdFromContext(ctx)
	res := model.LotteryResponse{
		Action:  "LotteryResponse",
		RetCode: 0,
	}
	var req model.LotteryRequest
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		xlog.Error(sessionId, "NewDecoder error", xlog.Fields{
			"request": req,
		})
		res.Message = err.Error()
		res.RetCode = ERR_PARSE_PARAMS_ERROR
		return SUCCESS_ON_ACTION_RETCODE, res
	}
	if req.PhoneNumber == 0 {
		xlog.Error(sessionId, "Lottery param invalid", xlog.Fields{
			"request": req,
		})
		res.Message = "Lottery param invalid"
		res.RetCode = ERR_PARSE_PARAMS_ERROR
		return SUCCESS_ON_ACTION_RETCODE, res
	}

	// get user info
	user, err := xorm.GetUserInfoByPhoneNum(req.PhoneNumber)
	if err != nil {
		xlog.Error(sessionId, "Get User Info ERROR", xlog.Fields{
			"request": req,
			"err":     err,
		})
		res.Message = err.Error()
		res.RetCode = ERR_XORM_ERROR
		return SUCCESS_ON_ACTION_RETCODE, res
	}

	// 1. permission check
	permit, err := lotteryPermitted(ctx, user.Id)
	if err != nil {
		xlog.Error(sessionId, "Lottery Permitted ERROR", xlog.Fields{
			"request": req,
			"err":     err,
		})
		res.Message = err.Error()
		res.RetCode = ERR_XORM_ERROR
		return SUCCESS_ON_ACTION_RETCODE, res
	}

	if !permit {
		xlog.Debug(sessionId, "Lottery Unpermitted", xlog.Fields{})
		res.Message = "once a day"
		return SUCCESS_ON_ACTION_RETCODE, res
	}

	// 2. lottery
	err = lottery(ctx, user.Id, req.PhoneNumber, &res)
	if err != nil {
		xlog.Error(sessionId, "Lottery ERROR", xlog.Fields{
			"UserId":      user.Id,
			"PhoneNumber": req.PhoneNumber,
			"err":         err,
		})
		res.Message = err.Error()
		res.RetCode = ERR_LOTTERY_ERROR
		return SUCCESS_ON_ACTION_RETCODE, res
	}

	return SUCCESS_ON_ACTION_RETCODE, res
}

// lottery
func lottery(ctx context.Context, userId string,
	phoneNumber uint64, lotteryRes *model.LotteryResponse) error {

	sessionId := utils.GetSessionIdFromContext(ctx)

	// Thank you for your patronage
	lotteryRes.PrizeId = "" //
	lotteryRes.PrizeName = cfg.NO_GIFT
	defer loadRecordToDb(ctx, userId, lotteryRes)

	lotteryCode := utils.NewLotteryCode()
	xlog.Info(sessionId, "Lottery Code", xlog.Fields{
		"userId":      userId,
		"lotteryCode": lotteryCode,
	})
	prize, err := lotteryPrize(lotteryCode)
	if err != nil {
		xlog.Error(sessionId, "Lottery Prize ERROR", xlog.Fields{
			"err":         err,
			"userId":      userId,
			"phoneNumber": phoneNumber,
			"lotteryRes":  lotteryRes,
		})
		return err
	}

	rd := xredis.GetRedisInstance()
	lockId := utils.NewRedisLockId()
	defer rd.UnLock(prize.Id, lockId)

	if !rd.SetLock(prize.Id, lockId) {
		lotteryRes.Message = "Fierce Competition"
		xlog.Error(sessionId, "Set Redis Lock Failed", xlog.Fields{
			"prize.Id": prize.Id,
			"lockId":   lockId,
		})
		return errors.New("Fierce Competition")
	}

	// Unlimited Prize
	if prize.Unlimited == 1 {
		lotteryRes.Win = true
		lotteryRes.PrizeName = prize.Name
		lotteryRes.PrizeId = prize.Id
		return nil
	}

	// No Prize Left
	if prize.Left == 0 { // no prize left
		xlog.Debug(sessionId, "No Prize Left", xlog.Fields{
			"prize": prize,
		})
		return nil
	}

	// Prize Day Count Limit
	if prize.DayLimit != 0 {
		dayCnt, err := xorm.PrizeDayCount(prize.Id)
		if err != nil {
			xlog.Error(sessionId, "Prize Day Count ERROR", xlog.Fields{
				"prize.Id": prize.Id,
				"err":      err,
			})
			return err
		}
		if dayCnt >= prize.DayLimit {
			xlog.Debug(sessionId, "Day Limit", xlog.Fields{
				"prize":  prize,
				"dayCnt": dayCnt,
			})
			return nil
		}
	}

	// Prize User Count Limit
	if prize.UserLimit != 0 {
		userCnt, err := xorm.PrizeUserCount(prize.Id, userId)
		if err != nil {
			xlog.Error(sessionId, "Prize User Count ERROR", xlog.Fields{
				"prize.Id": prize.Id,
				"userId":   userId,
				"err":      err,
			})
			return err
		}
		if userCnt >= prize.UserLimit {
			xlog.Debug(sessionId, "User Limit", xlog.Fields{
				"prize":   prize,
				"userCnt": userCnt,
			})
			return nil
		}
	}

	// Give Our Prize
	err = xorm.UpdatePrizeLeft(prize.Id)
	if err != nil {
		xlog.Error(sessionId, "Update Prize Left ERROR", xlog.Fields{
			"prize": prize,
		})
		return err
	}

	lotteryRes.Win = true
	lotteryRes.PrizeName = prize.Name
	lotteryRes.PrizeId = prize.Id
	return nil
}

// lotteryPermitted: check lottery permitted
// - permission is true
// - once a day
func lotteryPermitted(ctx context.Context, userId string) (bool, error) {
	sessionId := utils.GetSessionIdFromContext(ctx)
	// permission
	permitted, err := xorm.Permitted(userId)
	if err != nil {
		xlog.Error(sessionId, "XORM Permitted ERROR", xlog.Fields{
			"userId": userId,
			"err":    err,
		})
		return false, err
	}
	if !permitted {
		return permitted, nil
	}
	// once a day
	dayCnt, err := xorm.DaylotteryCount(userId)
	if err != nil {
		xlog.Error(sessionId, "XORM DaylotteryCount ERROR", xlog.Fields{
			"userId": userId,
			"err":    err,
		})
		return false, err
	}
	return dayCnt < 1, nil
}
