package consul

import (
	"bytes"
	"context"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/api/watch"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"

	"github.com/xiaohubai/go-grpc-layout/internal/consts"
	"github.com/xiaohubai/go-grpc-layout/pkg/email"
)

func (col *Consul) watcher(vp *viper.Viper, path string, conf any) {
	time.Sleep(time.Second * 10)
	var g errgroup.Group
	g.Go(func() error {
		w, err := watch.Parse(map[string]interface{}{"type": "keyprefix", "prefix": path})
		if err != nil {
			return err
		}
		w.Handler = func(u uint64, i interface{}) {
			kv := i.(api.KVPairs)
			for _, v := range kv {
				if v.Key == path {
					err = vp.ReadConfig(bytes.NewBuffer(v.Value))
					if err != nil {
						return
					}
					if err = vp.Unmarshal(conf); err != nil {
						return
					}
				}
			}
		}
		err = w.RunWithClientAndHclog(col.client, nil)
		if err != nil {
			return err
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		email.SendWarn(context.Background(), consts.Conf.Email, "viper remote watch", err.Error())
	}
}
