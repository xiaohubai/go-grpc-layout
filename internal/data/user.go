package data

import (
	"context"

	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
)

func (d *dataRepo) ListAllUser(ctx context.Context, u *model.User, p *v1.PageRequest) (users []*model.User, total int64, err error) {

	return
}

func (d *dataRepo) FirstUser(ctx context.Context, u *model.User) (users *model.User, err error) {
	db := d.data.db.User.WithContext(ctx)
	if u.Username != "" {
		db = db.Where(d.data.db.User.Username.Eq(u.Username))
	}
	if u.UID != "" {
		db = db.Where(d.data.db.User.UID.Eq(u.UID))
	}
	users, err = db.First()
	return
}

func (d *dataRepo) UpdateUserInfo(ctx context.Context, u *model.User) (err error) {
	db := d.data.db.User.WithContext(ctx)
	_, err = db.Where(d.data.db.User.UID.Eq(u.UID)).Updates(map[string]any{
		"nickname":    u.Nickname,
		"birth":       u.Birth,
		"phone":       u.Phone,
		"wechat":      u.Wechat,
		"email":       u.Email,
		"motto":       u.Motto,
		"update_user": u.UpdateUser,
	})
	return
}

func (d *dataRepo) UpdatePassword(ctx context.Context, u *model.User) (err error) {
	db := d.data.db.User.WithContext(ctx)
	_, err = db.Where(d.data.db.User.UID.Eq(u.UID)).Updates(map[string]any{
		"password":    u.Password,
		"salt":        u.Salt,
		"update_user": u.UpdateUser,
	})
	return
}
