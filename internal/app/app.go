package app

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/socialdistance/spa-test/internal/auth"
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
	Find(id uuid.UUID) (*storage.Post, error)
	FindAll() (int, error)
	Search(title, description string) ([]storage.Post, error)
	Pagination(limit, offset int) ([]storage.PostCount, error)
	CreateComment(c storage.Comment) error
	UpdateComment(c storage.Comment) error
	DeleteComment(id uuid.UUID) error
	FindAccount(username string) (*storage.User, error)
}

func New(logger Logger, storage Storage) *App {
	return &App{
		Logger:  logger,
		Storage: storage,
	}
}

func (a *App) SelectPostApp(ctx context.Context, postID string) (*storage.Post, error) {
	a.Logger.Debug("App.SelectPost %s", postID)

	postIdString, err := uuid.Parse(postID)
	if err != nil {
		a.Logger.Error("App.SelectPost error: find existing post error: %s", err)
	}

	selectedPost, err := a.Storage.Find(postIdString)
	if err != nil {
		a.Logger.Error("App.SelectPost error: find existing post error: %s", err)
		return nil, err
	}

	return selectedPost, nil
}

func (a *App) CreatePost(ctx context.Context, post storage.Post) error {
	var existingPost *storage.Post
	var err error

	a.Logger.Debug("App.CreatePost %s", post.ID)

	existingPost, err = a.Storage.Find(post.ID)
	if err != nil {
		a.Logger.Error("App.CreatePost error: find existing post error: %s", err)
		return err
	}

	if existingPost != nil {
		a.Logger.Warn("App.CreatePost error: post with ID %s already exists", post.ID)
		return fmt.Errorf("post with ID %s already exists", post.ID)
	}

	err = a.Storage.Create(post)
	if err != nil {
		a.Logger.Error("App.CreatePost error %s", err)
		return err
	}

	return nil
}

func (a *App) UpdatePost(ctx context.Context, post storage.Post) error {
	var existingPost *storage.Post
	var err error

	a.Logger.Debug("App.UpdatePost %s", post.ID)

	if existingPost, err = a.Storage.Find(post.ID); err != nil {
		a.Logger.Error("App.UpdatePost error: find existing post error: %s", err)
		return err
	}

	if existingPost == nil {
		a.Logger.Warn("App.UpdatePost error: post with ID %s not found", post.ID)
		return fmt.Errorf("post with ID %s already exists", post.ID)
	}

	if err = a.Storage.Update(post); err != nil {
		a.Logger.Error("App.UpdatePost error %s", err)
		return err
	}

	return nil
}

func (a *App) DeletePost(ctx context.Context, id uuid.UUID) error {
	var existingPost *storage.Post
	var err error

	a.Logger.Debug("App.DeletePost %s", id)

	if existingPost, err = a.Storage.Find(id); err != nil {
		a.Logger.Error("App.DeletePost error: find existing post error: %s", err)
		return err
	}

	if existingPost == nil {
		a.Logger.Warn("App.DeletePost error: post with ID %s already exists", id)
		return fmt.Errorf("post with ID %s already exists", id)
	}

	if err = a.Storage.Delete(id); err != nil {
		a.Logger.Error("App.DeletePost error %s", err)
		return err
	}

	return nil
}

func (a *App) PaginationPosts(ctx context.Context, page int) ([]storage.PostCount, error) {
	limit := 9
	offset := limit * (page - 1)

	posts, err := a.Storage.Pagination(limit, offset)
	if err != nil {
		a.Logger.Error("App.Pagination error %s", err)
		return nil, err
	}

	return posts, nil
}

func (a *App) SearchApp(ctx context.Context, title, description string) ([]storage.Post, error) {
	posts, err := a.Storage.Search(title, description)
	if err != nil {
		a.Logger.Error("App.Pagination error %s", err)
		return nil, err
	}

	return posts, nil
}

func (a *App) GetPosts(ctx context.Context) (int, error) {
	return a.Storage.FindAll()
}

func (a *App) CreateCommentApp(ctx context.Context, comment storage.Comment) error {
	a.Logger.Debug("App.CreateComment %s", comment.ID)
	err := a.Storage.CreateComment(comment)
	if err != nil {
		a.Logger.Error("App.CreatePost error %s", err)
		return err
	}

	return nil
}

func (a *App) UpdateCommentApp(ctx context.Context, comment storage.Comment) error {
	a.Logger.Debug("App.UpdateComment %s", comment.ID)

	if err := a.Storage.UpdateComment(comment); err != nil {
		a.Logger.Error("App.UpdatePost error %s", err)
		return err
	}

	return nil
}

func (a *App) DeleteCommentApp(ctx context.Context, id uuid.UUID) error {
	a.Logger.Debug("App.DeleteComment %s", id)

	if err := a.Storage.DeleteComment(id); err != nil {
		a.Logger.Error("App.DeleteComment error %s", err)
		return err
	}

	return nil
}

func (a *App) Authorize(ctx context.Context, username, password string) (*storage.User, error) {
	a.Logger.Debug("App.AuthorizeUser %s", username)

	exitingUser, err := a.Storage.FindAccount(username)
	if err != nil {
		a.Logger.Error("App.Authorize error: find existing user error: %s", err)
		return nil, err
	}

	//if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(user.Password)); err != nil {
	//return c.JSON(http.StatusNotFound, domain.HTTPError{Error: "incorrect username or password"})
	//}

	token, err := auth.GetToken(username)
	exitingUser.Token = token

	return exitingUser, nil
}
