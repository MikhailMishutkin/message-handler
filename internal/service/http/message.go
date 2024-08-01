package httpservice

import (
	"context"
	"message_handler/internal/models"
	"time"
)

func (s *MessageService) MessageService(ctx context.Context, message *models.Message) error {
	recieved_at := time.Now()
	authorId, err := s.mr.SaveAuthorToDB(ctx, message.Author)
	if err != nil {
		return err
	}

	message.RecievedAt = recieved_at
	messageToKafka, err := s.mr.SaveMessageToDB(ctx, message, authorId)
	if err != nil {
		return err
	}

	messageToKafka.Author = message.Author
	messageToKafka.Body = message.Body
	messageToKafka.RecievedAt = recieved_at
	err = s.km.Producer(ctx, messageToKafka)
	if err != nil {
		return err
	}

	messageToKafka.Handled = true
	err = s.mr.MessageHandled(ctx, messageToKafka)
	if err != nil {
		return err
	}

	return err
}
func (s *MessageService) StatisticsService(ctx context.Context, statistics *models.Statistics) (stat *models.Statistics, err error) {

	stat, err = s.mr.GetStatisticsFromDB(ctx, statistics)
	if err != nil {
		return nil, err
	}
	amount, err := s.mr.GetAmountFromDB(ctx, statistics)
	if err != nil {
		return nil, err
	}
	stat.HandledMessages = amount.HandledMessages
	stat.FirstDate = statistics.FirstDate
	stat.SecondDate = statistics.SecondDate

	return stat, err
}
