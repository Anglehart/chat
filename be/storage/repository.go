package storage

import "context"

type MessageRepo interface {
	SaveMessage(ctx context.Context, msg string) error
}
