package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

type key int

const (
	requestIDKey key = 0
)

type Server struct {
	srv *http.Server
}

func Get() *Server {
	return &Server{
		srv: &http.Server{},
	}
}

func (s *Server) WithAddr(host string) *Server {
	s.srv.Addr = host

	return s
}

func (s *Server) WithRouter(router *httprouter.Router) *Server {
	// FIXME: this is bad
	nextRequestID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}

	s.srv.Handler = cors.Default().Handler(tracing(nextRequestID)(logging()(router)))

	return s
}

func (s *Server) Start() error {
	s.srv.Addr = ":8080"

	return s.srv.ListenAndServe()
}

func (s *Server) Close() error {
	return s.srv.Close()
}

func logging() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				requestID, ok := r.Context().Value(requestIDKey).(string)
				if !ok {
					requestID = "unknown"
				}
				fmt.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
