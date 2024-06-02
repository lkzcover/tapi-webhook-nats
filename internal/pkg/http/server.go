package http

import (
	"github.com/nats-io/nats.go"
	"net/http"
)

type Server struct {
	router *http.ServeMux
	nc     *nats.Conn

	ncSubj    string
	debugMode bool
}

func NewServer(nc *nats.Conn, ncSubj string, debugMode bool) *Server {
	s := &Server{router: http.NewServeMux(), nc: nc, ncSubj: ncSubj, debugMode: debugMode}

	//s.router.HandleFunc("/livness", func())
	s.router.HandleFunc("/", s.webhook)

	return s
}

func (s *Server) Run(port string, cert, key string) error {
	return http.ListenAndServeTLS(":"+port, cert, key, s.router)
}
