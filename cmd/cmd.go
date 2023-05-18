package cmd

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/xiaohubai/go-grpc-layout/configs"
	"github.com/xiaohubai/go-grpc-layout/pkg/kafka"
	"github.com/xiaohubai/go-grpc-layout/pkg/zap"

	conf "github.com/xiaohubai/go-grpc-layout/pkg/configs"
	"github.com/xiaohubai/go-grpc-layout/pkg/tracing"

	_ "go.uber.org/automaxprocs"
)

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, rr registry.Registrar, g *configs.Global) *kratos.App {
	return kratos.New(
		kratos.ID(g.Id),
		kratos.Name(g.AppName),
		kratos.Version(g.Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Registrar(rr),
		kratos.Server(
			hs,
			gs,
		),
	)
}

func Run() (*kratos.App, func()) {
	cc, err := conf.Load()
	if err != nil {
		panic("load config failed")
	}
	logger, err := zap.New(cc.Zap, cc.Global)
	if err != nil {
		panic("load logger failed")
	}
	if err := tracing.NewTracerProvider(cc.Trace.Endpoint, cc.Global); err != nil {
		panic("load tracing failed")
	}
	if err := kafka.Server(cc.Kafka.Node); err != nil {
		panic("load kafka failed")
	}
	app, cleanup, err := wireApp(cc.Server, cc.Data, cc.Consul, cc.Global, logger)
	if err != nil {
		panic(err)
	}
	return app, cleanup
}
