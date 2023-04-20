package data

import (
	"context"
	"errors"

	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
)

func (d *dataRepo) UpdateSetting(ctx context.Context, s *model.Setting) error {
	db := d.data.db.Setting.WithContext(ctx)
	if s.UID == "" {
		return errors.New("UID must not be empty")
	}
	db = db.Where(d.data.db.Setting.UID.Eq(s.UID))
	_, err := db.Updates(s)
	return err
}

func (d *dataRepo) GetSetting(ctx context.Context, s *model.Setting) (setting *model.Setting, err error) {
	db := d.data.db.Setting.WithContext(ctx)
	if s.UID != "" {
		db = db.Where(d.data.db.Setting.UID.Eq(s.UID))
	}
	return db.First()
}
