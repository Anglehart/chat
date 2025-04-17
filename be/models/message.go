package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	GUID        uuid.UUID
	CreatedDate time.Time
	Text        string
}
