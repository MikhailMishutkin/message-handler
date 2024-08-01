package httpservice

import (
	"context"
	"message_handler/internal/models"
)

type MessageService struct {
	mr MessageRepositorier
	km KafkaMessenger
}

func NewMessageService(mr MessageRepositorier, km KafkaMessenger) *MessageService {
	return &MessageService{
		mr: mr,
		km: km,
	}
}

type MessageRepositorier interface {
	SaveAuthorToDB(context.Context, string) (int, error)
	SaveMessageToDB(context.Context, *models.Message, int) (*models.Message, error)
	MessageHandled(context.Context, *models.Message) error
	GetAmountFromDB(context.Context, *models.Statistics) (*models.Statistics, error)
	GetStatisticsFromDB(context.Context, *models.Statistics) (*models.Statistics, error)
}

type KafkaMessenger interface {
	Producer(context.Context, *models.Message) error
}
