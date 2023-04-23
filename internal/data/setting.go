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
	_, err := db.Updates(map[string]interface{}{
		"lang":                    s.Lang,
		"side_mode_color":         s.SideModeColor,
		"collapse":                s.Collapse,
		"breadcrumb":              s.Breadcrumb,
		"default_router":          s.DefaultRouter,
		"active_text_color":       s.ActiveTextColor,
		"active_background_color": s.ActiveBackgroundColor,
		"update_user":             s.UpdateUser,
	})
	return err
}

func (d *dataRepo) GetSetting(ctx context.Context, s *model.Setting) (setting *model.Setting, err error) {
	db := d.data.db.Setting.WithContext(ctx)
	if s.UID != "" {
		db = db.Where(d.data.db.Setting.UID.Eq(s.UID))
	}
	return db.First()
}
