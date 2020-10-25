package api

import (
	"context"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

const BASE_ERR_CODE_INDEX = 70000

const (
	SUCCESS_ON_ACTION_RETCODE int = 0

	ERR_HTTP_URI       int = BASE_ERR_CODE_INDEX
	ERR_HTTP_METHOD    int = BASE_ERR_CODE_INDEX + 1
	ERR_ACTION_INVALID int = BASE_ERR_CODE_INDEX + 2

	ERR_PARSE_PARAMS_ERROR int = BASE_ERR_CODE_INDEX + 3
	ERR_XORM_ERROR         int = BASE_ERR_CODE_INDEX + 3
)

type HTTPHandler func(ctx context.Context, v interface{}) (int, interface{})

var (
	httpHandlers    map[string]HTTPHandler
	ActionStructMap map[string]reflect.Type
	NewActionMap    map[string]func(c *gin.Context, ctx context.Context) (int, interface{})
)

func InitNewAction() {
	NewActionMap = make(map[string]func(c *gin.Context, ctx context.Context) (int, interface{}))
	NewActionMap["LogIn"] = LogIn
}

func init() {
	InitNewAction()
}

//	action: CreateService,GetService...
func ActionRouter(c *gin.Context, ctx context.Context, action string, v map[string]interface{}) {
	actionDesc := action + "Response"
	var (
		code int
		res  interface{}
	)

	// 如果读取到新定义的action，则走新的body解析流程
	if handler, ok := NewActionMap[action]; ok == true {
		code, res = handler(c, ctx)
	} else {
		handler, ok := httpHandlers[action]
		if !ok {
			APIResponseError(c, actionDesc, ERR_ACTION_INVALID, "action invalid")
			return
		}

		targetType, ok := ActionStructMap[action]
		if !ok {
			APIResponseError(c, actionDesc, ERR_PARSE_PARAMS_ERROR, "target struct don't find")
			return
		}

		target, err := TransformInput(v, targetType.Elem())
		if err != nil {
			APIResponseError(c, actionDesc, ERR_PARSE_PARAMS_ERROR, err.Error())
			return
		}

		code, res = handler(ctx, target)
	}

	if code != SUCCESS_ON_ACTION_RETCODE {
		if res == nil {
			APIResponseError(c, actionDesc, code, "action response is nil")
		} else {
			APIResponseError(c, actionDesc, code, res.(string))
		}
	} else {
		c.JSON(http.StatusOK, res)
	}
}
