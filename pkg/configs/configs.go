package configs

import (
	"errors"
	"flag"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
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
func Load() (*configs.Configs, error) {
	var f cmdFlags
	//flag.StringVar(&p,"args", "defaultValue","eg:xxx") p:绑定的对象, args:-选项, defaultValue:默认值,eg:说明
	flag.StringVar(&f.env, "env", "local", "runtime environment, eg: -env remote")
	flag.StringVar(&f.filePath, "conf", "configs/configs.yaml", "config path, eg: -conf configs.yaml")

	flag.StringVar(&f.remoteHost, "chost", "172.21.0.2:8500", "config server host, eg: -chost 172.21.0.2:8500")
	flag.StringVar(&f.remoteType, "ctype", "consul", "config server host, eg: -ctype consul")
	flag.StringVar(&f.remotePath, "cpath", "dev/config.yaml", "config server path, eg: -cpath dev/config.yaml")
	flag.Parse()

	var cc configs.Configs
	if f.env == "local" {
		if err := newFileConfig(f.filePath, &cc); err != nil {
			return nil, err
		}
	} else {
		if err := newRemoteConfigSource(f.remoteType, f.remoteHost, f.remotePath, &cc); err != nil {
			return nil, err
		}
	}
	consts.Cfg = &cc
	return &cc, nil
}

func newFileConfig(filePath string, conf any) error {
	v := viper.New()
	v.SetConfigFile(filePath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	v.WatchConfig()
	if err := v.Unmarshal(conf); err != nil {
		return err
	}
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.Unmarshal(conf); err != nil {
			log.Errorw(err)
		}
	})

	return nil
}

// newRemoteConfigSource 创建一个远程配置源
func newRemoteConfigSource(remoteType, remoteHost, remotePath string, conf any) error {
	switch remoteType {
	case "consul":
		return NewConsulConfigSource(remoteHost, remotePath, conf)
	}
	return errors.New("empty remote type source")
}

// NewConsulConfigSource 创建一个远程配置源 - Consul
func NewConsulConfigSource(remoteHost, remotePath string, conf any) error {
	v := viper.New()
	confType := strings.TrimSpace(remotePath[strings.LastIndex(remotePath, ".")+1:])
	v.AddRemoteProvider("consul", remoteHost, remotePath)
	v.SetConfigType(confType)
	if err := v.ReadRemoteConfig(); err != nil {
		return err
	}
	if err := v.Unmarshal(conf); err != nil {
		return err
	}
	go func() {
		for {
			time.Sleep(time.Second * 5)
			err := v.WatchRemoteConfig()
			if err != nil {
				log.Errorw(err)
				continue
			}
			v.Unmarshal(conf)
		}
	}()
	return nil
}
