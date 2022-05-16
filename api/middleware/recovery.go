package middleware

import (
	"blockchain-deal-hunter/api/utility"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"os"
	"runtime/debug"
	"strings"
)

func ResponseOutput(c *gin.Context){
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*utility.TraceContext)
	traceId := ""
	if traceContext != nil {
		traceId = traceContext.TraceId
	}

	ref := ""
	code := "ok"
	statusCode := 200
	var msg string
	var responseErr error
	if respErr, set := c.Get("responseErr"); set{
		code = "err"
		if responseCode, set := c.Get("responseCode"); set {
			code = responseCode.(string)
		}
		switch respErr.(type) {
		case utility.ParamError:
			responseErr = respErr.(utility.ParamError)
			code = "param_err"
			msg = "请求错误, 请检查数据"
			statusCode = 400
		default :
			responseErr = respErr.(utility.FatalError)
			code = "fatal_err"
			msg = "系统错误, 请稍候再试"
			statusCode = 500
		}
		if c.Query("is_debug") == "1" || os.Getenv("is_debug") == "1" {
			ref = strings.Replace(fmt.Sprintf("%+v", responseErr), responseErr.Error()+"\n", "", -1)
		}
	}
	if respMsg, set := c.Get("responseMsg"); set{
		msg = respMsg.(string)
	}
	responseData := make(map[string]interface{})
	if respData, set := c.Get("responseData"); set{
		responseData = respData.(map[string]interface{})
	}
	output := &struct {
		Code string `json:"code"`
		Msg  string       `json:"msg"`
		Data      interface{}  `json:"data"`
		TraceId   string  `json:"tid"`
		Ref     interface{}  `json:"ref"`
	}{code, msg, responseData, traceId, ref}
	c.JSON(statusCode, output)

	response, _ := json.Marshal(output)

	c.Set("response", string(response))

	if responseErr != nil{
		c.AbortWithError(statusCode, responseErr)
	}
}
// RecoveryMiddleware捕获所有panic，并且返回错误信息
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Error(map[string]interface{}{
					"error": fmt.Sprint(err),
					"stack": string(debug.Stack()),
				})
				c.Set("responseErr", utility.FatalError{err.(error)})
				ResponseOutput(c)
			}
		}()
		c.Next()

		ResponseOutput(c)
	}
}