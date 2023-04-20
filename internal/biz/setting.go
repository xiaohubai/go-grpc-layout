package biz

import (
	"context"

	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
)

// GetCasbinList 获取权限列表
func (uc *HttpUsecase) GetSetting(ctx context.Context, s *model.Setting) (*v1.SettingResponse, error) {
	setting, err := uc.repo.GetSetting(ctx, s)
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
func (uc *HttpUsecase) UpdateSetting(ctx context.Context, s *v1.SettingRequest) error {
	return uc.repo.UpdateSetting(ctx, &model.Setting{
		UID:                   s.UID,
		Lang:                  s.Lang,
		SideModeColor:         s.SideModeColor,
		Collapse:              s.Collapse,
		Breadcrumb:            s.Breadcrumb,
		DefaultRouter:         s.DefaultRouter,
		ActiveTextColor:       s.ActiveTextColor,
		ActiveBackgroundColor: s.ActiveBackgroundColor,
	})
}
