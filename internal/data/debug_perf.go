package data

import (
	"context"

	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
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
