package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	v1 "github.com/xiaohubai/go-grpc-layout/api/http/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/data/model"
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

// data层共享  入参 返参只接收表model和其他条件
type Repo interface {
	ListAllUser(context.Context, *model.User, *v1.PageRequest) ([]*model.User, int64, error)
	FirstUser(context.Context, *model.User) (*model.User, error)
	UpdateUserInfo(context.Context, *model.User) error
	UpdatePassword(context.Context, *model.User) error

	UpdateSetting(context.Context, *model.Setting) error
	GetSetting(context.Context, *model.Setting) (*model.Setting, error)

	AddRoleMenu(context.Context, *model.Menu) error
	DeleteRoleMenuByID(context.Context, *model.Menu) error
	UpdateRoleMenu(context.Context, *model.Menu) error
	ListAllMenu(context.Context) ([]*model.Menu, error)
	ListRoleMenu(context.Context, *model.Menu) ([]*model.Menu, error)

	ListRoleCasbin(context.Context, *model.CasbinRule, *v1.PageRequest) ([]*model.CasbinRule, int64, error)
	AddRoleCasbin(context.Context, *model.CasbinRule) error
	UpdateRoleCasbin(context.Context, *model.CasbinRule) error
	DeleteRoleCasbin(context.Context, *model.CasbinRule) error
}
