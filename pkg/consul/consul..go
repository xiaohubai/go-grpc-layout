package consul

import (
	"bytes"
	"context"
	"errors"
	"strings"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	ggrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"google.golang.org/grpc"

	"github.com/xiaohubai/go-grpc-layout/configs/conf"
)

func NewConsulClient(remoteHost, remoteToken string) (*api.Client, error) {
	cli, err := api.NewClient(&api.Config{
		Address: remoteHost,
		Token:   remoteToken,
	})
	if err != nil {
		return nil, err
	}
	return cli, err
}

// NewConsulConfigSource 创建一个远程配置源 - Consul
func NewConsulConfigSource(remoteHost, remotePath, remoteToken string, conf any) error {
	cli, err := NewConsulClient(remoteHost, remoteToken)
	if err != nil {
		return err
	}
	v, err := GetConsulKV(cli, remotePath, conf)
	if err == nil {
		watcher(cli, v, remotePath, conf)
	}
	return err
}

func NewRegistry(cul *conf.Consul) (registry.Registrar, error) {
	cli, err := NewConsulClient(cul.Host, cul.Token)
	if err != nil {
		return nil, err
	}
	r := consul.New(cli, consul.WithHealthCheck(cul.HealthCheck))
	return r, nil
}

func NewDiscovery(cul *conf.Consul, endpoint string) (*grpc.ClientConn, error) {
	cli, err := NewConsulClient(cul.Host, cul.Token)
	if err != nil {
		return nil, err
	}
	r := consul.New(cli, consul.WithHealthCheck(cul.HealthCheck))
	return ggrpc.DialInsecure(context.Background(), ggrpc.WithEndpoint(endpoint), ggrpc.WithDiscovery(r))
}

// GetConsulKV
func GetConsulKV(cli *api.Client, remotePath string, conf any) (*viper.Viper, error) {
	kv, _, err := cli.KV().Get(remotePath, nil)
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
