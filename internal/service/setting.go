package service

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
	"github.com/xiaohubai/go-grpc-layout/internal/errors"
	"github.com/xiaohubai/go-grpc-layout/pkg/jwt"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/request"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

func (s *HttpService) UpdateSetting(c *gin.Context) {
	req := &v1.SettingRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, errors.TokenFailed, err)
		return
	}
	claims, ok := c.Get("claims")
	if !ok {
		response.Fail(c, errors.UpdateSettingsFailed, nil)
		return
	}
	userInfo := claims.(*jwt.Claims)
	req.UID = userInfo.UID
	err = s.uc.UpdateSetting(c.Request.Context(), req)
	if err != nil {
		response.Fail(c, errors.UpdateSettingsFailed, err)
		return
	}
	response.Success(c, nil)
}

func (s *HttpService) GetSetting(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		response.Fail(c, errors.TokenFailed, nil)
		return
	}
	userInfo := claims.(*jwt.Claims)
	data, err := s.uc.GetSetting(c.Request.Context(), &model.Setting{UID: userInfo.UID})
	if err != nil {
		response.Fail(c, errors.GetSettingsFailed, err)
		return
	}
	response.Success(c, data)
}
