package data

import (
	"context"

	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
)

func (d *dataRepo) ListCasbinList(ctx context.Context, casbin *model.CasbinRule, p *v1.PageRequest) (
	casbinList []*model.CasbinRule, total int64, err error) {
	db := d.data.db.CasbinRule.WithContext(ctx)
	total, err = db.Count()
	if err != nil {
		return
	}
	if casbin.V0 != "" {
		db = db.Where(d.data.db.CasbinRule.V0.Eq(casbin.V0))
	}
	if casbin.V1 != "" {
		db = db.Where(d.data.db.CasbinRule.V1.Eq(casbin.V1))
	}
	if casbin.V2 != "" {
		db = db.Where(d.data.db.CasbinRule.V2.Eq(casbin.V2))
	}
	casbinList, err = db.Limit(int(p.PageSize)).Offset(int(p.PageSize * (p.Page - 1))).Find()
	return
}
