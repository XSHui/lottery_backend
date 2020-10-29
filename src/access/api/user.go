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

// LogIn: new user login
func LogIn(c *gin.Context, ctx context.Context) (int, interface{}) {
	sessionId := utils.GetSessionIdFromContext(ctx)

	xlog.DebugSimple(sessionId, xlog.Fields{
		"LogIn": LogIn,
	})

	res := model.LogInResponse{
		Action:  "LogInResponse",
		RetCode: 0,
	}
	var req model.LogInRequest
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		xlog.Error(sessionId, "NewDecoder error", xlog.Fields{
			"request": req,
		})
		res.Message = err.Error()
		res.RetCode = ERR_PARSE_PARAMS_ERROR
		return SUCCESS_ON_ACTION_RETCODE, res
	}
	if req.PhoneNumber == 0 { // TODO: check phone number
		xlog.Error(sessionId, "LogIn param invalid", xlog.Fields{
			"res": req,
		})
		res.Message = "LogIn param invalid"
		res.RetCode = ERR_PARSE_PARAMS_ERROR
		return SUCCESS_ON_ACTION_RETCODE, res
	}
	// TODO: phone number check
	userId := utils.NewId()
	err = xorm.InsertUser(&xmodel.User{
		Id:          userId,
		PhoneNumber: req.PhoneNumber,
		CreateTime:  utils.NowTimestamp(),
		ModifyTime:  utils.NowTimestamp(),
	})
	if err != nil {
		xlog.Error(sessionId, "insert user error", xlog.Fields{
			"err": err,
			"res": req,
		})
		res.Message = err.Error()
		res.RetCode = ERR_XORM_ERROR
		return SUCCESS_ON_ACTION_RETCODE, res
	}
	res.UserId = userId
	return SUCCESS_ON_ACTION_RETCODE, res
}

// UserExist: check whether user exist, if not send verification code
func UserExist(c *gin.Context, ctx context.Context) (int, interface{}) {
	sessionId := utils.GetSessionIdFromContext(ctx)
	res := model.UserExistResponse{
		Action:  "UserExistResponse",
		RetCode: 0,
	}
	var req model.UserExistRequest
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
		xlog.Error(sessionId, "UserExist param invalid", xlog.Fields{
			"res": req,
		})
		res.Message = "UserExist param invalid"
		res.RetCode = ERR_PARSE_PARAMS_ERROR
		return SUCCESS_ON_ACTION_RETCODE, res
	}
	_, err = xorm.GetUserInfoByPhoneNum(req.PhoneNumber)
	if err != nil {
		xlog.Error(sessionId, "insert user error", xlog.Fields{
			"err": err,
			"res": req,
		})
		res.Message = err.Error()
		res.RetCode = ERR_XORM_ERROR
		return SUCCESS_ON_ACTION_RETCODE, res
	}
	res.Exist = true
	return SUCCESS_ON_ACTION_RETCODE, res
}
