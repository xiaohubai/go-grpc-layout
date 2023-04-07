package dao

import (
	"context"

	"github.com/xiaohubai/go-grpc-layout/internal/model"
)

func (d *dataRepo) ListAllUser(ctx context.Context, u *model.User, p *model.PageInfo) (users []*model.User, err error) {

	return
}

func (d *dataRepo) FirstUser(ctx context.Context, u *model.User) (users *model.User, err error) {
	q := d.dao.db.User.WithContext(ctx)
	if u.Username != "" {
		q = q.Where(d.dao.db.User.Username.Eq(u.Username))
	}
	users, err = q.First()
	return
}
