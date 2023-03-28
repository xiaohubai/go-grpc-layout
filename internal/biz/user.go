package biz

import (
	"context"

	"github.com/spf13/cast"
	pb "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/pkg/tracing"
	"go.opentelemetry.io/otel/attribute"
)

type User struct {
	ID    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Login 用户登录
func (uc *HttpUsecase) Login(ctx context.Context, u *User) (*pb.Data, error) {
	ctx, span := tracing.NewSpan(ctx, "biz-Login")
	defer span.End()

	userList, err := uc.repo.GetUserInfo(ctx, u)
	if err != nil {
		return nil, err
	}
	span.SetAttributes(attribute.Key("user").String(cast.ToString(userList)))
	result := &pb.Data{
		UserName: "xxx",
		NickName: "sss",
	}
	return result, nil
}
