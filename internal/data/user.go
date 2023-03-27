package data

import (
	"context"

	"github.com/xiaohubai/go-grpc-layout/internal/biz"
)

func (r *dataRepo) GetUserInfo(ctx context.Context, dict *biz.User) ([]*biz.User, error) {
	d := make([]*biz.User, 0)
	return d, nil
}
