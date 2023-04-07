package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"

	"github.com/xiaohubai/go-grpc-layout/pkg/serviceInfo"
	"github.com/xiaohubai/go-grpc-layout/pkg/zap"

	"github.com/xiaohubai/go-grpc-layout/pkg/configs"
	"github.com/xiaohubai/go-grpc-layout/pkg/tracing"

	_ "go.uber.org/automaxprocs"
)

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, rr registry.Registrar, s *serviceInfo.ServiceInfo) *kratos.App {
	return kratos.New(
		kratos.ID(s.Id),
		kratos.Name(s.Name),
		kratos.Version(s.Version),
		kratos.Metadata(s.Metadata),
		kratos.Logger(logger),
		kratos.Registrar(rr),
		kratos.Server(
			hs,
			gs,
		),
	)
}

func main() {
	cc, cr := configs.LoadConfig()
	if cc == nil || cr == nil {
		panic("load config failed")
	}
	serviceInfo := serviceInfo.NewServiceInfo(cc.Global)

	lg := zap.New(cc.Zap, &serviceInfo)

	if err := tracing.NewTracerProvider(cc.Trace.Endpoint, &serviceInfo); err != nil {
		panic("load tracing failed")
	}

	app, cleanup, err := wireApp(cc.Server, cc.Dao, cr, lg, &serviceInfo)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
