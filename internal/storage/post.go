package storage

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID          uuid.UUID
	Title       string
	Created     time.Time
	Description string
	UserID      uuid.UUID
	Comments    []Comment
}

type PostCount struct {
	ID          uuid.UUID
	Title       string
	Created     time.Time
	Description string
	UserID      uuid.UUID
	Count       int
}

type Comment struct {
	ID       uuid.UUID
	Username string
	Content  string
	UserID   uuid.UUID
	PostID   uuid.UUID
}

func NewPost(title string, created time.Time, description string, userID uuid.UUID) *Post {
	return &Post{
		ID:          uuid.New(),
		Title:       title,
		Created:     created,
		Description: description,
		UserID:      userID,
	}
}

func NewComment(username, content string, userID uuid.UUID, postID uuid.UUID) *Comment {
	return &Comment{
		ID:       uuid.New(),
		Username: username,
		Content:  content,
		UserID:   userID,
		PostID:   postID,
	}
}
