package consul

import (
	"bytes"
	"context"
	"errors"
	"strings"
	"sync"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	ggrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"github.com/xiaohubai/go-grpc-layout/configs"
	"google.golang.org/grpc"
)

var (
	consulClient *api.Client
	once         sync.Once
)

func NewConsulClient(remoteHost, remoteToken string) *api.Client {
	once.Do(func() {
		cli, err := api.NewClient(&api.Config{
			Address: remoteHost,
			Token:   remoteToken,
		})
		if err != nil {
			panic(err)
		}
		consulClient = cli
	})
	return consulClient
}

// NewConsulConfigSource 创建一个远程配置源 - Consul
func NewConsulConfigSource(remoteHost, remotePath, remoteToken string, conf any) error {
	NewConsulClient(remoteHost, remoteToken)
	v, err := GetConsulKV(remotePath, conf)
	if err == nil {
		watcher(v, remotePath, conf)
	}
	return err
}

func NewRegistry(cul *configs.Consul) registry.Registrar {
	cli := NewConsulClient(cul.Host, cul.Token)
	r := consul.New(cli, consul.WithHealthCheck(cul.HealthCheck))
	return r
}

func NewDiscovery(cul *configs.Consul) (*grpc.ClientConn, error) {
	cli := NewConsulClient(cul.Host, cul.Token)
	r := consul.New(cli, consul.WithHealthCheck(cul.HealthCheck))
	return ggrpc.DialInsecure(context.Background(), ggrpc.WithEndpoint(cul.Discovery.GoGrpcLayout), ggrpc.WithDiscovery(r))
}

// GetConsulKV
func GetConsulKV(remotePath string, conf any) (*viper.Viper, error) {
	kv, _, err := consulClient.KV().Get(remotePath, nil)
	if err != nil {
		return nil, errors.New("consul获取配置失败")
	}
	v := viper.New()
	confType := strings.TrimSpace(remotePath[strings.LastIndex(remotePath, ".")+1:])
	v.SetConfigType(confType)
	v.SetConfigFile(remotePath)
	err = v.ReadConfig(bytes.NewBuffer(kv.Value))
	if err != nil {
		return nil, errors.New("Viper解析配置失败")
	}
	if err := v.Unmarshal(conf); err != nil {
		return nil, err
	}
	return v, nil
}
