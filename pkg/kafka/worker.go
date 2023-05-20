package kafka

import (
	"context"
	"errors"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/xiaohubai/go-grpc-layout/configs/conf"
)

var (
	workers = make(map[string]Handler)
	lock    sync.RWMutex
)

func newWorkerByTopic(topic string) Handler {
	lock.RLock()
	w := workers[topic]
	lock.RUnlock()
	return w
}

type WorkerHandler struct {
	next Handler
}

type Handler interface {
	Do(ctx context.Context, msg *sarama.ConsumerMessage) error
	SetNext(h Handler) Handler
	Run(ctx context.Context, msg *sarama.ConsumerMessage) error
}

func (n *WorkerHandler) SetNext(h Handler) Handler {
	n.next = h
	return h
}

func (n *WorkerHandler) Run(ctx context.Context, msg *sarama.ConsumerMessage) (err error) {
	if n.next != nil {
		if err = (n.next).Do(ctx, msg); err != nil {
			return err
		}
		return (n.next).Run(ctx, msg)
	}
	return
}

// 注册yaml和biz对应consumer实例
var workerHandlers = make(map[string]Handler)

func Register(name string, maker Handler) {
	workerHandlers[name] = maker
}

// 首链
type NullHandler struct {
	WorkerHandler
}

func (h *NullHandler) Do(ctx context.Context, msg *sarama.ConsumerMessage) (err error) {
	return
}
func HandlerInterface() Handler {
	return &NullHandler{}
}

// NewConsumerWorker yaml配置topic对于的func调用链式实例
func NewConsumerWorker(nodes []*conf.Kafka_Consumer) error {
	for _, v := range nodes {
		head := HandlerInterface()
		temp := head
		for _, vv := range v.Func {
			w, ok := workerHandlers[vv]
			if !ok {
				return errors.New("not found")
			}
			temp = temp.SetNext(w)
		}
		workers[v.Topic] = head
	}
	return nil
}
