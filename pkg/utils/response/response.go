package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	v1 "github.com/xiaohubai/go-grpc-layout/api/errors/v1"
	"github.com/xiaohubai/go-grpc-layout/pkg/tracing"
)

type Body struct {
	Code    int32       `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	TraceID string      `json:"traceID"`
}

func Result(c *gin.Context, code int32, data interface{}, opts ...interface{}) {
	if data == nil {
		data = make(map[string]string, 0)
	}
	var msg string
	for _, m := range opts {
		msg = msg + ": " + fmt.Sprintf("%s", m)
	}
	resp := Body{
		Code:    code,
		Data:    data,
		Msg:     v1.Error_name[code] + msg,
		TraceID: tracing.TraceID(c.Request.Context()),
	}
	c.JSON(http.StatusOK, resp)
}

func Success(c *gin.Context, data interface{}, opts ...interface{}) {
	Result(c, int32(v1.Error_Success), data, opts)
}

func Fail(c *gin.Context, code v1.Error, opts ...interface{}) {
	Result(c, int32(code), nil, opts)
}
