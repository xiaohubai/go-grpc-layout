//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/xiaohubai/go-grpc-layout/configs"
	"github.com/xiaohubai/go-grpc-layout/internal/biz"
	"github.com/xiaohubai/go-grpc-layout/pkg/serviceInfo"

	"github.com/xiaohubai/go-grpc-layout/internal/dao"
	"github.com/xiaohubai/go-grpc-layout/internal/server"
	"github.com/xiaohubai/go-grpc-layout/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*configs.Server, *configs.Dao, *configs.Registry, log.Logger, *serviceInfo.ServiceInfo) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, dao.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
