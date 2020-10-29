package api

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"lottery_backend/src/utils"
	"lottery_backend/src/xlog"
)

const BASE_ERR_CODE_INDEX = 70000

const (
	SUCCESS_ON_ACTION_RETCODE int = 0

	ERR_HTTP_URI       int = BASE_ERR_CODE_INDEX
	ERR_HTTP_METHOD    int = BASE_ERR_CODE_INDEX + 1
	ERR_ACTION_INVALID int = BASE_ERR_CODE_INDEX + 2

	ERR_PARSE_PARAMS_ERROR int = BASE_ERR_CODE_INDEX + 3
	ERR_XORM_ERROR         int = BASE_ERR_CODE_INDEX + 4
	ERR_LOTTERY_ERROR      int = BASE_ERR_CODE_INDEX + 5
)

type HTTPHandler func(ctx context.Context, v interface{}) (int, interface{})

var (
	//httpHandlers map[string]HTTPHandler
	ActionMap map[string]func(c *gin.Context, ctx context.Context) (int, interface{})
)

func InitAction() {
	ActionMap = make(map[string]func(c *gin.Context, ctx context.Context) (int, interface{}))
	ActionMap["LogIn"] = LogIn
	ActionMap["UserExist"] = UserExist
	ActionMap["SubmitArticle"] = SubmitArticle
	ActionMap["ListArticle"] = ListArticle
	ActionMap["Lottery"] = Lottery
	ActionMap["ListRecord"] = ListRecord
	ActionMap["SubOneDayForRecord"] = SubOneDayForRecord
}

func init() {
	InitAction()
}

//	action: CreateService,GetService...
func ActionRouter(c *gin.Context, ctx context.Context, action string, v map[string]interface{}) {
	actionDesc := action + "Response"
	var (
		code int
		res  interface{}
	)

	// 如果读取到新定义的action，则走新的body解析流程
	if handler, ok := ActionMap[action]; ok == true {
		xlog.DebugSimple(utils.GetSessionIdFromContext(ctx), xlog.Fields{
			"action": action,
		})
		code, res = handler(c, ctx)
	} else {
		xlog.ErrorSimple(utils.GetSessionIdFromContext(ctx), xlog.Fields{
			"action": action,
		})
		APIResponseError(c, actionDesc, ERR_ACTION_INVALID, "action invalid")
		return
	}
	if code != SUCCESS_ON_ACTION_RETCODE {
		if res == nil {
			APIResponseError(c, actionDesc, code, "action response is nil")
		} else {
			APIResponseError(c, actionDesc, code, res.(string))
		}
	} else {
		xlog.Info(utils.GetSessionIdFromContext(ctx), "[http-response-success]", xlog.Fields{
			"code":     code,
			"action":   action,
			"response": res,
		})
		c.JSON(http.StatusOK, res)
	}
}
