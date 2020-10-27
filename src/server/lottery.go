package api

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin"

	cfg "lottery_backend/src/config"
	"lottery_backend/src/model"
	xredis "lottery_backend/src/redis"
	"lottery_backend/src/xlog"
)

func Lottery(c *gin.Context, ctx context.Context) (int, interface{}) {
	sessionId := GetSessionIdFromContext(ctx)
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
			"res": req,
		})
		res.Message = "Lottery param invalid"
		res.RetCode = ERR_PARSE_PARAMS_ERROR
	}

	// TODO

	return SUCCESS_ON_ACTION_RETCODE, res
}

func lottery(ctx context.Context, userId string,
	phoneNumber uint64, lotteryRes *model.LotteryResponse) error {

	sessionId := GetSessionIdFromContext(ctx)

	// TODO: once a day

	// Thank you for your patronage
	lotteryRes.PrizeId = "" //
	lotteryRes.PrizeName = cfg.NO_GIFT
	defer loadRecordToDb(ctx, userId, lotteryRes)

	prize, err := lotteryPrize(NewLotteryCode())
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
	lockId := NewRedisLockId()
	defer rd.UnLock(prize.Id, lockId)

	if !rd.SetLock(prize.Id, lockId) {
		lotteryRes.Message = "Fierce Competition"
		xlog.Error(sessionId, "Set Redis Lock Failed", xlog.Fields{
			"prize.Id": prize.Id,
			"lockId":   lockId,
		})
		return errors.New("Fierce Competition")
	}

	if prize.Unlimited == 1 {
		lotteryRes.Win = true
		lotteryRes.PrizeName = prize.Id
		lotteryRes.PrizeId = prize.Name
		return nil
	}

	// TODO: prize condition

	lotteryRes.Win = true
	lotteryRes.PrizeName = prize.Id
	lotteryRes.PrizeId = prize.Name
	return nil
}

// TODO
//func confirmWin() bool {
//
//}
