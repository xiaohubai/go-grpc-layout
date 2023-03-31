package biz

import (
	"context"

	"github.com/spf13/cast"
	pb "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/model"
	"github.com/xiaohubai/go-grpc-layout/pkg/tracing"
	"go.opentelemetry.io/otel/attribute"
)

// Login 用户登录
func (uc *HttpUsecase) Login(ctx context.Context, u *model.User) (*pb.LoginResponse, error) {
	ctx, span := tracing.NewSpan(ctx, "biz-Login")
	defer span.End()

	userList, err := uc.repo.ListAllUser(ctx, u)
	if err != nil {
		return nil, err
	}
	span.SetAttributes(attribute.Key("userList").String(cast.ToString(userList)))
	result := &pb.LoginResponse{
		UserName: "xxx",
		NickName: "sss",
	}
	return result, nil
}
