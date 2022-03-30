package app

import (
	"github.com/google/uuid"
	"github.com/socialdistance/spa-test/internal/storage"
)

type App struct {
	Logger  Logger
	Storage Storage
}

type Logger interface {
	Debug(message string, params ...interface{})
	Info(message string, params ...interface{})
	Error(message string, params ...interface{})
	Warn(message string, params ...interface{})
}

type Storage interface {
	Create(e storage.Post) error
	Update(e storage.Post) error
	Delete(id uuid.UUID) error
	FindAll() ([]storage.Post, error)
}

func New(logger Logger, storage Storage) *App {
	return &App{
		Logger:  logger,
		Storage: storage,
	}
}
