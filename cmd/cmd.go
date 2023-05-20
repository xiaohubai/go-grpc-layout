package cmd

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/xiaohubai/go-grpc-layout/configs/conf"
	"github.com/xiaohubai/go-grpc-layout/pkg/holmes"
	"github.com/xiaohubai/go-grpc-layout/pkg/kafka"
	"github.com/xiaohubai/go-grpc-layout/pkg/zap"

	"github.com/xiaohubai/go-grpc-layout/pkg/tracing"
	"github.com/xiaohubai/go-grpc-layout/pkg/viper"

	_ "github.com/xiaohubai/go-grpc-layout/internal/biz" //init biz
	_ "go.uber.org/automaxprocs"
)

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, rr registry.Registrar, g *conf.Global) *kratos.App {
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

func NewApp() *kratos.App {
	cc, err := viper.Load()
	if err != nil {
		panic("load config failed")
	}
	logger, err := zap.New(cc.Zap, cc.Global)
	if err != nil {
		panic("load logger failed")
	}
	if err := tracing.RegisterTracer(cc.Trace.Endpoint, cc.Global); err != nil {
		panic("load tracing failed")
	}
	if err := kafka.RegisterProducer(cc.Kafka.Producer); err != nil {
		panic("load kafka producer failed")
	}
	if err := kafka.RegisterConsumer(cc.Kafka.Consumer); err != nil {
		panic("load kafka consumer failed")
	}
	if err := holmes.NewRegisterHolmes(cc.Holmes); err != nil {
		panic("load holmes failed")
	}

	app, err := wireApp(cc.Server, cc.Data, cc.Consul, cc.Global, logger)
	if err != nil {
		panic(err)
	}
	return app
}
