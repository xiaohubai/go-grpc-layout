package server

import (
	"github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"

	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/xiaohubai/go-grpc-layout/configs"
	m "github.com/xiaohubai/go-grpc-layout/pkg/metrics"

	gpb "github.com/xiaohubai/go-grpc-layout/api/grpc/v1"

	"github.com/xiaohubai/go-grpc-layout/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *configs.Server, g *service.GrpcService, lg log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			logging.Server(lg),
			tracing.Server(),
			recovery.Recovery(),
			validate.Validator(),
			metrics.Server(
				metrics.WithSeconds(prometheus.NewHistogram(m.MetricSeconds)),
				metrics.WithRequests(prometheus.NewCounter(m.MetricRequests)),
			),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	gpb.RegisterGrpcServer(srv, g)
	return srv
}
