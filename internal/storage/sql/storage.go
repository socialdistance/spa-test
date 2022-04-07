package sqlstorage

import "C"
import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/socialdistance/spa-test/internal/storage"
	"os"
	"time"
)

type Storage struct {
	ctx  context.Context
	conn *pgxpool.Pool
	url  string
}

func New(ctx context.Context, url string) *Storage {
	return &Storage{
		ctx: ctx,
		url: url,
	}
}

func (s *Storage) Connect(ctx context.Context) error {
	//conn, err := pgx.Connect(ctx, s.url)
	conn, err := pgxpool.Connect(ctx, s.url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect database %s", err)
		os.Exit(1)
	}

	s.conn = conn

	return nil
}

func (s *Storage) Close(ctx context.Context) {
	s.conn.Close()
}

func (s *Storage) FindAccount(username string) (*storage.User, error) {
	var u storage.User
	sql := `
		select * from users where username = $1
	`

	err := s.conn.QueryRow(s.ctx, sql, username).Scan(
		&u.ID,
		&u.Username,
		&u.Password,
	)

	if err == nil {
		return &u, nil
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	return nil, fmt.Errorf("failed to scan SQL result into struct: %w", err)
}

func (s *Storage) Create(p storage.Post) error {
	sql := `
		insert into posts (id, title, created, description, user_id) values ($1, $2, $3, $4, $5)
	`

	_, err := s.conn.Exec(s.ctx, sql, p.ID.String(), p.Title, p.Created.Format(time.RFC3339), p.Description, p.UserID)

	return err
}

func (s *Storage) CreateComment(c storage.Comment) error {
	sql := `
		insert into comments (id, username, content, user_id, post_id) values ($1, $2, $3, $4, $5)
	`

	_, err := s.conn.Exec(s.ctx, sql, c.ID.String(), c.Username, c.Content, c.UserID, c.PostID)

	return err
}

func (s *Storage) Update(p storage.Post) error {
	sql := `
		update posts set title = $2, created = $3, description = $4, user_id = $5 where id = $1
	`
	_, err := s.conn.Exec(s.ctx, sql, p.ID.String(), p.Title, p.Created.Format(time.RFC3339), p.Description, p.UserID)

	return err
}

func (s *Storage) UpdateComment(c storage.Comment) error {
	sql := `
		update comments set username = $2, content = $3, user_id = $4, post_id = $5 where id = $1
	`

	_, err := s.conn.Exec(s.ctx, sql, c.ID.String(), c.Username, c.Content, c.UserID, c.PostID)

	return err
}

func (s *Storage) Delete(id uuid.UUID) error {
	sql := "delete from posts where id = $1"

	_, err := s.conn.Exec(s.ctx, sql, id)

	return err
}

func (s *Storage) DeleteComment(id uuid.UUID) error {
	sql := "delete from comments where id = $1"

	_, err := s.conn.Exec(s.ctx, sql, id)

	return err
}

func (s *Storage) Find(id uuid.UUID) (*storage.Post, error) {
	var comments []storage.Comment
	var p storage.Post

	sql := `
		select id, title, created, description, user_id from posts where id = $1
	`

	sqlComments := `
		select id, username, content, user_id, post_id from comments where post_id=$1;
	`

	//err := s.conn.QueryRow(s.ctx, sql, id).Scan(
	//	&p.ID,
	//	&p.Title,
	//	&p.Created,
	//	&p.Description,
	//	&p.UserID,
	//)

	rowsPost, err := s.conn.Query(s.ctx, sql, id)
	defer rowsPost.Close()

	for rowsPost.Next() {
		if err := rowsPost.Scan(
			&p.ID,
			&p.Title,
			&p.Created,
			&p.Description,
			&p.UserID); err != nil {
			return nil, fmt.Errorf("unable to transform array result into struct: %w", err)
		}
	}

	rows, err1 := s.conn.Query(s.ctx, sqlComments, id)
	defer rows.Close()

	for rows.Next() {
		var c storage.Comment
		if err := rows.Scan(
			&c.ID,
			&c.Username,
			&c.Content,
			&c.UserID,
			&c.PostID,
		); err1 != nil {
			return nil, fmt.Errorf("unable to transform array result into struct: %w", err)
		}

		comments = append(comments, c)
	}

	p.Comments = comments

	if err == nil {
		return &p, nil
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}

	return nil, fmt.Errorf("failed to scan SQL result into struct: %w", err)
}

func (s *Storage) FindAll() (int, error) {
	posts := make([]storage.Post, 0)

	sql := `
		select id, title, created, description, user_id from posts
	`

	rows, err := s.conn.Query(s.ctx, sql)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var p storage.Post
		if err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Created,
			&p.Description,
			&p.UserID,
		); err != nil {
			return 0, fmt.Errorf("unable to transform array result into struct: %w", err)
		}

		posts = append(posts, p)
	}

	if err := rows.Err(); err != nil {
		return 0, err
	}

	return len(posts), nil
}

func (s *Storage) Pagination(limit, offset int) ([]storage.PostCount, error) {
	posts := make([]storage.PostCount, 0)

	sql := `
	 select posts.id, title, created, description, posts.user_id, count(comments.post_id) from posts LEFT JOIN comments on posts.id=post_id group by posts.id, created, title, description, posts.user_id order by created limit $2 offset $1;
	`

	rows, err := s.conn.Query(s.ctx, sql, offset, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p storage.PostCount
		if err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Created,
			&p.Description,
			&p.UserID,
			&p.Count,
		); err != nil {
			return nil, fmt.Errorf("unable to transform array result into struct: %w", err)
		}

		posts = append(posts, p)
	}

	return posts, nil
}

func (s *Storage) Search(title, description string) ([]storage.Post, error) {
	posts := make([]storage.Post, 0)

	sql := `
		SELECT * FROM posts WHERE title ILIKE '%' || $1 || '%' AND description ILIKE '%' || $2 || '%'
	`

	//sql := `
	//	SELECT * FROM posts WHERE title LIKE $1 OR description LIKE $2
	//`

	rows, err := s.conn.Query(s.ctx, sql, title, description)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var p storage.Post
		if err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Created,
			&p.Description,
			&p.UserID,
		); err != nil {
			return nil, fmt.Errorf("unable to transform array result into struct: %w", err)
		}

		posts = append(posts, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}
