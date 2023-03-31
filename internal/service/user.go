package service

import (
	"context"

	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-grpc-layout/api/errors/v1"
	gpb "github.com/xiaohubai/go-grpc-layout/api/grpc/v1"
	hpb "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/model"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/request"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils/response"
)

func (s *HttpService) Login(c *gin.Context) {
	req := &hpb.LoginRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, v1.Error_RequestFail, "请求失败", err)
		return
	}

	data, err := s.uc.Login(c.Request.Context(), &model.User{
		UserName: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		response.Fail(c, v1.Error_LoginFail, "登录失败", err)
		return
	}
	response.Success(c, data)
}

func (s *GrpcService) GetUserInfo(ctx context.Context, req *gpb.UserInfoRequest) (*gpb.UserInfoResponse, error) {
	return &gpb.UserInfoResponse{}, nil
}
