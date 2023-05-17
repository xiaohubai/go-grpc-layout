// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/xiaohubai/go-grpc-layout/configs"
	"github.com/xiaohubai/go-grpc-layout/internal/biz"
	"github.com/xiaohubai/go-grpc-layout/internal/data"
	"github.com/xiaohubai/go-grpc-layout/internal/server"
	"github.com/xiaohubai/go-grpc-layout/internal/service"
	"github.com/xiaohubai/go-grpc-layout/pkg/consul"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(configsServer *configs.Server, configsData *configs.Data, configsConsul *configs.Consul, global *configs.Global, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(configsData, logger)
	if err != nil {
		return nil, nil, err
	}
	repo := data.NewDataRepo(dataData, logger)
	httpUsecase := biz.NewHttpUsecase(repo, logger)
	httpService := service.NewHttpService(httpUsecase, logger)
	httpServer := server.NewHTTPServer(configsServer, httpService, logger)
	grpcUsecase := biz.NewGrpcUsecase(repo, logger)
	grpcService := service.NewGrpcService(grpcUsecase, logger)
	grpcServer := server.NewGRPCServer(configsServer, grpcService, logger)
	registrar := consul.NewRegistry(configsConsul)
	app := newApp(logger, httpServer, grpcServer, registrar, global)
	return app, func() {
		cleanup()
	}, nil
}
