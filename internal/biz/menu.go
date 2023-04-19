package biz

import (
	"context"

	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
)

// GetAllMenuList 获取全部路由列表
func (uc *HttpUsecase) GetAllMenuList(ctx context.Context, p *v1.PageRequest) (*v1.PageResponse, error) {
	menuList, total, err := uc.repo.ListAllMenu(ctx, p)
	if err != nil {
		return nil, err
	}
	list := menuTreeHandler(menuList)
	res := &v1.PageResponse{
		Page:     p.Page,
		PageSize: p.PageSize,
		Total:    total,
		List:     list,
	}
	return res, nil
}

// GetRoleMenuList 获取角色路由列表
func (uc *HttpUsecase) GetRoleMenuList(ctx context.Context, m *model.Menu) ([]*v1.MenuResponse, error) {
	menuList, err := uc.repo.ListRoleMenu(ctx, m)
	if err != nil {
		return nil, err
	}
	return menuTreeHandler(menuList), nil
}

func menuTreeHandler(menuList []*model.Menu) []*v1.MenuResponse {
	m := make(map[int]*v1.MenuResponse, 0)
	res := make([]*v1.MenuResponse, 0)

	for _, v := range menuList {
		data := v1.MenuResponse{
			Path:      v.Path,
			Name:      v.Name,
			Component: v.Component,
			Redirect:  v.Redirect,
			Meta: &v1.MenuResponse_Meta{
				ID:        v.ID,
				ParentID:  v.ParentID,
				RoleIDs:   v.RoleIDs,
				Title:     v.Title,
				Icon:      v.Icon,
				Hidden:    v.Hidden,
				KeepAlive: v.KeepAlive,
				Sort:      v.Sort,
			},
		}
		if v.ParentID == 0 {
			res = append(res, &data)
		} else {
			m[int(v.ParentID)].Children = append(m[int(v.ParentID)].Children, &data)
		}
		m[int(v.ID)] = &data
	}
	return res
}
