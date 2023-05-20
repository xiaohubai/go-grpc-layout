package holmes

import (
	"github.com/spf13/viper"
	"github.com/xiaohubai/go-grpc-layout/configs/conf"
	"mosn.io/holmes"
	mlog "mosn.io/pkg/log"
)

// NewRegisterHolmes 异常捕获
func NewRegisterHolmes(c *conf.Holmes) error {
	h, err := holmes.New(
		holmes.WithCollectInterval(c.CollectInterval),
		holmes.WithDumpPath(c.Path),
		holmes.WithTextDump(),
		holmes.WithLogger(holmes.NewFileLog(c.Log, mlog.INFO)),

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
