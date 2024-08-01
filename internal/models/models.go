package models

import "time"

type Message struct {
	UUID       int       `json:"uuid,omitempty"`
	Author     string    `json:"author"`
	Body       string    `json:"body"`
	RecievedAt time.Time `json:"recieved_at,omitempty"`
	Handled    bool      `json:"handled,omitempty"`
}

type Statistics struct {
	HandledMessages int       `json:"handled_messages,omitempty"`
	Messages        []string  `json:"messages,omitempty"`
	FirstDate       time.Time `json:"first_date"`
	SecondDate      time.Time `json:"second_date"`
}
