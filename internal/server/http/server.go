package internalhttp

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/socialdistance/spa-test/internal/app"
	"net"
	"net/http"
)

type Server struct {
	host   string
	port   string
	logger Logger
	server *http.Server
}

type Logger interface {
	Debug(message string, params ...interface{})
	Info(message string, params ...interface{})
	Error(message string, params ...interface{})
	Warn(message string, params ...interface{})
	LogHTTP(r *http.Request, code, length int)
}

func NewServer(logger Logger, app *app.App, host, port string) *Server {
	server := &Server{
		host:   host,
		port:   port,
		logger: logger,
		server: nil,
	}

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: loggingMiddleware(NewRouter(app), logger),
	}

	server.server = httpServer

	return server
}

func NewRouter(app *app.App) http.Handler {
	handlers := NewServerHandlers(app)

	r := mux.NewRouter()

	//helloWorldFn := http.HandlerFunc(handlers.HelloWorld)
	//createPostFn := http.HandlerFunc(handlers.CreatePost)
	//updatePostFn := http.HandlerFunc(handlers.UpdatePost)
	//deletePostFn := http.HandlerFunc(handlers.DeletePost)
	//listPostFn := http.HandlerFunc(handlers.ListPost)
	//selectedPostFn := http.HandlerFunc(handlers.SelectedPost)
	//createCommentFn := http.HandlerFunc(handlers.CreateComment)
	//updateCommentFn := http.HandlerFunc(handlers.UpdateComment)
	//deleteCommentFn := http.HandlerFunc(handlers.DeleteComment)
	//
	//r.HandleFunc("/", handlers.HelloWorld).Methods("GET")
	//r.Handle("/posts/create", authMiddleware(createPostFn)).Methods("POST")
	//r.Handle("/posts/update/{id}", authMiddleware(updatePostFn)).Methods("PUT")
	//r.Handle("/posts/delete/{id}", authMiddleware(deletePostFn)).Methods("DELETE")
	//r.HandleFunc("/posts/{page}", handlers.PaginationHandler).Methods("GET")
	//r.HandleFunc("/posts/search", handlers.SearchHandler).Methods("POST")
	//r.Handle("/posts", authMiddleware(listPostFn)).Methods("GET")
	//r.Handle("/post", authMiddleware(selectedPostFn)).Methods("POST")
	//
	//r.Handle("/comments/create", authMiddleware(createCommentFn)).Methods("POST")
	//r.Handle("/comments/update/{id}", authMiddleware(updateCommentFn)).Methods("PUT")
	//r.Handle("/comments/delete/{id}", authMiddleware(deleteCommentFn)).Methods("DELETE")
	//
	//r.HandleFunc("/login", handlers.LoginUser).Methods("POST")

	r.HandleFunc("/", handlers.HelloWorld).Methods("GET")
	r.HandleFunc("/posts/create", handlers.CreatePost).Methods("POST")
	r.HandleFunc("/posts/update/{id}", handlers.UpdatePost).Methods("PUT")
	r.HandleFunc("/posts/delete/{id}", handlers.DeletePost).Methods("DELETE")
	r.HandleFunc("/posts/{page}", handlers.PaginationHandler).Methods("GET")
	r.HandleFunc("/posts/search", handlers.SearchHandler).Methods("POST")
	r.HandleFunc("/posts", handlers.ListPost).Methods("GET")
	r.HandleFunc("/post", handlers.SelectedPost).Methods("POST")

	r.HandleFunc("/comments/create", handlers.CreateComment).Methods("POST")
	r.HandleFunc("/comments/update/{id}", handlers.UpdateComment).Methods("PUT")
	r.HandleFunc("/comments/delete/{id}", handlers.DeleteComment).Methods("DELETE")
	
	r.HandleFunc("/login", handlers.LoginUser).Methods("POST")

	return r
}

func (s *Server) Start(ctx context.Context) error {
	s.logger.Info("[+] Http server start and listen %s:%s", s.host, s.port)
	err := s.server.ListenAndServe()
	if err != nil {
		return err
	}
	<-ctx.Done()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
