package biz

import (
	"errors"

	"github.com/gin-gonic/gin"
	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
	"github.com/xiaohubai/go-grpc-layout/pkg/jwt"
)

// GetCasbinList 获取权限列表
func (uc *HttpUsecase) GetSetting(c *gin.Context) (*v1.SettingResponse, error) {
	claims, ok := c.Get("claims")
	if !ok {
		return nil, errors.New("token解析失败")
	}
	userInfo := claims.(*jwt.Claims)
	setting, err := uc.repo.GetSetting(c.Request.Context(), &model.Setting{UID: userInfo.UID})
	if err != nil {
		return nil, err
	}
	res := &v1.SettingResponse{
		ID:                    setting.ID,
		UID:                   setting.UID,
		Lang:                  setting.Lang,
		SideModeColor:         setting.SideModeColor,
		Collapse:              setting.Collapse,
		Breadcrumb:            setting.Breadcrumb,
		DefaultRouter:         setting.DefaultRouter,
		ActiveTextColor:       setting.ActiveTextColor,
		ActiveBackgroundColor: setting.ActiveBackgroundColor,
	}
	return res, nil
}

// SetSetting 设置layout配置
func (uc *HttpUsecase) UpdateSetting(c *gin.Context, s *v1.SettingRequest) error {
	claims, ok := c.Get("claims")
	if !ok {
		return errors.New("token解析失败")
	}
	userInfo := claims.(*jwt.Claims)

	return uc.repo.UpdateSetting(c.Request.Context(), &model.Setting{
		UID:                   userInfo.UID,
		Lang:                  s.Lang,
		SideModeColor:         s.SideModeColor,
		Collapse:              s.Collapse,
		Breadcrumb:            s.Breadcrumb,
		DefaultRouter:         s.DefaultRouter,
		ActiveTextColor:       s.ActiveTextColor,
		ActiveBackgroundColor: s.ActiveBackgroundColor,
		UpdateUser:            userInfo.UserName,
	})
}
