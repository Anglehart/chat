package service

import (
	"be/models"
	"be/storage"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type MessageService struct {
	storage storage.MessageRepository // Интерфейс с методом SaveMessage
}

func NewMessageService(storage storage.MessageRepository) *MessageService {
	return &MessageService{storage: storage}
}

func (s *MessageService) SaveMessage(ctx context.Context, text string) (*models.Message, error) {
	msg := models.Message{
		GUID:        uuid.New(),       // Генерация GUID
		CreatedDate: time.Now().UTC(), // Установка времени
		Text:        text,
	}

	if err := s.storage.SaveMessage(ctx, msg); err != nil {
		return nil, fmt.Errorf("не удалось сохранить: %w", err)
	}

	return &msg, nil
}
