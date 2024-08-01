package kafkarepo

import "github.com/segmentio/kafka-go"

type KafkaRepo struct {
	k *kafka.Conn
}

func NewKafkaRepo(k *kafka.Conn) *KafkaRepo {
	return &KafkaRepo{k: k}
}
