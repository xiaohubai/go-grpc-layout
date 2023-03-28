package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	pb "github.com/xiaohubai/go-grpc-layout/api/admin/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewHttpService, NewGrpcService)

type HttpService struct {
	pb.UnimplementedAdminServer
	log *log.Helper
	uc  *biz.HttpUsecase
}

type GrpcService struct {
	pb.UnimplementedAdminServer
	log *log.Helper
	uc  *biz.GrpcUsecase
}

func NewHttpService(uc *biz.HttpUsecase, lg log.Logger) *HttpService {
	return &HttpService{
		uc:  uc,
		log: log.NewHelper(log.With(lg, "service", "NewHttpService")),
	}
}
func NewGrpcService(uc *biz.GrpcUsecase, lg log.Logger) *GrpcService {
	return &GrpcService{
		uc:  uc,
		log: log.NewHelper(log.With(lg, "service", "NewGrpcService")),
	}
}
