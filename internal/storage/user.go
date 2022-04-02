package storage

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Username string
	Password string
	Token    string
}

//func NewUser(username, password string) *User {
//	return &User{
//		ID:       uuid.New(),
//		Username: username,
//		Password: password,
//	}
//}
