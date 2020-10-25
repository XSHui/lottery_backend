package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
	c.AbortWithStatusJSON(http.StatusOK, &Response{
		Action:  action,
		RetCode: code,
		Message: message,
	})
}

func HttpMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//ctx, _ = common.NewContext(context.Background())
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

	router.POST(BASE_URI, func(c *gin.Context) {
		// 此处为了兼容新的body解析方式而重新将Body赋值
		reader, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(reader))

		v := make(map[string]interface{}, 0)
		err := json.NewDecoder(bytes.NewReader(reader)).Decode(&v)
		if err != nil {
			APIResponseError(c, "", ERR_PARSE_PARAMS_ERROR, "request body json decode error "+err.Error())
			return
		}
		reqAction, ok := v["Action"]
		if !ok {
			APIResponseError(c, "", ERR_ACTION_INVALID, "miss action")
			return
		}
		newCtx := NewContextWithSession(ctx, NewSessionId())
		ActionRouter(c, newCtx, reqAction.(string), v)
	})

	// 启动 HTTP 服务
	srv := &http.Server{
		Handler:     router,
		Addr:        fmt.Sprintf("%s:%d", ip, port),
		ReadTimeout: 5 * time.Second,
	}

	srv.ListenAndServe()
}
