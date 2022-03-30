package storage

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
}

type Post struct {
	ID          uuid.UUID
	Title       string
	Created     time.Time
	Description string
	UserID      uuid.UUID
}
