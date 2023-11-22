package holmes

import (
	"context"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"mosn.io/holmes"
	mlog "mosn.io/pkg/log"

	"github.com/xiaohubai/go-grpc-layout/configs/conf"
	"github.com/xiaohubai/go-grpc-layout/internal/consts"
	"github.com/xiaohubai/go-grpc-layout/pkg/email"
)

// NewRegisterHolmes 异常捕获
func RegisterHolmes(c *conf.Holmes) error {
	h, err := holmes.New(
		holmes.WithCollectInterval(c.CollectInterval),
		holmes.WithProfileReporter(&ReporterImpl{}),
		holmes.WithDumpPath(c.Path),
		holmes.WithTextDump(),
		holmes.WithLogger(holmes.NewFileLog(c.Log, mlog.ERROR)),

		//现场占用率,突增,总占用率,两次dump操作之间最小时间间隔
		holmes.WithCPUDump(int(c.CPU.Min), int(c.CPU.Diff), int(c.CPU.Abs), viper.GetDuration(c.CPU.CoolDown)),
		holmes.WithMemDump(int(c.Mem.Min), int(c.Mem.Diff), int(c.Mem.Abs), viper.GetDuration(c.Mem.CoolDown)),
		holmes.WithGCHeapDump(int(c.GCHeap.Min), int(c.GCHeap.Diff), int(c.GCHeap.Abs), viper.GetDuration(c.GCHeap.CoolDown)),
		holmes.WithGoroutineDump(int(c.Goroutine.Min), int(c.Goroutine.Diff), int(c.Goroutine.Abs), int(c.Goroutine.Max), viper.GetDuration(c.Goroutine.CoolDown)),
	)
	if err != nil {
		return err
	}

	h.EnableCPUDump()
	h.EnableMemDump()
	h.EnableGCHeapDump()
	h.EnableGoroutineDump()
	h.Start()

	return nil
}

type ReporterImpl struct{}

func (r *ReporterImpl) Report(pType string, filename string, reason holmes.ReasonType, eventID string, sampleTime time.Time, pprofBytes []byte, scene holmes.Scene) error {

	msg := fmt.Sprintf("pType:%s filename:%s", pType, filename)
	filePath := fmt.Sprintf("%s/%s", consts.PwdPath, filename)
	email.SendWarnWithFile(context.Background(), consts.EmailTitlePprof, filePath, msg)
	return nil
}
