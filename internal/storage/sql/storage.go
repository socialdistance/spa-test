package sqlstorage

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/socialdistance/spa-test/internal/storage"
	"os"
)

type Storage struct {
	ctx  context.Context
	conn *pgx.Conn
	url  string
}

func New(ctx context.Context, url string) *Storage {
	return &Storage{
		ctx: ctx,
		url: url,
	}
}

func (s *Storage) Connect(ctx context.Context) error {
	conn, err := pgx.Connect(ctx, s.url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect database %s", err)
		os.Exit(1)
	}

	s.conn = conn

	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	return s.conn.Close(ctx)
}

func (s *Storage) Create(e storage.Post) error {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) Update(e storage.Post) error {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) Delete(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (s *Storage) FindAll() ([]storage.Post, error) {
	//TODO implement me
	panic("implement me")
}
