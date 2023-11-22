//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package cmd

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/xiaohubai/go-grpc-layout/configs/conf"
	"github.com/xiaohubai/go-grpc-layout/internal/biz"
	"github.com/xiaohubai/go-grpc-layout/internal/data"
	"github.com/xiaohubai/go-grpc-layout/internal/server"
	"github.com/xiaohubai/go-grpc-layout/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Consul, *conf.Global, log.Logger) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
