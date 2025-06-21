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

func (s *PostgresStorage) GetMessages(ctx context.Context) ([]models.Message, error) {
	rows, err := s.db.QueryContext(ctx,
		`SELECT guid, created_date, message 
         FROM messages 
         ORDER BY created_date DESC 
         LIMIT 100`,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query messages: %w", err)
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var msg models.Message
		if err := rows.Scan(&msg.GUID, &msg.CreatedDate, &msg.Text); err != nil {
			return nil, fmt.Errorf("failed to scan message: %w", err)
		}
		messages = append(messages, msg)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	// Возвращаем сообщения в правильном порядке (новые в конце)
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	return messages, nil
}
