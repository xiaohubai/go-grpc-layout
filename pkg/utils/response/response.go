package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/xiaohubai/go-grpc-layout/internal/errors"
	"github.com/xiaohubai/go-grpc-layout/pkg/tracing"
)

type Body struct {
	Code    int         `json:"code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
	TraceID string      `json:"traceID"`
}

func Result(c *gin.Context, code int, data interface{}, err interface{}) {
	if data == nil {
		data = make(map[string]string, 0)
	}
	resp := Body{
		Code:    code,
		Data:    data,
		Msg:     errors.Msg[code],
		TraceID: tracing.TraceID(c.Request.Context()),
	}

	if e, ok := err.(error); ok {
		resp.Msg = resp.Msg + ": " + e.Error()
	}
	c.JSON(http.StatusOK, resp)
}

func Success(c *gin.Context, data interface{}) {
	Result(c, errors.Success, data, nil)
}

func Fail(c *gin.Context, code int, err interface{}) {
	Result(c, code, nil, err)
}
