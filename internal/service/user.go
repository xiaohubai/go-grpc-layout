package service

import (
	"context"

	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-grpc-layout/api/errors/v1"
	gpb "github.com/xiaohubai/go-grpc-layout/api/grpc/v1"
	hpb "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/biz"
	"github.com/xiaohubai/go-grpc-layout/pkg/request"
	"github.com/xiaohubai/go-grpc-layout/pkg/response"
)

func (s *HttpService) Login(c *gin.Context) {
	req := &hpb.LoginRequest{}
	err := request.ShouldBindJSON(c, req)
	if err != nil {
		response.Fail(c, v1.Error_RequestFail, err)
		return
	}

	data, _ := s.uc.Login(c.Request.Context(), &biz.User{
		UserName: req.UserName,
		Password: req.Password,
	})
	response.Ok(c, data)
}

func (s *GrpcService) GetUserInfo(ctx context.Context, req *gpb.UserInfoRequest) (*gpb.UserInfoResponse, error) {
	return &gpb.UserInfoResponse{}, nil
}
