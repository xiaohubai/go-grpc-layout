package biz

import (
	"context"

	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
)

// GetCasbinList 获取权限列表
func (uc *HttpUsecase) GetCasbinList(ctx context.Context, casbin *model.CasbinRule, p *v1.PageRequest) (
	*v1.PageResponse, error) {
	casbinList, total, err := uc.repo.ListCasbinList(ctx, casbin, p)
	if err != nil {
		return nil, err
	}
	list := make([]*v1.GetCasbinResponse, 0)
	for _, v := range casbinList {
		data := &v1.GetCasbinResponse{
			ID:      int32(v.ID),
			RoleIDs: v.V0,
			Path:    v.V1,
			Method:  v.V2,
			Desc:    v.Desc,
		}
		list = append(list, data)
	}
	res := &v1.PageResponse{
		Page:     p.Page,
		PageSize: p.PageSize,
		Total:    total,
		List:     list,
	}
	return res, nil
}
