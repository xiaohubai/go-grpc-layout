package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/ecode"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/request"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

func (s *HttpService) DebugPerf(c *gin.Context) {
	req := &v1.DebugPerfRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, ecode.ParamsFailed, err)
		return
	}
	res, err := s.uc.DebugPerf(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	response.Success(c, res)
}
