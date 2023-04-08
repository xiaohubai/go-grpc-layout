package configs

import (
	"errors"
	"flag"

	"github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/hashicorp/consul/api"
	"github.com/xiaohubai/go-grpc-layout/configs"
	"github.com/xiaohubai/go-grpc-layout/internal/consts"
)

type cmdFlags struct {
	env        string //环境:local remote
	filePath   string //local的文件地址
	remoteHost string //remote的地址
	remoteType string //remote的类型选项:consul
	remotePath string //remote的文件位置
}

// LoadConfig 读取一个本地文件或远程配置
func LoadConfig() (*configs.Configs, *configs.Registry) {
	var f cmdFlags
	//flag.StringVar(&p,"args", "defaultValue","eg:xxx") p:绑定的对象, args:-选项, defaultValue:默认值,eg:说明
	flag.StringVar(&f.env, "env", "local", "runtime environment, eg: -env remote")
	flag.StringVar(&f.filePath, "conf", "../configs/configs.yaml", "config path, eg: -conf conf.yaml")

	flag.StringVar(&f.remoteHost, "chost", "172.21.0.2:8500", "config server host, eg: -chost 172.21.0.2:8500")
	flag.StringVar(&f.remoteType, "ctype", "consul", "config server host, eg: -ctype consul")
	flag.StringVar(&f.remotePath, "cpath", "go-grpc-layout/dev/config.yaml", "config server path, eg: -cpath go-grpc-layout/dev/config.yaml")
	flag.Parse()
	c, err := newConfigProvider(f.env, f.remoteType, f.remoteHost, f.filePath, f.remotePath)
	if err != nil {
		panic(err)
	}
	if err := c.Load(); err != nil {
		panic(err)
	}
	var cc configs.Configs
	if err := c.Scan(&cc); err != nil {
		panic(err)
	}
	consts.Conf = &cc

	var cr configs.Registry
	if err := c.Scan(&cr); err != nil {
		panic(err)
	}
	return &cc, &cr
}

// newConfigProvider 创建一个配置
func newConfigProvider(env, remoteType, remoteHost, filePath, remotePath string) (config.Config, error) {
	switch env {
	case "local":
		return config.New(config.WithSource(file.NewSource(filePath))), nil
	}
	source, err := newRemoteConfigSource(remoteType, remoteHost, remotePath)
	if err != nil {
		return nil, err
	}
	return config.New(config.WithSource(source)), nil
}

// newRemoteConfigSource 创建一个远程配置源
func newRemoteConfigSource(remoteType, remoteHost, remotePath string) (config.Source, error) {
	switch remoteType {
	case "consul":
		return newConsulConfigSource(remoteHost, remotePath)
	}
	return nil, errors.New("empty remote type source")
}

// NewConsulConfigSource 创建一个远程配置源 - Consul
func newConsulConfigSource(remoteHost, remotePath string) (config.Source, error) {
	consulClient, err := api.NewClient(&api.Config{
		Address: remoteHost,
	})
	if err != nil {
		return nil, err
	}

	consulSource, err := consul.New(consulClient,
		consul.WithPath(remotePath),
	)
	if err != nil {
		return nil, err
	}

	return consulSource, nil
}

// NewRemoteConf 创建一个远程配置源
func NewRemoteConf(remoteType, remoteHost, remotePath string, conf any) (err error) {
	var source config.Source
	switch remoteType {
	case "consul":
		source, err = newConsulConfigSource(remoteHost, remotePath)
	}
	c := config.New(config.WithSource(source))
	if err := c.Load(); err != nil {
		return err
	}
	if err := c.Scan(conf); err != nil {
		return err
	}
	return err
}
