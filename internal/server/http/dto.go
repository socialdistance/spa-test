package internalhttp

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/socialdistance/spa-test/internal/storage"
	"time"
)

type PostsDto struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Created     string `json:"created"`
	Description string `json:"description"`
	UserID      string `json:"userId"`
}

type CommentDto struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Content  string `json:"content"`
	UserID   string `json:"userId"`
	PostID   string `json:"postId"`
}

type ErrorDto struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

type SearchDto struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type CountPosts struct {
	Count int `json:"count"`
}

type UserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PostComments struct {
	PostsDto
	Comments []storage.Comment
}

type PostCountComments struct {
	PostsDto
	Count int
}

func (p *PostsDto) GetModel() (*storage.Post, error) {
	created, err := time.Parse("2006-01-02T15:04:05Z", p.Created)
	if err != nil {
		return nil, fmt.Errorf("error: Created exprected to be 'yyyy-mm-dd hh:mm:ss', got: %s, %w", p.Created, err)
	}

	id, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, fmt.Errorf("ID exprected to be uuid, got: %s, %w", p.ID, err)
	}

	userID, err := uuid.Parse(p.UserID)
	if err != nil {
		return nil, fmt.Errorf("userID exprected to be uuid, got: %s, %w", p.UserID, err)
	}

	appPost := storage.NewPost(p.Title, created, p.Description, userID)
	appPost.ID = id

	return appPost, nil
}

func CreatePostDtoFromModel(post storage.Post) PostsDto {
	postDto := PostsDto{}
	postDto.ID = post.ID.String()
	postDto.Title = post.Title
	postDto.Created = post.Created.Format(time.RFC3339)
	postDto.Description = post.Description
	postDto.UserID = post.UserID.String()

	return postDto
}

func CreatePostWithCommentsDtoFromModel(post storage.Post) PostComments {
	postDto := PostComments{}
	postDto.ID = post.ID.String()
	postDto.Title = post.Title
	postDto.Created = post.Created.Format(time.RFC3339)
	postDto.Description = post.Description
	postDto.UserID = post.UserID.String()
	postDto.Comments = post.Comments

	return postDto
}

func CreatePostCountDtoFromModel(post storage.PostCount) PostCountComments {
	postDto := PostCountComments{}
	postDto.ID = post.ID.String()
	postDto.Title = post.Title
	postDto.Created = post.Created.Format(time.RFC3339)
	postDto.Description = post.Description
	postDto.UserID = post.UserID.String()
	postDto.Count = post.Count

	return postDto
}

func CreateCommentDtoModel(comment storage.Comment) CommentDto {
	commentDto := CommentDto{}
	commentDto.ID = comment.ID.String()
	commentDto.Username = comment.Username
	commentDto.Content = comment.Content
	commentDto.UserID = comment.UserID.String()
	commentDto.PostID = comment.PostID.String()

	return commentDto
}

func (p *CommentDto) GetModelComment() (*storage.Comment, error) {
	id, err := uuid.Parse(p.ID)
	if err != nil {
		return nil, fmt.Errorf("ID exprected to be uuid, got: %s, %w", p.ID, err)
	}

	userID, err := uuid.Parse(p.UserID)
	if err != nil {
		return nil, fmt.Errorf("userID exprected to be uuid, got: %s, %w", p.UserID, err)
	}

	postID, err := uuid.Parse(p.PostID)
	if err != nil {
		return nil, fmt.Errorf("postID exprected to be uuid, got: %s, %w", p.UserID, err)
	}

	appComment := storage.NewComment(p.Username, p.Content, userID, postID)
	appComment.ID = id

	return appComment, nil
}
