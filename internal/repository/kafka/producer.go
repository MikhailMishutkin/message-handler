package kafkarepo

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
	"message_handler/internal/models"
)

func (k *KafkaRepo) Producer(ctx context.Context, message *models.Message) error {

	m, err := json.Marshal(message)
	if err != nil {
		log.Println("failed to marshal messages:", err)
		return err
	}
	_, err = k.k.WriteMessages(
		kafka.Message{Value: m},
	)
	if err != nil {
		log.Fatal("failed to write messages:", err)
		return err
	}

	//if err := k.k.Close(); err != nil {
	//	log.Println("failed to close writer:", err)
	//	return err
	//
	//}

	return nil
}
