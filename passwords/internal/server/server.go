package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
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
	s.srv.Handler = router

	return s
}

func (s *Server) Start() error {
	s.srv.Addr = ":8080"

	return s.srv.ListenAndServe()
}

func (s *Server) Close() error {
	return s.srv.Close()
}
