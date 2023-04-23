package service

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/errors"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/request"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

func (s *HttpService) GetAllMenuList(c *gin.Context) {
	data, err := s.uc.GetAllMenuList(c)
	if err != nil {
		response.Fail(c, errors.MenuListFailed, err)
		return
	}
	response.Success(c, data)
}

func (s *HttpService) GetRoleMenuList(c *gin.Context) {
	data, err := s.uc.GetRoleMenuList(c)
	if err != nil {
		response.Fail(c, errors.MenuListFailed, err)
		return
	}
	response.Success(c, data)
}

func (s *HttpService) AddRoleMenu(c *gin.Context) {
	req := &v1.AddRoleMenuRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, errors.ParamsFailed, err)
		return
	}
	err = s.uc.AddRoleMenu(c, req)
	if err != nil {
		response.Fail(c, errors.MenuListFailed, err)
		return
	}
	response.Success(c, nil)
}

func (s *HttpService) DeleteRoleMenuByID(c *gin.Context) {
	req := &v1.DeleteRoleMenuRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, errors.ParamsFailed, err)
		return
	}
	err = s.uc.DeleteRoleMenuByID(c, req)
	if err != nil {
		response.Fail(c, errors.MenuListFailed, err)
		return
	}
	response.Success(c, nil)
}

func (s *HttpService) UpdateRoleMenu(c *gin.Context) {
	req := &v1.UpdateRoleMenuRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, errors.ParamsFailed, err)
		return
	}
	err = s.uc.UpdateRoleMenu(c, req)
	if err != nil {
		response.Fail(c, errors.MenuListFailed, err)
		return
	}
	response.Success(c, nil)
}
