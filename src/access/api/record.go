package api

import (
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"

	"lottery_backend/src/access/model"
	"lottery_backend/src/utils"
	"lottery_backend/src/xlog"
	"lottery_backend/src/xorm"
	xmodel "lottery_backend/src/xorm/model"
)

// ListRecord: list user record
func ListRecord(c *gin.Context, ctx context.Context) (int, interface{}) {
	sessionId := utils.GetSessionIdFromContext(ctx)
	res := model.ListRecordResponse{
		Action:  "ListRecordResponse",
		RetCode: 0,
	}
	var req model.ListRecordRequest
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		xlog.Error(sessionId, "NewDecoder error", xlog.Fields{
			"request": req,
		})
		res.Message = err.Error()
		res.RetCode = ERR_PARSE_PARAMS_ERROR
		return SUCCESS_ON_ACTION_RETCODE, res
	}
	if req.Offset < 0 || req.Limit == 0 {
		xlog.Error(sessionId, "ListRecord param invalid", xlog.Fields{
			"res": req,
		})
		res.Message = "ListRecord param invalid"
		res.RetCode = ERR_PARSE_PARAMS_ERROR
		return SUCCESS_ON_ACTION_RETCODE, res
	}
	records, err := xorm.ListRecord(req.Offset, req.Limit)
	if err != nil {
		xlog.Error(sessionId, "list record error", xlog.Fields{
			"err": err,
			"res": req,
		})
		res.Message = err.Error()
		res.RetCode = ERR_XORM_ERROR
		return SUCCESS_ON_ACTION_RETCODE, res
	}
	res.DataSet = records
	return SUCCESS_ON_ACTION_RETCODE, res
}

// loadRecordToDb: asynchronous load record to db
func loadRecordToDb(ctx context.Context, userId string, lotteryRes *model.LotteryResponse) {
	sessionId := utils.GetSessionIdFromContext(ctx)
	if !lotteryRes.Win {
		xlog.Info(sessionId, "DOES NOT WIN", xlog.Fields{
			"userId": userId,
		})
		return
	}
	go func() {
		err := xorm.InsertRecord(&xmodel.Record{
			Id:         utils.NewId(),
			UserId:     userId,
			PrizeId:    lotteryRes.PrizeId,
			CreateTime: utils.NowTimestamp(),
			ModifyTime: utils.NowTimestamp(),
		})
		if err != nil {
			xlog.Error(sessionId, "Insert Record ERROR", xlog.Fields{
				"userId":  userId,
				"prizeId": lotteryRes.PrizeId,
			})
		}
		return
	}()
}

// SubOneDayForRecord: sub one day for record
func SubOneDayForRecord(c *gin.Context, ctx context.Context) (int, interface{}) {
	sessionId := utils.GetSessionIdFromContext(ctx)
	res := model.SubOneDayForRecordResponse{
		Action:  "SubOneDayForRecordResponse",
		RetCode: 0,
	}
	var req model.SubOneDayForRecordRequest
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		xlog.Error(sessionId, "NewDecoder error", xlog.Fields{
			"request": req,
		})
		res.Message = err.Error()
		res.RetCode = ERR_PARSE_PARAMS_ERROR
		return SUCCESS_ON_ACTION_RETCODE, res
	}
	err = xorm.SubOneDayForRecord()
	if err != nil {
		xlog.Error(sessionId, "sub record time error", xlog.Fields{
			"err": err,
		})
		res.Message = err.Error()
		res.RetCode = ERR_XORM_ERROR
		return SUCCESS_ON_ACTION_RETCODE, res
	}
	return SUCCESS_ON_ACTION_RETCODE, res
}
