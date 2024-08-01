package postgres

import (
	"context"
	"fmt"
	"message_handler/internal/models"
)

// ...
func (r *Repo) SaveAuthorToDB(ctx context.Context, author string) (authorId int, err error) {
	const query = `INSERT INTO authors (name) VALUES ($1) RETURNING uuid`
	err = r.DB.QueryRow(ctx, query, author).Scan(&authorId)
	return authorId, err
}

// ...
func (r *Repo) SaveMessageToDB(ctx context.Context, message *models.Message, authorId int) (messageToKafka *models.Message, err error) {
	const query = `INSERT INTO messages (author_uuid, message, recieved_at, handled) VALUES ($1, $2, $3, $4) RETURNING uuid`
	var uuid int
	messageToKafka = &models.Message{}
	row := r.DB.QueryRow(ctx, query, authorId, message.Body, message.RecievedAt, message.Handled)
	err = row.Scan(&uuid)
	if err != nil {
		return nil, fmt.Errorf("can't insert message to DB: ", err)
	}
	messageToKafka.UUID = uuid

	return messageToKafka, nil
}

func (r *Repo) MessageHandled(ctx context.Context, message *models.Message) error {

	const query = `UPDATE messages SET handled = $1 WHERE uuid = $2`

	_, err := r.DB.Exec(ctx, query, message.Handled, message.UUID)
	if err != nil {
		return fmt.Errorf("can't upate status message in DB: ", err)
	}
	return nil
}

// ...
func (r *Repo) GetAmountFromDB(ctx context.Context, statIn *models.Statistics) (statOut *models.Statistics, err error) {
	const query = `SELECT COUNT(*) FROM messages 
                WHERE handled = true AND recieved_at >= $1::date
					AND    recieved_at   <=$2::date
                `
	var amount int
	err = r.DB.QueryRow(ctx, query, statIn.FirstDate, statIn.SecondDate).Scan(&amount)
	if err != nil {
		return nil, fmt.Errorf("can't get amount of handled messages from DB: ", err)
	}
	statOut = &models.Statistics{HandledMessages: amount}

	return statOut, err
}

// ...
func (r *Repo) GetStatisticsFromDB(ctx context.Context, statIn *models.Statistics) (statOut *models.Statistics, err error) {
	const query = `SELECT message FROM messages 
                WHERE handled = true AND recieved_at >= $1::date
					AND    recieved_at   <=$2::date
                `
	var messages []string
	rows, err := r.DB.Query(ctx, query, statIn.FirstDate, statIn.SecondDate)
	if err != nil {
		return nil, fmt.Errorf("can't get statistics from DB: ", err)
	}
	for rows.Next() {

		var message string
		if err = rows.Scan(&message); err != nil {
			return nil, fmt.Errorf("trouble with rows.Next: %s", err)
		}
		messages = append(messages, message)
	}
	statOut = &models.Statistics{Messages: messages}

	return statOut, err
}
