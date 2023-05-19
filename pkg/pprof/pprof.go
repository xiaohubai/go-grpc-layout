package pprof

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/xiaohubai/go-grpc-layout/pkg/email"
	"github.com/xiaohubai/go-grpc-layout/pkg/kafka"
	"golang.org/x/sync/errgroup"
)

func RegisterPprof() error {

	return nil
}

func Report() {
	var g errgroup.Group
	g.Go(func() error {
		record := "pprof 上报"
		producer, err := kafka.NewProducer("pprof")
		if err != nil {
			return err
		}
		return producer.Send(record)
	})
	if err := g.Wait(); err != nil {
		email.SendWarn(context.Background(), err.Error())
		log.Errorw("key", "warn", "msg", err.Error())
	}
}
