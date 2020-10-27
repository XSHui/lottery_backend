package api

import (
	"context"
	"encoding/json"

	"github.com/gin-gonic/gin"

	"lottery_backend/src/model"
	"lottery_backend/src/xlog"
	"lottery_backend/src/xorm"
	xmodel "lottery_backend/src/xorm/model"
)

// ListRecord: list user record
func ListRecord(c *gin.Context, ctx context.Context) (int, interface{}) {
	sessionId := GetSessionIdFromContext(ctx)
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
	if req.Offset == 0 || req.Limit == 0 {
		xlog.Error(sessionId, "ListRecord param invalid", xlog.Fields{
			"res": req,
		})
		res.Message = "ListRecord param invalid"
		res.RetCode = ERR_PARSE_PARAMS_ERROR
	}
	records, err := xorm.ListRecord(req.Offset, req.Limit)
	if err != nil {
		xlog.Error(sessionId, "list record error", xlog.Fields{
			"err": err,
			"res": req,
		})
		res.Message = err.Error()
		res.RetCode = ERR_XORM_ERROR
	}
	res.DataSet = records
	return SUCCESS_ON_ACTION_RETCODE, res
}

// load record to db
func loadRecordToDb(ctx context.Context, userId string, lotteryRes *model.LotteryResponse) {
	sessionId := GetSessionIdFromContext(ctx)
	if !lotteryRes.Win {
		xlog.Info(sessionId, "DOES NOT WIN", xlog.Fields{
			"userId": userId,
		})
		return
	}
	go func() {
		err := xorm.InsertRecord(&xmodel.Record{
			Id:         NewId(),
			UserId:     userId,
			PrizeId:    lotteryRes.PrizeId,
			CreateTime: NowTimestamp(),
			ModifyTime: NowTimestamp(),
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
