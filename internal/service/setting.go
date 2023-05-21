package service

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/ecode"

	"github.com/xiaohubai/go-grpc-layout/pkg/utils/request"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

func (s *HttpService) UpdateSetting(c *gin.Context) {
	req := &v1.SettingRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, ecode.ParamsFailed, err)
		return
	}
	err = s.uc.UpdateSetting(c, req)
	if err != nil {
		response.Fail(c, ecode.UpdateSettingsFailed, err)
		return
	}
	response.Success(c, nil)
}

func (s *HttpService) GetSetting(c *gin.Context) {
	data, err := s.uc.GetSetting(c)
	if err != nil {
		response.Fail(c, ecode.GetSettingsFailed, err)
		return
	}
	response.Success(c, data)
}
