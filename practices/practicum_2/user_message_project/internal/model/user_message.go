package model

import (
	"time"
)

// UserMessage – структура, описывающая сообщение пользователя
type UserMessage struct {
	UserId    int    `json:"user_id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}
