package service

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
	"github.com/xiaohubai/go-grpc-layout/internal/errors"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/request"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

func (s *HttpService) GetCasbinList(c *gin.Context) {
	req := &v1.GetCasbinRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, errors.ParamsFailed, err)
		return
	}

	data, err := s.uc.GetCasbinList(c.Request.Context(),
		&model.CasbinRule{
			V0: req.RoleIDs,
			V1: req.Path,
			V2: req.Method,
		}, &v1.PageRequest{Page: req.Page, PageSize: req.PageSize})
	if err != nil {
		response.Fail(c, errors.GetCasbinListFailed, err)
		return
	}
	response.Success(c, data)
}

func (s *HttpService) AddCasbin(c *gin.Context) {
	req := &v1.GetCasbinRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, errors.ParamsFailed, err)
		return
	}

	data, err := s.uc.GetCasbinList(c.Request.Context(),
		&model.CasbinRule{
			V0: req.RoleIDs,
			V1: req.Path,
			V2: req.Method,
		}, &v1.PageRequest{Page: req.Page, PageSize: req.PageSize})
	if err != nil {
		response.Fail(c, errors.GetCasbinListFailed, err)
		return
	}
	response.Success(c, data)
}

func (s *HttpService) DeleteCasbin(c *gin.Context) {
	req := &v1.GetCasbinRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, errors.ParamsFailed, err)
		return
	}

	data, err := s.uc.DeleteCasbin(c.Request.Context(),
		&model.CasbinRule{
			V0: req.RoleIDs,
			V1: req.Path,
			V2: req.Method,
		}, &v1.PageRequest{Page: req.Page, PageSize: req.PageSize})
	if err != nil {
		response.Fail(c, errors.GetCasbinListFailed, err)
		return
	}
	response.Success(c, data)
}

func (s *HttpService) UpdateCasbin(c *gin.Context) {
	req := &v1.GetCasbinRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, errors.ParamsFailed, err)
		return
	}

	data, err := s.uc.UpdateCasbin(c.Request.Context(),
		&model.CasbinRule{
			V0: req.RoleIDs,
			V1: req.Path,
			V2: req.Method,
		}, &v1.PageRequest{Page: req.Page, PageSize: req.PageSize})
	if err != nil {
		response.Fail(c, errors.GetCasbinListFailed, err)
		return
	}
	response.Success(c, data)
}
