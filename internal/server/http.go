package server

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/viper"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/xiaohubai/go-grpc-layout/configs"
	"github.com/xiaohubai/go-grpc-layout/internal/service"
	m "github.com/xiaohubai/go-grpc-layout/pkg/middleware"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *configs.Server, s *service.HttpService, lg log.Logger) *http.Server {
	var opts = []http.ServerOption{}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != "" {
		opts = append(opts, http.Timeout(viper.GetDuration(c.Http.Timeout)))
	}

	srv := http.NewServer(opts...)
	srv.HandlePrefix("/", routers(s))
	return srv
}

func routers(s *service.HttpService) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(m.Recovery(), m.Tracing("go-grpc-layout"), m.Metrics("go-grpc-layout"))
	r := router.Group("")
	{
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
		r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}
	r1 := r.Group("")
	{
		r1.POST("/v1/login", s.Login)    //登录
		r1.GET("/v1/captcha", s.Captcha) //获取验证码
	}
	r2 := r.Group("").Use(m.Jwt(), m.Casbin())
	{
		r2.GET("/v1/get/dictList", s.GetDictList)         //获取字典序
		r2.POST("/v1/get/allMenuList", s.GetAllMenuList)  //获取全部路由菜单
		r2.GET("/v1/get/roleMenuList", s.GetRoleMenuList) //获取角色路由菜单
	}

	return router
}
