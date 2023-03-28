package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
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

func NewHttpUsecase(repo Repo, logger log.Logger) *HttpUsecase {
	return &HttpUsecase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "biz", "NewHttpUsecase")),
	}
}

func NewGrpcUsecase(repo Repo, logger log.Logger) *GrpcUsecase {
	return &GrpcUsecase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "biz", "NewGrpcUsecase")),
	}
}

// data层共享
type Repo interface {
	//Save Update FindByID ListAll
	GetUserInfo(context.Context, *User) ([]*User, error)
}
