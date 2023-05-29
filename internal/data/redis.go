package data

import (
	"context"
	"time"

	predis "github.com/xiaohubai/go-grpc-layout/pkg/redis"
)

func (d *dataRepo) Get(ctx context.Context, key string) (string, error) {
	r := predis.New(d.data.rdb)
	return r.Get(ctx, key)
}

func (d *dataRepo) Set(ctx context.Context, key, value string) error {
	r := predis.New(d.data.rdb)
	return r.Set(ctx, key, value)
}

func (d *dataRepo) SetEx(ctx context.Context, key, value string, expiration time.Duration) error {
	r := predis.New(d.data.rdb)
	return r.SetEx(ctx, key, value, expiration)
}
