package service

import (
	"context"

	"github.com/spf13/cast"
	pb "github.com/xiaohubai/go-grpc-layout/api/admin/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/biz"
	"github.com/xiaohubai/go-grpc-layout/pkg/tracing"
	"go.opentelemetry.io/otel/attribute"
)

func (s *HttpService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	ctx, span := tracing.NewSpan(ctx, "service-Login")
	defer span.End()
	//处理入参和 反参
	user := &biz.User{}

	data, err := s.uc.Login(ctx, user)
	if err != nil {
		return nil, err
	}

	resp := &pb.LoginResponse{
		Code:    200,
		Msg:     "success",
		TraceID: tracing.TraceID(ctx),
		Data:    data,
	}

	span.SetAttributes(attribute.Key("resp").String(cast.ToString(resp)))
	return resp, nil
}
