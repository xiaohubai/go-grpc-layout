package kafka

import (
	"context"

	"github.com/Shopify/sarama"

	"github.com/xiaohubai/go-grpc-layout/configs/conf"
	"github.com/xiaohubai/go-grpc-layout/internal/consts"
	"github.com/xiaohubai/go-grpc-layout/pkg/email"
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

func (Consumer) Setup(_ sarama.ConsumerGroupSession) error { return nil }

func (Consumer) Cleanup(_ sarama.ConsumerGroupSession) error { return nil }

func (c Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		if err := newWorkerByTopic(msg.Topic).Run(context.Background(), msg); err != nil {
			email.SendWarn(context.Background(), consts.Conf.Email, consts.EmailTitleKafkaConsumer, err.Error())
		}
		session.MarkMessage(msg, "")
	}
	return nil
}
