package biz

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/data/model"

	"github.com/xiaohubai/go-grpc-layout/pkg/jwt"
)

// DebugPerf 性能测试
func (uc *HttpUsecase) DebugPerf(c *gin.Context, req *v1.DebugPerfRequest) (*v1.DebugPerfResponse, error) {
	claims, ok := c.Get("claims")
	if !ok {
		return nil, errors.New("token解析失败")
	}
	userInfo := claims.(*jwt.Claims)
	/* 	lockKey := fmt.Sprintf("%s:lock:%s", consts.Conf.Global.Env, "debugPerf")
	   	l := redis.NewRedisLock(consts.RDB, lockKey, "debugPerf", 3*time.Second)
	   	if ok := l.Lock(c.Request.Context()); !ok {
	   		return nil, errors.New("资源已被持有,稍后再试")
	   	}
	   	defer func() {
	   		l.UnLock(c.Request.Context(), lockKey, "debugPerf")
	   	}()
	   	err := uc.repo.TransactionDebugPerf(c.Request.Context(), &model.DebugPerf{
	   		UID:        userInfo.UID,
	   		Text:       req.Text,
	   		CreateUser: userInfo.UserName,
	   		UpdateUser: userInfo.UserName,
	   	})
	   	if err != nil {
	   		return nil, err
	   	} */
	res, err := uc.repo.FirstDebugPerf(c.Request.Context(), &model.DebugPerf{
		UID: userInfo.UID,
	})
	if err != nil {
		return nil, err
	}

	return &v1.DebugPerfResponse{
		ID:         res.ID,
		UID:        res.UID,
		UserName:   res.Username,
		Motto:      res.Motto,
		Text:       res.Text,
		CreateTime: res.CreateAt.Format(time.DateTime),
		CreateUser: res.CreateUser,
	}, nil
}
