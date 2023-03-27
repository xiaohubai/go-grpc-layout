package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "github.com/xiaohubai/go-grpc-layout/api/admin/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewService)

type Service struct {
	pb.UnimplementedAdminServer
	log *log.Helper
	uc  *biz.Usecase
}

func NewService(uc *biz.Usecase, lg log.Logger) *Service {
	return &Service{
		uc:  uc,
		log: log.NewHelper(log.With(lg, "service", "NewService")),
	}
}
