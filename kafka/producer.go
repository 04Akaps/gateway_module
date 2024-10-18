package kafka

import (
	"github.com/04Akaps/gateway_module/config"
	"github.com/04Akaps/gateway_module/log"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"go.uber.org/zap"
)

const (
	_allAcks = "all"
)

type Producer struct {
	cfg      config.Producer
	producer *kafka.Producer
}

//BatchSize int64  `yaml:"batch_size"`
//BatchTime int64  `yaml:"batch_time"`

func NewProducer(
	config config.Producer,
) Producer {
	url := config.URL
	id := config.ClientID
	acks := config.Acks

	if acks == "" {
		acks = _allAcks
	}

	conf := &kafka.ConfigMap{
		"bootstrap.servers": url,
		"client.id":         id,
		"acks":              acks,
	}

	producer, err := kafka.NewProducer(conf)

	if err != nil {
		log.Log.Panic("Failed to create producer", zap.Error(err))
	}

	return Producer{
		producer: producer,
		cfg:      config,
	}
}

func (p Producer) SendEvent(value []byte) {
	topic := p.cfg.Topic

	err := p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: value,
	}, nil)

	if err != nil {
		log.Log.Error("Failed to send topic",
			zap.String("topic", topic),
			zap.String("value", string(value)),
			zap.Error(err),
		)
	} else {
		log.Log.Info("Success to send topic",
			zap.String("topic", topic),
		)
	}
}
