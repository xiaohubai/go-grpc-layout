package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewUsecase)

type Usecase struct {
	repo Repo
	log  *log.Helper
}

func NewUsecase(repo Repo, logger log.Logger) *Usecase {
	return &Usecase{
		repo: repo,
		log:  log.NewHelper(log.With(logger, "biz", "NewUsecase")),
	}
}

type Repo interface {
	//Save Update FindByID ListAll
	GetUserInfo(context.Context, *User) ([]*User, error)
}
