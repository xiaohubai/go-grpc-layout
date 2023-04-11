package service

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
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
	if !store.Verify(req.CaptchaID, req.Captcha, true) {
		response.Fail(c, errors.CaptchaFailed, nil)
		return
	}

	data, err := s.uc.Login(c.Request.Context(), &model.User{
		Username: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		response.Fail(c, errors.LoginFailed, err)
		return
	}
	response.Success(c, data)
}
