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
