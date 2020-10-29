package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	//"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"lottery_backend/src/utils"
	"lottery_backend/src/xlog"
)

const (
	BASE_URI = "/"
)

var ctx context.Context

type Response struct {
	Action  string `json:"Action,omitempty" comment:"Response Action"`
	RetCode int    `json:"RetCode" comment:"返回码，0表示成功，其它表示错误"`
	Message string `json:"Message" comment:"返回消息"`
}

func APIResponseError(c *gin.Context, action string, code int, message string) {
	xlog.ErrorSimple("[http-response-error]", xlog.Fields{
		"action":  action,
		"code":    code,
		"message": message,
	})
	c.AbortWithStatusJSON(http.StatusOK, &Response{
		Action:  action,
		RetCode: code,
		Message: message,
	})
}

func HttpMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionId := ""
		ctx, sessionId = utils.NewContext(context.Background())
		xlog.Debug(sessionId, "[http-request]", xlog.Fields{
			"uri":    c.Request.RequestURI,
			"method": c.Request.Method,
		})
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "false")
		c.Set("content-type", "application/json")
		if c.Request.RequestURI != BASE_URI {
			APIResponseError(c, "", ERR_HTTP_URI, "no support this uri "+c.Request.RequestURI)
			return
		}
		if c.Request.Method == http.MethodPost {
			c.Next()
		} else {
			APIResponseError(c, "", ERR_HTTP_METHOD, "request method must be POST, no support this http method "+c.Request.Method)
		}
	}
}

func StartHttpServer(ip string, port int) {

	router := gin.Default()

	router.Use(HttpMiddleWare())

	xlog.DebugSimple("Starting Http Server....", xlog.Fields{})

	router.POST(BASE_URI, func(c *gin.Context) {
		reader, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(reader))

		v := make(map[string]interface{}, 0)
		err := json.NewDecoder(bytes.NewReader(reader)).Decode(&v)
		if err != nil {
			APIResponseError(c, "", ERR_PARSE_PARAMS_ERROR, "request body json decode error "+err.Error())
			return
		}
		xlog.DebugSimple("[http-request-post-body]", xlog.Fields{
			"request": v,
			"err":     err,
		})
		reqAction, ok := v["Action"]
		if !ok {
			xlog.ErrorSimple("No Action", xlog.Fields{})
			APIResponseError(c, "", ERR_ACTION_INVALID, "miss action")
			return
		}
		xlog.DebugSimple("after [http-request-post-body]", xlog.Fields{
			"reqAction": reqAction,
		})
		ActionRouter(c, ctx, reqAction.(string), v)
	})

	// 启动 HTTP 服务
	srv := &http.Server{
		Handler:     router,
		Addr:        fmt.Sprintf("%s:%d", ip, port),
		ReadTimeout: 5 * time.Second,
	}

	xlog.DebugSimple("Http Server Started", xlog.Fields{})
	srv.ListenAndServe()
}
