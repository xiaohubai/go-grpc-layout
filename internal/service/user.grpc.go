package service

import (
	"context"

	v1 "github.com/xiaohubai/go-grpc-layout/api/grpc/v1"
)

func (s *GrpcService) GetUserInfo(ctx context.Context, req *v1.UserInfoRequest) (*v1.UserInfoResponse, error) {
	return &v1.UserInfoResponse{
		Code: 21312,
	}, nil
}
