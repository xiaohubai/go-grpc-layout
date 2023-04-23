package biz

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/consts"
	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
	pJwt "github.com/xiaohubai/go-grpc-layout/pkg/jwt"
	"github.com/xiaohubai/go-grpc-layout/pkg/tracing"
	"github.com/xiaohubai/go-grpc-layout/pkg/utils"
	"go.opentelemetry.io/otel/attribute"
)

// Login 用户登录
func (uc *HttpUsecase) Login(c *gin.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	ctx, span := tracing.NewSpan(c.Request.Context(), "biz-Login")
	defer span.End()

	userInfo, err := uc.repo.FirstUser(ctx, &model.User{Username: req.Username})
	if err != nil {
		return nil, err
	}

	if userInfo.Password != utils.Md5([]byte(req.Password+userInfo.Salt)) {
		return nil, errors.New("密码错误")
	}

	token, err := pJwt.Create(pJwt.Claims{
		UID:        userInfo.UID,
		UserName:   userInfo.Username,
		Phone:      userInfo.Phone,
		RoleID:     userInfo.RoleID,
		RoleName:   userInfo.RoleName,
		State:      userInfo.State,
		BufferTime: consts.Cfg.Jwt.BufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + int64(consts.Cfg.Jwt.ExpiresTime), // 过期时间
			Issuer:    consts.Cfg.Jwt.Issuer,                                 // 签名的发行者
		},
	})
	if err != nil {
		return nil, errors.New("生成token失败")
	}

	span.SetAttributes(attribute.Key("userList").String(cast.ToString(userInfo)))
	res := &v1.LoginResponse{
		ID:           userInfo.ID,
		UID:          userInfo.UID,
		UserName:     userInfo.Username,
		NickName:     userInfo.Nickname,
		Birth:        userInfo.Birth.Local().Format("2006-01-02"),
		Avatar:       userInfo.Avatar,
		RoleID:       userInfo.RoleID,
		RoleName:     userInfo.RoleName,
		Phone:        userInfo.Phone,
		Wechat:       userInfo.Wechat,
		Email:        userInfo.Email,
		State:        userInfo.State,
		Motto:        userInfo.Motto,
		Token:        token,
		RefreshToken: token,
	}
	return res, nil
}
