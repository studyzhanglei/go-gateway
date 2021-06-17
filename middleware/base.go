package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gateway/global"
	"go-gateway/model/response"
	"go-gateway/utils"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func Base() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Writer.Status() == http.StatusNotFound {
			c.JSON(http.StatusNotFound, &response.Response{
				Code: http.StatusNotFound,
				Data: map[string]interface{}{},
				Msg: "404 not found",
			})
			c.Abort()
			return
		}

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		//新增request-id
		requestId := string(c.Request.Header.Get("X-request-id"))
		c.Set("trace-id", requestId)
		c.Set(utils.REQUET_LOG_KEY, global.LOG.With(zap.String("trace-id", requestId)))
		c.Header("X-request-id", requestId)  	//给客户端返回X-request-id以便于排错


		//记录请求日志
		data, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) //body复用 不然读取一次就没了
		utils.GetTraceLog(c).Info(fmt.Sprintf("request：body %s header %s", string(data), string(fmt.Sprintf("%v", c.Request.Header))))

		c.Next()

		// 处理请求
		if blw.body.String() == "" {
			response.Result(200, map[string]interface{}{}, "success", c)
		}


		//记录响应日志
		utils.GetTraceLog(c).Info(fmt.Sprintf("response：%s", blw.body.String()))
	}
}
