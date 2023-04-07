package configs

import (
	"flag"

	"github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/hashicorp/consul/api"
	"github.com/xiaohubai/go-grpc-layout/configs"
)

type commandFlags struct {
	configEnv  string //环境:local remote
	configFile string //local的文件地址
	ConfigHost string //remote的地址
	ConfigType string //remote的类型选项:consul
	configPath string //remote的文件位置
}

var Cfg *configs.Configs

// LoadConfig 读取一个本地文件或远程配置
func LoadConfig() (*configs.Configs, *configs.Registry) {
	var f commandFlags
	//flag.StringVar(&p,"args", "defaultValue","eg:xxx") p:绑定的对象, args:-选项, defaultValue:默认值,eg:说明
	flag.StringVar(&f.configFile, "conf", "../configs/configs.yaml", "config path, eg: -conf conf.yaml")
	flag.StringVar(&f.configEnv, "env", "local", "runtime environment, eg: -env remote")
	flag.StringVar(&f.ConfigHost, "chost", "172.21.0.2:8500", "config server host, eg: -chost 172.21.0.2:8500")
	flag.StringVar(&f.ConfigType, "ctype", "consul", "config server host, eg: -ctype consul")
	flag.StringVar(&f.configPath, "cpath", "rpc-layout/dev/config.yaml", "config server path, eg: -cpath rpc-layout/dev/config.yaml")
	flag.Parse()
	c := newConfigProvider(f.ConfigType, f.ConfigHost, f.configEnv, f.configFile, f.configPath)
	if err := c.Load(); err != nil {
		panic(err)
	}
	var cc configs.Configs
	if err := c.Scan(&cc); err != nil {
		panic(err)
	}
	Cfg = &cc
	var cr configs.Registry
	if err := c.Scan(&cr); err != nil {
		panic(err)
	}
	return &cc, &cr
}

// newConfigProvider 创建一个配置
func newConfigProvider(configType, configHost, configEnv, configFile, configPath string) config.Config {
	switch configEnv {
	case "local":
		return config.New(config.WithSource(file.NewSource(configFile)))
	case "remote":
		return config.New(config.WithSource(newRemoteConfigSource(configType, configHost, configPath)))
	}
	return config.New(
		config.WithSource(
			file.NewSource(configFile),
			newRemoteConfigSource(configType, configHost, configPath),
		),
	)
}

// newRemoteConfigSource 创建一个远程配置源
func newRemoteConfigSource(configType, configHost, configPath string) config.Source {
	switch configType {
	case "consul":
		return newConsulConfigSource(configHost, configPath)
	}
	return nil
}

// newConsulConfigSource 创建一个远程配置源 - Consul
func newConsulConfigSource(configHost, configPath string) config.Source {
	consulClient, err := api.NewClient(&api.Config{
		Address: configHost,
	})
	if err != nil {
		panic(err)
	}

	consulSource, err := consul.New(consulClient,
		consul.WithPath(configPath),
	)
	if err != nil {
		panic(err)
	}

	return consulSource
}
