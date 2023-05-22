package biz

import (
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
	"github.com/xiaohubai/go-grpc-layout/pkg/jwt"
	"gorm.io/gorm"
)

var (
	i int
)

// SetSetting 设置layout配置
func (uc *HttpUsecase) DebugPerf(c *gin.Context, req *v1.DebugPerfRequest) (*v1.DebugPerfResponse, error) {
	claims, ok := c.Get("claims")
	if !ok {
		return nil, errors.New("token解析失败")
	}
	userInfo := claims.(*jwt.Claims)
	//查询 user表->debug_perf插入数据->debug_perf更新数据->查询debug_perf->删除debu_perf数据
	debugPerf, err := uc.repo.FirstDebugPerf(c.Request.Context(), &model.DebugPerf{
		UID: userInfo.UID,
	})
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errors.New("获取测试信息失败")
	}
	if debugPerf == nil {
		err = uc.repo.AddDebugPerf(c.Request.Context(), &model.DebugPerf{
			UID:  userInfo.UID,
			Text: req.Text,
		})
		if err != nil {
			return nil, errors.New("插入测试表数据失败")
		}
	}
	user, err := uc.repo.FirstUser(c.Request.Context(), &model.User{
		UID: userInfo.UID,
	})
	if err != nil {
		return nil, errors.New("获取用户信息失败")
	}
	i = i + 1
	motto := fmt.Sprintf("%s-%d", user.Motto, i)
	err = uc.repo.UpdateDebugPerf(c.Request.Context(), &model.DebugPerf{
		UID:        userInfo.UID,
		Username:   user.Username,
		Motto:      motto,
		CreateUser: userInfo.UserName,
		UpdateUser: userInfo.UserName,
	})
	if err != nil {
		return nil, errors.New("更新测试表信息失败")
	}
	data, err := uc.repo.FirstDebugPerf(c.Request.Context(), &model.DebugPerf{
		UID: userInfo.UID,
	})
	if err != nil {
		return nil, errors.New("获取测试信息失败")
	}

	/* 	err = uc.repo.DeleteDebugPerf(c.Request.Context(), &model.DebugPerf{
	   		UID: userInfo.UID,
	   	})
	   	if err != nil {
	   		return nil, errors.New("删除测试信息失败")
	   	}
	*/
	return &v1.DebugPerfResponse{
		ID:         data.ID,
		UID:        data.UID,
		UserName:   data.Username,
		Motto:      data.Motto,
		Text:       data.Text,
		CreateTime: data.CreateTime.Format(time.DateTime),
		CreateUser: data.CreateUser,
	}, err
}
