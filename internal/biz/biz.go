package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/xiaohubai/go-grpc-layout/internal/model"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewHttpUsecase, NewGrpcUsecase)

type HttpUsecase struct {
	repo Repo
	log  *log.Helper
}

type GrpcUsecase struct {
	repo Repo
	log  *log.Helper
}

func NewHttpUsecase(repo Repo, lg log.Logger) *HttpUsecase {
	return &HttpUsecase{
		repo: repo,
		log:  log.NewHelper(lg),
	}
}

func NewGrpcUsecase(repo Repo, lg log.Logger) *GrpcUsecase {
	return &GrpcUsecase{
		repo: repo,
		log:  log.NewHelper(lg),
	}
}

// data层共享
type Repo interface {
	ListAllUser(context.Context, *model.User, *model.PageInfo) ([]*model.User, error)
	FirstUser(context.Context, *model.User) (*model.User, error)
}
