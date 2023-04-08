package biz

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	pb "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/consts"
	"github.com/xiaohubai/go-grpc-layout/internal/model"
	pJwt "github.com/xiaohubai/go-grpc-layout/pkg/jwt"
	"github.com/xiaohubai/go-grpc-layout/pkg/tracing"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils"
	"go.opentelemetry.io/otel/attribute"
)

// Login 用户登录
func (uc *HttpUsecase) Login(ctx context.Context, u *model.User) (*pb.LoginResponse, error) {
	ctx, span := tracing.NewSpan(ctx, "biz-Login")
	defer span.End()

	userInfo, err := uc.repo.FirstUser(ctx, &model.User{Username: u.Username})
	if err != nil {
		return nil, err
	}

	if userInfo.Password != utils.Md5([]byte(u.Password+userInfo.Salt)) {
		return nil, errors.New("密码错误")
	}

	token, err := pJwt.Create(model.Claims{
		UID:        userInfo.UID,
		UserName:   userInfo.Username,
		Phone:      userInfo.Phone,
		RoleID:     userInfo.RoleID,
		RoleName:   userInfo.RoleName,
		State:      int(userInfo.State),
		BufferTime: int64(consts.Conf.Jwt.BufferTime),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                               // 签名生效时间
			ExpiresAt: time.Now().Unix() + int64(consts.Conf.Jwt.ExpiresTime), // 过期时间
			Issuer:    consts.Conf.Jwt.Issuer,                                 // 签名的发行者
		},
	})
	if err != nil {
		return nil, errors.New("生成token失败")
	}

	span.SetAttributes(attribute.Key("userList").String(cast.ToString(userInfo)))
	result := &pb.LoginResponse{
		UserName:     userInfo.Username,
		NickName:     userInfo.Nickname,
		Birth:        userInfo.Birth.Local().Format("2006-01-02"),
		Token:        token,
		RefreshToken: token,
	}
	return result, nil
}
