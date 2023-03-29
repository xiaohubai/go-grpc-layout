package biz

import (
	"context"

	"github.com/spf13/cast"
	pb "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/pkg/tracing"
	"go.opentelemetry.io/otel/attribute"
)

type User struct {
	ID       int32  `json:"id"`
	UID      string `json:"uid"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	NickName string `json:"nickName"`
	Birth    string `json:"birth"`
	Avatar   string `json:"avatar"`
	RoleID   string `json:"roleID"`
	RoleName string `json:"roleName"`
	Phone    string `json:"phone"`
	Wechat   string `json:"wechat"`
	Email    string `json:"email"`
	State    string `json:"state"`
	Motto    string `json:"motto"`
}

// Login 用户登录
func (uc *HttpUsecase) Login(ctx context.Context, u *User) (*pb.LoginResponse, error) {
	ctx, span := tracing.NewSpan(ctx, "biz-Login")
	defer span.End()

	userList, err := uc.repo.GetUserInfo(ctx, u)
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
