package data

import (
	"context"

	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/model"
)

func (d *dataRepo) ListAllUser(ctx context.Context, u *model.User, p *v1.PageRequest) (users []*model.User, total int64, err error) {

	return
}

func (d *dataRepo) FirstUser(ctx context.Context, u *model.User) (users *model.User, err error) {
	db := d.data.db.User.WithContext(ctx)
	if u.Username != "" {
		db = db.Where(d.data.db.User.Username.Eq(u.Username))
	}
	users, err = db.First()
	return
}
