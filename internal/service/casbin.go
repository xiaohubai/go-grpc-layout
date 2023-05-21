package service

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/ecode"

	"github.com/xiaohubai/go-grpc-layout/pkg/utils/request"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

func (s *HttpService) GetRoleCasbinList(c *gin.Context) {
	req := &v1.GetCasbinRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, ecode.ParamsFailed, err)
		return
	}
	data, err := s.uc.GetRoleCasbinList(c, req)
	if err != nil {
		response.Fail(c, ecode.GetCasbinListFailed, err)
		return
	}
	response.Success(c, data)
}

func (s *HttpService) AddRoleCasbin(c *gin.Context) {
	req := &v1.AddCasbinRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, ecode.ParamsFailed, err)
		return
	}
	err = s.uc.AddRoleCasbin(c, req)
	if err != nil {
		response.Fail(c, ecode.GetCasbinListFailed, err)
		return
	}
	response.Success(c, nil)
}

func (s *HttpService) DeleteRoleCasbin(c *gin.Context) {
	req := &v1.DeleteCasbinRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, ecode.ParamsFailed, err)
		return
	}
	err = s.uc.DeleteRoleCasbin(c, req)
	if err != nil {
		response.Fail(c, ecode.GetCasbinListFailed, err)
		return
	}
	response.Success(c, nil)
}

func (s *HttpService) UpdateRoleCasbin(c *gin.Context) {
	req := &v1.UpdateCasbinRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, ecode.ParamsFailed, err)
		return
	}
	err = s.uc.UpdateRoleCasbin(c, req)
	if err != nil {
		response.Fail(c, ecode.GetCasbinListFailed, err)
		return
	}
	response.Success(c, nil)
}
