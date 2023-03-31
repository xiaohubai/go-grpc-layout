package dao

import (
	"context"

	"github.com/xiaohubai/go-grpc-layout/internal/model"
)

func (r *dataRepo) ListAllUser(ctx context.Context, dict *model.User) ([]*model.User, error) {
	d := make([]*model.User, 0)
	return d, nil
}
