package biz

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"

	"github.com/xiaohubai/go-grpc-layout/internal/consts"
	"github.com/xiaohubai/go-grpc-layout/pkg/kafka"
	"github.com/xiaohubai/go-grpc-layout/pkg/metric"
)

type OperationRecordReportES struct {
	kafka.WorkerHandler
}

func (h *OperationRecordReportES) Do(ctx context.Context, msg *sarama.ConsumerMessage) (err error) {
	var data map[string]interface{}
	_ = json.Unmarshal(msg.Value, &data)
	err = repoUsecase.repo.ESInsertDoc(ctx, consts.ESIndexOperationRecord, data)
	if err != nil {
		fmt.Println(err.Error())
		metric.Count.With(fmt.Sprintf("consumer_%s_to_es_error", msg.Topic)).Inc()
		return err
	}
	return
}

func init() {
	kafka.Register("OperationRecordReportES", &OperationRecordReportES{})
}
