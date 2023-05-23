package data

import (
	"context"
	"errors"

	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
	"golang.org/x/sync/errgroup"
	"gorm.io/gorm"
)

func (d *dataRepo) FirstDebugPerf(ctx context.Context, debugPerf *model.DebugPerf) (
	res *model.DebugPerf, err error) {
	db := d.data.db.DebugPerf.WithContext(ctx)
	if debugPerf.UID != "" {
		db = db.Where(d.data.db.DebugPerf.UID.Eq(debugPerf.UID))
	}
	return db.First()
}

func (d *dataRepo) AddDebugPerf(ctx context.Context, debugPerf *model.DebugPerf) (err error) {
	db := d.data.db.DebugPerf.WithContext(ctx)
	return db.Create(debugPerf)
}

func (d *dataRepo) DeleteDebugPerf(ctx context.Context, debugPerf *model.DebugPerf) (err error) {
	db := d.data.db.DebugPerf.WithContext(ctx)
	_, err = db.Where(d.data.db.DebugPerf.UID.Eq(debugPerf.UID)).Unscoped().Delete()
	return
}

func (d *dataRepo) UpdateDebugPerf(ctx context.Context, debugPerf *model.DebugPerf) (err error) {
	db := d.data.db.DebugPerf.WithContext(ctx)
	_, err = db.Omit(d.data.db.DebugPerf.ID).Where(d.data.db.DebugPerf.UID.Eq(debugPerf.UID)).Updates(debugPerf)
	return
}

func (d *dataRepo) TransactionDebugPerf(ctx context.Context, debugPerf *model.DebugPerf) (err error) {
	tx := d.data.db.Begin()
	var g errgroup.Group
	g.Go(func() error {
		data, err := tx.DebugPerf.WithContext(ctx).Where(d.data.db.DebugPerf.UID.Eq(debugPerf.UID)).First()
		if err != nil && err != gorm.ErrRecordNotFound {
			_ = tx.Rollback()
			return errors.New("获取测试信息失败")
		}
		if data == nil {
			err = tx.DebugPerf.WithContext(ctx).Create(&model.DebugPerf{
				UID:  debugPerf.UID,
				Text: debugPerf.Text,
			})
			if err != nil {
				_ = tx.Rollback()
				return errors.New("插入测试表数据失败")
			}
		}
		user, err := tx.User.WithContext(ctx).Where(d.data.db.User.UID.Eq(debugPerf.UID)).First()
		if err != nil {
			_ = tx.Rollback()
			return errors.New("获取用户信息失败")
		}
		_, err = tx.DebugPerf.WithContext(ctx).Omit(d.data.db.DebugPerf.ID).Where(d.data.db.DebugPerf.UID.Eq(debugPerf.UID)).Updates(&model.DebugPerf{
			UID:        debugPerf.UID,
			Username:   user.Username,
			Motto:      user.Motto,
			Text:       debugPerf.Text,
			CreateUser: debugPerf.CreateUser,
			UpdateUser: debugPerf.CreateUser,
		})
		if err != nil {
			_ = tx.Rollback()
			return errors.New("更新测试表信息失败")
		}
		return nil

	})
	if err := g.Wait(); err != nil {
		return err
	}
	_ = tx.Commit()
	return nil
}
