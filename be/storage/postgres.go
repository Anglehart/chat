package storage

import (
	"be/models"
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

// NewPostgresStorage теперь принимает контекст и строку подключения
func NewPostgresStorage(ctx context.Context, connStr string) (*PostgresStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения: %w", err)
	}

	// Проверяем подключение
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ошибка ping: %w", err)
	}

	return &PostgresStorage{db: db}, nil
}

func (s *PostgresStorage) Close() error {
	return s.db.Close()
}

func (s *PostgresStorage) SaveMessage(ctx context.Context, msg models.Message) error {
	_, err := s.db.ExecContext(ctx,
		`INSERT INTO messages (guid, created_date, message) 
        VALUES ($1, $2, $3)`,
		msg.GUID,
		msg.CreatedDate,
		msg.Text,
	)
	return err
}
