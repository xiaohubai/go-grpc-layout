package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/xiaohubai/go-grpc-layout/pkg/configs"
	"github.com/xiaohubai/go-grpc-layout/pkg/logger"
	metrics "github.com/xiaohubai/go-grpc-layout/pkg/metrics"
	"github.com/xiaohubai/go-grpc-layout/pkg/tracing"

	_ "go.uber.org/automaxprocs"
)

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, rr registry.Registrar, serviceInfo *configs.ServiceInfo) *kratos.App {
	return kratos.New(
		kratos.ID(serviceInfo.GetInstanceId()),
		kratos.Name(serviceInfo.Name),
		kratos.Version(serviceInfo.Version),
		kratos.Metadata(serviceInfo.Metadata),
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
	serviceInfo := configs.NewServiceInfo(cc.Global.AppName, cc.Global.Env, cc.Global.Version, "")

	lg := logger.New(&serviceInfo)

	prometheus.MustRegister(metrics.MetricSeconds, metrics.MetricRequests)

	if err := tracing.NewTracerProvider(cc.Trace.Endpoint, &serviceInfo); err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(cc.Server, cc.Data, cr, lg, &serviceInfo)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}
