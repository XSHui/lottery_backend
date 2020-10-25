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

// LogIn: new user login
func SubmitArticle(c *gin.Context, ctx context.Context) (int, interface{}) {
	sessionId := GetSessionIdFromContext(ctx)
	res := model.SubmitArticleResponse{
		Action:  "SubmitArticleResponse",
		RetCode: 0,
	}
	var req model.SubmitArticleRequest
	err := json.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		xlog.Error(sessionId, "NewDecoder error", xlog.Fields{
			"request": req,
		})
		res.Message = err.Error()
		res.RetCode = ERR_PARSE_PARAMS_ERROR
		return SUCCESS_ON_ACTION_RETCODE, res
	}
	if req.UserId == "" || req.Text == "" || len(req.Text) > 500 {
		xlog.Error(sessionId, "SubmitArticle param invalid", xlog.Fields{
			"res": req,
		})
		res.Message = "SubmitArticle param invalid"
		res.RetCode = ERR_PARSE_PARAMS_ERROR
	}
	err = xorm.InsertArticle(&xmodel.Article{
		Id:          NewId(),
		UserId: req.UserId,
		Text: req.Text,
	})

	if err != nil {
		xlog.Error(sessionId, "insert article error", xlog.Fields{
			"err": err,
			"res": req,
		})
		res.Message = err.Error()
		res.RetCode = ERR_XORM_ERROR
	}
	return SUCCESS_ON_ACTION_RETCODE, res
}

// LogIn: new user login
func ListArticle(c *gin.Context, ctx context.Context) (int, interface{}) {
	sessionId := GetSessionIdFromContext(ctx)
	res := model.ListArticleResponse{
		Action:  "ListArticleResponse",
		RetCode: 0,
	}
	var req model.ListArticleRequest
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
		xlog.Error(sessionId, "ListArticle param invalid", xlog.Fields{
			"res": req,
		})
		res.Message = "ListArticle param invalid"
		res.RetCode = ERR_PARSE_PARAMS_ERROR
	}
	articles, err := xorm.ListUserArticle(req.Offset, req.Limit)
	if err != nil {
		xlog.Error(sessionId, "list article error", xlog.Fields{
			"err": err,
			"res": req,
		})
		res.Message = err.Error()
		res.RetCode = ERR_XORM_ERROR
	}
	res.DataSet = articles
	return SUCCESS_ON_ACTION_RETCODE, res
}



