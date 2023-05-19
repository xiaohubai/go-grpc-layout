package kafka

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/xiaohubai/go-grpc-layout/configs/conf"
	"github.com/xiaohubai/go-grpc-layout/pkg/metric"
)

func RegisterConsumer(nodes []*conf.Kafka_Consumer) error {
	err := NewConsumerWorker(nodes)
	if err == nil {
		go NewConsumerGroup(nodes)
	}
	return err
}

func NewConsumerGroup(nodes []*conf.Kafka_Consumer) {
	var address, topics []string
	for _, v := range nodes {
		address = append(address, v.Address...)
		topics = append(topics, v.Topic)
	}

	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	group, err := sarama.NewConsumerGroup(address, "group", config)
	if err != nil {
		panic(err)
	}

	for {
		err := group.Consume(context.Background(), topics, Consumer{})
		if err != nil {
			panic(err)
		}
	}
}

type Consumer struct{}

func (Consumer) Setup(_ sarama.ConsumerGroupSession) error   { return nil }
func (Consumer) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }
func (c Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		if err := consumerMap[msg.Topic].Run(context.Background(), msg); err != nil {
			metric.Count.With(fmt.Sprintf("%s_consumer_error", msg.Topic)).Inc()
		}
		session.MarkMessage(msg, "")
	}
	return nil
}
