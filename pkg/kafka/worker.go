package kafka

import (
	"context"
	"errors"

	"github.com/Shopify/sarama"
	"github.com/xiaohubai/go-grpc-layout/configs/conf"
)

var (
	consumerMap = make(map[string]Handler)
)

type WorkerHandler struct {
	next Handler
}

type Handler interface {
	Do(c context.Context, msg *sarama.ConsumerMessage) error
	SetNext(h Handler) Handler
	Run(c context.Context, msg *sarama.ConsumerMessage) error
}

func (n *WorkerHandler) SetNext(h Handler) Handler {
	n.next = h
	return h
}

func (n *WorkerHandler) Run(c context.Context, msg *sarama.ConsumerMessage) (err error) {
	if n.next != nil {
		if err = (n.next).Do(c, msg); err != nil {
			return
		}
		return (n.next).Run(c, msg)
	}
	return
}

var (
	workerHandlers = make(map[string]Handler)
)

func Register(name string, maker Handler) {
	workerHandlers[name] = maker
}

type NullHandler struct {
	WorkerHandler
}

func (h *NullHandler) Do(c context.Context, msg *sarama.ConsumerMessage) (err error) {
	return
}
func HandlerInterface() Handler {
	return &NullHandler{}
}

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
		consumerMap[v.Topic] = head
	}
	return nil
}
