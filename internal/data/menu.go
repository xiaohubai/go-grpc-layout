package data

import (
	"context"

	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
)

func (d *dataRepo) ListAllMenu(ctx context.Context) (menuList []*model.Menu, err error) {
	db := d.data.db.Menu.WithContext(ctx)
	return db.Find()
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
	return db.Create(m)
}

func (d *dataRepo) DeleteRoleMenuByID(ctx context.Context, m *model.Menu) (err error) {
	db := d.data.db.Menu.WithContext(ctx)
	_, err = db.Where(d.data.db.Menu.ID.Eq(m.ID)).Delete()
	return
}

func (d *dataRepo) UpdateRoleMenu(ctx context.Context, m *model.Menu) (err error) {
	db := d.data.db.Menu.WithContext(ctx)
	_, err = db.Where(d.data.db.Menu.ID.Eq(m.ID)).Updates(map[string]any{
		"path":        m.Path,
		"name":        m.Name,
		"redirect":    m.Redirect,
		"component":   m.Component,
		"parentId":    m.ParentID,
		"roleIDs":     m.RoleIDs,
		"title":       m.Title,
		"icon":        m.Icon,
		"hidden":      m.Hidden,
		"keepAlive":   m.KeepAlive,
		"sort":        m.Sort,
		"update_user": m.UpdateUser,
	})
	return
}
