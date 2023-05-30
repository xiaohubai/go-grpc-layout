package viper

import (
	"context"
	"errors"
	"flag"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"

	"github.com/xiaohubai/go-grpc-layout/configs/conf"
	"github.com/xiaohubai/go-grpc-layout/internal/consts"
	"github.com/xiaohubai/go-grpc-layout/pkg/consul"
	"github.com/xiaohubai/go-grpc-layout/pkg/email"
)

type cmdFlags struct {
	env         string //环境:local remote
	filePath    string //local的文件地址
	remoteHost  string //remote的地址
	remoteType  string //remote的类型选项:consul
	remotePath  string //remote的文件位置
	remoteToken string //remote的密钥
}

// Load 读取一个本地文件或远程配置
func Load() (*conf.Conf, error) {
	var f cmdFlags
	//flag.StringVar(&p,"args", "defaultValue","eg:xxx") p:绑定的对象, args:-选项, defaultValue:默认值,eg:说明
	flag.StringVar(&f.env, "env", "local", "runtime environment, eg: -env remote")
	flag.StringVar(&f.filePath, "conf", "configs/conf/conf.yaml", "config path, eg: -conf configs.yaml")

	flag.StringVar(&f.remoteHost, "chost", "172.21.0.2:8500", "config server host, eg: -chost 172.21.0.2:8500")
	flag.StringVar(&f.remoteType, "ctype", "consul", "remote config server host, eg: -ctype consul")
	flag.StringVar(&f.remotePath, "cpath", "dev/conf.yaml", "remote config server path, eg: -cpath dev/conf.yaml")
	flag.StringVar(&f.remoteToken, "ctoken", "ac9b7b85-8819-cffb-c3f6-1bbd43ca1402", "remote config server token")
	flag.Parse()

	var cc conf.Conf
	if f.env == "local" {
		if err := newFileConfig(f.filePath, conf.C); err != nil {
			return nil, err
		}
	} else {
		if err := newRemoteConfigSource(f.remoteType, f.remoteHost, f.remotePath, f.remoteToken, &cc); err != nil {
			return nil, err
		}
	}
	return conf.C, nil
}

func newFileConfig(filePath string, cf any) error {
	v := viper.New()
	confType := strings.TrimSpace(filePath[strings.LastIndex(filePath, ".")+1:])
	v.SetConfigFile(filePath)
	v.SetConfigType(confType)
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	v.WatchConfig()
	if err := v.Unmarshal(cf); err != nil {
		return err
	}
	v.OnConfigChange(func(e fsnotify.Event) {
		if err := v.Unmarshal(cf); err != nil {
			email.SendWarn(context.Background(), consts.EmailTitleViperLocalWatch, err.Error())
		}
	})

	return nil
}

// newRemoteConfigSource 创建一个远程配置源
func newRemoteConfigSource(remoteType, remoteHost, remotePath, remoteToken string, conf any) error {
	switch remoteType {
	case "consul":
		return consul.NewConsulConfigSource(remoteHost, remotePath, remoteToken, conf)
	}
	return errors.New("empty remote type source")
}
