package internalhttp

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/socialdistance/spa-test/internal/app"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ServerHandlers struct {
	app *app.App
}

func NewServerHandlers(a *app.App) *ServerHandlers {
	return &ServerHandlers{app: a}
}

func (s *ServerHandlers) HelloWorld(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello, world!\n")
	w.WriteHeader(200)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(msg)
}

func (s *ServerHandlers) SelectedPost(w http.ResponseWriter, r *http.Request) {
	var dto PostComments
	err := ParseRequest(r, &dto)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
		return
	}

	post, err := dto.GetModel()
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
		return
	}

	selectedPost, err := s.app.SelectPostApp(r.Context(), *post)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err)
		return
	}

	responseData, err := json.Marshal(selectedPost)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseData)
}

func (s *ServerHandlers) CreatePost(w http.ResponseWriter, r *http.Request) {
	var dto PostsDto
	err := ParseRequest(r, &dto)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
		return
	}

	post, err := dto.GetModel()
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
		return
	}

	err = s.app.CreatePost(r.Context(), *post)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err)
		return
	}

	responseData, _ := json.Marshal(dto)
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseData)
}

func (s *ServerHandlers) UpdatePost(w http.ResponseWriter, r *http.Request) {
	var dto PostsDto
	err := ParseRequest(r, &dto)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
		return
	}

	vars := mux.Vars(r)
	dto.ID = vars["id"]

	posts, err := dto.GetModel()
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
		return
	}

	err = s.app.UpdatePost(r.Context(), *posts)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err)
		return
	}

	responseData, _ := json.Marshal(dto)
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseData)
}

func (s *ServerHandlers) DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err)
	}

	err = s.app.DeletePost(r.Context(), id)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusNoContent)
}

func (s *ServerHandlers) ListPost(w http.ResponseWriter, r *http.Request) {
	var dto CountPosts
	posts, err := s.app.GetPosts(r.Context())
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
		return
	}

	dto.Count = posts

	res, err := json.Marshal(dto)

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (s *ServerHandlers) PaginationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])
	posts, err := s.app.PaginationPosts(r.Context(), page)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
		return
	}

	postsDto := make([]PostCountComments, 0, len(posts))
	for _, t := range posts {
		postsDto = append(postsDto, CreatePostCountDtoFromModel(t))
	}

	response, err := json.Marshal(postsDto)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *ServerHandlers) SearchHandler(w http.ResponseWriter, r *http.Request) {
	var dto SearchDto
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
	}

	err = json.Unmarshal(data, &dto)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
	}

	posts, err := s.app.SearchApp(r.Context(), dto.Title, dto.Description)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
		return
	}

	postsDto := make([]PostsDto, 0, len(posts))
	for _, t := range posts {
		postsDto = append(postsDto, CreatePostDtoFromModel(t))
	}

	response, err := json.Marshal(postsDto)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *ServerHandlers) CreateComment(w http.ResponseWriter, r *http.Request) {
	var dto CommentDto
	err := ParseRequest(r, &dto)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
		return
	}

	comment, err := dto.GetModelComment()
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
		return
	}

	err = s.app.CreateCommentApp(r.Context(), *comment)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err)
		return
	}

	responseData, _ := json.Marshal(dto)
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseData)
}

func (s *ServerHandlers) UpdateComment(w http.ResponseWriter, r *http.Request) {
	var dto CommentDto
	err := ParseRequest(r, &dto)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
		return
	}

	vars := mux.Vars(r)
	dto.ID = vars["id"]

	comment, err := dto.GetModelComment()
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
		return
	}

	err = s.app.UpdateCommentApp(r.Context(), *comment)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err)
		return
	}

	responseData, _ := json.Marshal(dto)
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
	w.Write(responseData)
}

func (s *ServerHandlers) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err)
	}

	err = s.app.DeleteCommentApp(r.Context(), id)
	fmt.Println(err)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusNoContent)
}

func (s *ServerHandlers) LoginUser(w http.ResponseWriter, r *http.Request) {
	var dto UserDto

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
	}

	err = json.Unmarshal(data, &dto)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
	}

	user, err := s.app.Authorize(r.Context(), dto.Username, dto.Password)
	if err != nil {
		RespondError(w, http.StatusBadRequest, err)
		return
	}

	response, err := json.Marshal(user)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func ParseRequest(r *http.Request, dto interface{}) error {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("failed to read body: %w", err)
	}

	err = json.Unmarshal(data, dto)
	if err != nil {
		return fmt.Errorf("failed to decode JSON request: %w", err)
	}

	return nil
}

func RespondError(w http.ResponseWriter, code int, err error) {
	data, err := json.Marshal(ErrorDto{
		false,
		err.Error(),
	})
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Failed to marshall error dto"))
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(code)
	w.Write(data)
}
