package service

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/errors"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/request"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

func (s *HttpService) GetRoleCasbinList(c *gin.Context) {
	req := &v1.GetCasbinRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, errors.ParamsFailed, err)
		return
	}
	data, err := s.uc.GetRoleCasbinList(c, req)
	if err != nil {
		response.Fail(c, errors.GetCasbinListFailed, err)
		return
	}
	response.Success(c, data)
}

func (s *HttpService) AddRoleCasbin(c *gin.Context) {
	req := &v1.AddCasbinRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, errors.ParamsFailed, err)
		return
	}
	err = s.uc.AddRoleCasbin(c, req)
	if err != nil {
		response.Fail(c, errors.GetCasbinListFailed, err)
		return
	}
	response.Success(c, nil)
}

func (s *HttpService) DeleteRoleCasbin(c *gin.Context) {
	req := &v1.DeleteCasbinRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, errors.ParamsFailed, err)
		return
	}
	err = s.uc.DeleteRoleCasbin(c, req)
	if err != nil {
		response.Fail(c, errors.GetCasbinListFailed, err)
		return
	}
	response.Success(c, nil)
}

func (s *HttpService) UpdateRoleCasbin(c *gin.Context) {
	req := &v1.UpdateCasbinRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, errors.ParamsFailed, err)
		return
	}
	err = s.uc.UpdateRoleCasbin(c, req)
	if err != nil {
		response.Fail(c, errors.GetCasbinListFailed, err)
		return
	}
	response.Success(c, nil)
}
