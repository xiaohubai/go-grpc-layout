package biz

import (
	"context"

	"github.com/spf13/cast"
	pb "github.com/xiaohubai/go-grpc-layout/api/admin/v1"
	"github.com/xiaohubai/go-grpc-layout/pkg/tracing"
	"go.opentelemetry.io/otel/attribute"
)

type User struct {
	ID    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Login 用户登录
func (uc *Usecase) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	_, span := tracing.NewSpan(ctx, "")

	_, err := uc.repo.GetUserInfo(ctx, &User{})
	if err != nil {
		return nil, err
	}

	result := &pb.LoginResponse{
		Code: 200,
		Msg:  "success",
		Data: &pb.LoginResponse_Data{
			DictMap: map[string]string{
				"ss": "sss",
			},
		},
	}
	span.SetAttributes(attribute.Key("resp").String(cast.ToString(result)))
	return result, nil
}
