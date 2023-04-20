package data

import (
	"context"

	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
)

func (d *dataRepo) ListAllMenu(ctx context.Context, p *v1.PageRequest) (menuList []*model.Menu, total int64, err error) {
	db := d.data.db.Menu.WithContext(ctx)
	total, err = db.Count()
	if err != nil {
		return
	}
	menuList, err = db.Limit(int(p.PageSize)).Offset(int(p.PageSize * (p.Page - 1))).Find()
	return
}

func (d *dataRepo) ListRoleMenu(ctx context.Context, m *model.Menu) (menuList []*model.Menu, err error) {
	db := d.data.db.Menu.WithContext(ctx)
	if m.RoleIDs != "" {
		db = db.Where(d.data.db.Menu.RoleIDs.FindInSetWith(m.RoleIDs))
	}
	menuList, err = db.Find()
	return
}

func (d *dataRepo) AddRoleMenu(ctx context.Context, m *model.Menu) (err error) {
	db := d.data.db.Menu.WithContext(ctx)
	if m.RoleIDs != "" {
		db = db.Where(d.data.db.Menu.RoleIDs.FindInSetWith(m.RoleIDs))
	}
	menuList, err = db.Find()
	return
}

func (d *dataRepo) DeleteRoleMenu(ctx context.Context, m *model.Menu) (err error) {
	db := d.data.db.Menu.WithContext(ctx)
	if m.RoleIDs != "" {
		db = db.Where(d.data.db.Menu.RoleIDs.FindInSetWith(m.RoleIDs))
	}
	menuList, err = db.Find()
	return
}

func (d *dataRepo) UpdateRoleMenu(ctx context.Context, m *model.Menu) (err error) {
	db := d.data.db.Menu.WithContext(ctx)
	if m.RoleIDs != "" {
		db = db.Where(d.data.db.Menu.RoleIDs.FindInSetWith(m.RoleIDs))
	}
	menuList, err = db.Find()
	return
}
