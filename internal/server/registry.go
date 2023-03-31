package server

import (
	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	consulAPI "github.com/hashicorp/consul/api"
	"github.com/xiaohubai/go-grpc-layout/configs"
)

func NewConsulRegistry(configs *configs.Registry) registry.Registrar {
	c := consulAPI.DefaultConfig()
	c.Address = configs.Consul.Address
	c.Scheme = configs.Consul.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(configs.Consul.HealthCheck))
	return r
}
