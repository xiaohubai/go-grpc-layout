package kafka

import (
	"errors"
	"time"

	"github.com/Shopify/sarama"
	"github.com/xiaohubai/go-grpc-layout/configs"
)

var (
	m map[string]sarama.SyncProducer
)

type Producer struct {
	Topic        string              `json:"topic"`
	SyncProducer sarama.SyncProducer `json:"map"`
}

func Server(node []*configs.Kafka_Node) error {
	m = make(map[string]sarama.SyncProducer, 0)
	for _, v := range node {
		cfg := sarama.NewConfig()
		cfg.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
		cfg.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
		cfg.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
		cfg.Producer.Retry.Max = 3                             // 重试三次
		cfg.Producer.Retry.Backoff = 100 * time.Millisecond
		producer, err := sarama.NewSyncProducer(v.Address, cfg)
		if err != nil {
			return err
		}
		m[v.Topic] = producer
	}
	return nil
}

func NewProducer(topic string) (*Producer, error) {
	k, ok := m[topic]
	if !ok {
		return nil, errors.New("NewProducer topic not found")
	}
	return &Producer{
		Topic:        topic,
		SyncProducer: k,
	}, nil
}

func (p *Producer) Send(data string) error {
	msg := &sarama.ProducerMessage{
		Topic: p.Topic,
		Value: sarama.StringEncoder(data),
	}
	_, _, err := p.SyncProducer.SendMessage(msg)
	return err
}
