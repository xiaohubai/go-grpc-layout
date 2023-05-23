package biz

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/data/model"

	"github.com/xiaohubai/go-grpc-layout/pkg/jwt"
)

// DebugPerf 性能测试(查询 user表->debug_perf插入数据->debug_perf更新数据->查询debug_perf->删除debu_perf数据)
func (uc *HttpUsecase) DebugPerf(c *gin.Context, req *v1.DebugPerfRequest) (*v1.DebugPerfResponse, error) {
	claims, ok := c.Get("claims")
	if !ok {
		return nil, errors.New("token解析失败")
	} 
	userInfo := claims.(*jwt.Claims)

	err := uc.repo.TransactionDebugPerf(c.Request.Context(), &model.DebugPerf{
		UID:        userInfo.UID,
		Text:       req.Text,
		CreateUser: userInfo.UserName,
		UpdateUser: userInfo.UserName,
	})
	if err != nil {
		return nil, err
	}
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
		CreateTime: res.CreateTime.Format(time.DateTime),
		CreateUser: res.CreateUser,
	}, nil
}
