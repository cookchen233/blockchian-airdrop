package middleware

import (
	"blockchain-deal-hunter/api/utility"
	"bytes"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

// 请求进入日志
func requestInput(c *gin.Context) {
	traceContext := utility.NewTrace()
	if traceId := c.Request.Header.Get("com-header-rid"); traceId != "" {
		traceContext.TraceId = traceId
	}
	if spanId := c.Request.Header.Get("com-header-spanid"); spanId != "" {
		traceContext.SpanId = spanId
	}

	c.Set("startExecTime", time.Now())
	c.Set("trace", traceContext)

	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes)) // Write body back

	log.Info([]string{"request", c.Request.Method, c.Request.RequestURI, string(bodyBytes)}, map[string]interface{}{
		"tid" : traceContext.TraceId,
		"ip":   c.ClientIP(),
		"csid" : traceContext.CSpanId,
		"sid" : traceContext.SpanId,
	})
}

// 请求输出日志
func responseOutput(c *gin.Context) {
	// after request
	endExecTime := time.Now()
	response, _ := c.Get("response")
	st, _ := c.Get("startExecTime")

	startExecTime, _ := st.(time.Time)
	var traceContext *utility.TraceContext
	tc, exists := c.Get("trace")
	if exists {
		traceContext, _ = tc.(*utility.TraceContext)
	}else{

		traceContext = utility.NewTrace()
	}
	log.Info([]string{"response", c.Request.Method, c.Request.RequestURI, response.(string)}, map[string]interface{}{
		"tid" : traceContext.TraceId,
		"ip":      c.ClientIP(),
		"csid" : traceContext.CSpanId,
		"sid" : traceContext.SpanId,
		"runtime": endExecTime.Sub(startExecTime).Seconds(),
	})
}

func RequestLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestInput(c)
		defer func() {
			responseOutput(c)
		}()

		c.Next()

	}
}