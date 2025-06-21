package storage

import (
	"be/models"
	"context"
)

type MessageRepository interface {
	SaveMessage(ctx context.Context, msg models.Message) error
	GetMessages(ctx context.Context) ([]models.Message, error)
}
