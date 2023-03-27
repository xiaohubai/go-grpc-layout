package server

import (
	"github.com/go-kratos/kratos/contrib/metrics/prometheus/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metrics"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/xiaohubai/go-grpc-layout/internal/conf"

	pb "github.com/xiaohubai/go-grpc-layout/api/admin/v1"
	"github.com/xiaohubai/go-grpc-layout/internal/service"
	prom "github.com/xiaohubai/go-grpc-layout/pkg/prometheus"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, GRPC *service.Service, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			tracing.Server(),
			logging.Server(logger),
			metrics.Server(
				metrics.WithSeconds(prometheus.NewHistogram(prom.MetricSeconds)),
				metrics.WithRequests(prometheus.NewCounter(prom.MetricRequests)),
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
	pb.RegisterAdminServer(srv, GRPC)
	return srv
}
