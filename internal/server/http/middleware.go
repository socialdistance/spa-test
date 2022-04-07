package internalhttp

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/socialdistance/spa-test/internal/auth"
	"net/http"
	"strings"
)

type ResposeWriter struct {
	http.ResponseWriter
	StatusCode int
	Bytes      int
}

func (w *ResposeWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *ResposeWriter) Write(bytes []byte) (int, error) {
	r, err := w.ResponseWriter.Write(bytes)
	w.Bytes += r

	return r, err
}

func loggingMiddleware(next http.Handler, logger Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		wrt := &ResposeWriter{w, 0, 0}
		next.ServeHTTP(wrt, r)
		logger.LogHTTP(r, wrt.StatusCode, wrt.Bytes)
	})
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers:", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
		return
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Missing Authorization Header"))
			return
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := auth.VerifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Error verifying JWT token: " + err.Error()))
			return
		}
		name := claims.(jwt.MapClaims)["name"].(string)

		r.Header.Set("name", name)

		next.ServeHTTP(w, r)
	})
}
