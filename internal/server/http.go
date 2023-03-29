package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/xiaohubai/go-grpc-layout/internal/conf"
	"github.com/xiaohubai/go-grpc-layout/internal/service"
	m "github.com/xiaohubai/go-grpc-layout/pkg/middleware"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, s *service.HttpService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", routers(s))
	return srv
}

func routers(s *service.HttpService) *gin.Engine {
	var router = gin.Default()

	router.Use(m.Recovery(), otelgin.Middleware("router"))

	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	router.POST("/v1/login", s.Login)

	return router
}
