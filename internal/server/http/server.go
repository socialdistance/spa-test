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
