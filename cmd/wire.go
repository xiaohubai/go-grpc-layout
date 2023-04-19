//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package cmd

import (
	"github.com/xiaohubai/go-grpc-layout/configs"
	"github.com/xiaohubai/go-grpc-layout/internal/biz"

	"github.com/xiaohubai/go-grpc-layout/internal/data"
	"github.com/xiaohubai/go-grpc-layout/internal/server"
	"github.com/xiaohubai/go-grpc-layout/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*configs.Server, *configs.Data, *configs.Registry, *configs.Global, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
