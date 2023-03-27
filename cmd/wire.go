//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/xiaohubai/go-grpc-layout/internal/biz"
	"github.com/xiaohubai/go-grpc-layout/internal/conf"
	"github.com/xiaohubai/go-grpc-layout/internal/data"
	"github.com/xiaohubai/go-grpc-layout/internal/server"
	"github.com/xiaohubai/go-grpc-layout/internal/service"
	"github.com/xiaohubai/go-grpc-layout/pkg/configs"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data,*conf.Registry, log.Logger, *configs.ServiceInfo) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
