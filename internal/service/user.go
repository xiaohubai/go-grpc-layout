package service

import (
	"context"

	pb "github.com/xiaohubai/go-grpc-layout/api/admin/v1"
)

func (s *Service) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	result, err := s.uc.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return result, nil
}
