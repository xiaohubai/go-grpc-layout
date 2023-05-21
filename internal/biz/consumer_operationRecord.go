package biz

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"

	"github.com/xiaohubai/go-grpc-layout/pkg/kafka"
	"github.com/xiaohubai/go-grpc-layout/pkg/metric"
)

type OperationRecordReportES struct {
	kafka.WorkerHandler
}

func (h *OperationRecordReportES) Do(ctx context.Context, msg *sarama.ConsumerMessage) (err error) {
	//fmt.Println("OperationRecord记录写入es...")
	/* fmt.Printf("Message Value:%s,Message topic:%q partition:%d offset:%d\n", string(msg.Value), msg.Topic, msg.Partition, msg.Offset) */
	// 如果 写入es失败 埋点
	metric.Count.With(fmt.Sprintf("consumer_%s_to_es_error", msg.Topic)).Inc()
	return
}

func init() {
	kafka.Register("OperationRecordReportES", &OperationRecordReportES{})
}
