package consul

import (
	"bytes"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
	"github.com/spf13/viper"
)

func watcher(vp *viper.Viper, path string, conf any) {
	time.Sleep(time.Second * 10)
	w, err := watch.Parse(map[string]interface{}{"type": "keyprefix", "prefix": path})
	if err != nil {
		log.Errorw("key", "loading", "msg", fmt.Sprintf("%s:%s", "watch.Parse", err))
	}
	w.Handler = func(u uint64, i interface{}) {
		kv := i.(api.KVPairs)
		for _, v := range kv {
			if v.Key == path {
				err := vp.ReadConfig(bytes.NewBuffer(v.Value))
				if err != nil {
					log.Errorw("key", "loading", "msg", fmt.Sprintf("%s:%s", "vp.ReadConfig", err))
				}
				if err := vp.Unmarshal(conf); err != nil {
					log.Errorw("key", "loading", "msg", fmt.Sprintf("%s:%s", "vp.Unmarshal", err))
				}
			}
		}
	}
	err = w.RunWithClientAndHclog(consulClient, nil)
	if err != nil {
		log.Errorw("key", "loading", "msg", fmt.Sprintf("%s:%s", "w.RunWithClientAndHclog", err))
	}
}
