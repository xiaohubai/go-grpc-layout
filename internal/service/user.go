package service

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/errors"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/request"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

func (s *HttpService) Login(c *gin.Context) {
	req := &v1.LoginRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, errors.ParamsFailed, err)
		return
	}
	/* 	if !store.Verify(req.CaptchaID, req.Captcha, true) {
		response.Fail(c, errors.CaptchaFailed, nil)
		return
	} */

	data, err := s.uc.Login(c, req)
	if err != nil {
		response.Fail(c, errors.LoginFailed, err)
		return
	}
	response.Success(c, data)
}

func (s *HttpService) GetUserInfo(c *gin.Context) {
	data, err := s.uc.GetUserInfo(c)
	if err != nil {
		response.Fail(c, errors.GetUserInfoFailed, err)
		return
	}
	response.Success(c, data)
}

func (s *HttpService) UpdateUserInfo(c *gin.Context) {
	req := &v1.UpdateUserInfoRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, errors.ParamsFailed, err)
		return
	}
	err = s.uc.UpdateUserInfo(c, req)
	if err != nil {
		response.Fail(c, errors.UpdateUserInfoFailed, err)
		return
	}
	response.Success(c, nil)
}

func (s *HttpService) UpdatePassword(c *gin.Context) {
	req := &v1.UpdatePasswordRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, errors.ParamsFailed, err)
		return
	}
	err = s.uc.UpdatePassword(c, req)
	if err != nil {
		response.Fail(c, errors.UpdatePasswordFailed, err)
		return
	}
	response.Success(c, nil)
}
