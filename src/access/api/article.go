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

// SubmitArticle: submit article
func SubmitArticle(c *gin.Context, ctx context.Context) (int, interface{}) {
	sessionId := utils.GetSessionIdFromContext(ctx)
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
		return SUCCESS_ON_ACTION_RETCODE, res
	}
	_, err = xorm.InsertAndUpdateArticle(&xmodel.Article{
		//Id:     utils.NewId(),
		UserId: req.UserId,
		Text:   req.Text,
	})

	if err != nil {
		xlog.Error(sessionId, "insert article error", xlog.Fields{
			"err": err,
			"res": req,
		})
		res.Message = err.Error()
		res.RetCode = ERR_XORM_ERROR
		return SUCCESS_ON_ACTION_RETCODE, res
	}
	return SUCCESS_ON_ACTION_RETCODE, res
}

// ListArticle: list article
func ListArticle(c *gin.Context, ctx context.Context) (int, interface{}) {
	sessionId := utils.GetSessionIdFromContext(ctx)
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
	if req.Offset < 0 || req.Limit == 0 {
		xlog.Error(sessionId, "ListArticle param invalid", xlog.Fields{
			"res": req,
		})
		res.Message = "ListArticle param invalid"
		res.RetCode = ERR_PARSE_PARAMS_ERROR
		return SUCCESS_ON_ACTION_RETCODE, res
	}
	articles, err := xorm.ListUserArticle(req.Offset, req.Limit)
	if err != nil {
		xlog.Error(sessionId, "list article error", xlog.Fields{
			"err": err,
			"res": req,
		})
		res.Message = err.Error()
		res.RetCode = ERR_XORM_ERROR
		return SUCCESS_ON_ACTION_RETCODE, res
	}
	res.DataSet = articles
	return SUCCESS_ON_ACTION_RETCODE, res
}
