package biz

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/xiaohubai/go-grpc-layout/pkg/kafka"
)

type TestRecord struct {
	kafka.WorkerHandler
}

func (h *TestRecord) Do(c context.Context, msg *sarama.ConsumerMessage) (err error) {
	fmt.Println("TestRecord链式操作")
	fmt.Printf("Message Value:%s,Message topic:%q partition:%d offset:%d\n", string(msg.Value), msg.Topic, msg.Partition, msg.Offset)
	return
}

func init() {
	kafka.Register("TestRecord", &TestRecord{})
}
