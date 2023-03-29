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

func Result(c *gin.Context, code int32, data, msg interface{}) {
	resp := Body{
		Code:    code,
		Data:    data,
		Msg:     v1.Error_name[code],
		TraceID: tracing.TraceID(c.Request.Context()),
	}
	if e, ok := msg.(error); ok {
		resp.Msg = fmt.Sprintf("%s：%s", resp.Msg, e.Error())
	}
	c.JSON(http.StatusOK, resp)
}

func Ok(c *gin.Context, data interface{}) {
	if data == nil {
		data = make(map[string]string, 0)
	}
	Result(c, int32(v1.Error_Success), data, "")
}

func Fail(c *gin.Context, code v1.Error, err interface{}) {
	data := make(map[string]string, 0)
	Result(c, int32(code), data, err)
}
