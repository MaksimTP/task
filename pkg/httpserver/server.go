package httpserver

import (
	"net/http"
	"time"
)

type Server interface {
	Run(port string, maxHeaderBytes, readTimeout, writeTimeout int, handler http.Handler) error
}

type server struct {
	httpServer *http.Server
}

func New() *server {
	return &server{}
}

func (s *server) Run(port string, maxHeaderBytes, readTimeout, writeTimeout int, handler http.Handler) error {
	s.httpServer = &http.Server{Addr: ":" + port, Handler: handler, MaxHeaderBytes: maxHeaderBytes, ReadTimeout: time.Duration(readTimeout) * time.Second, WriteTimeout: time.Duration(writeTimeout) * time.Second}
	return s.httpServer.ListenAndServe()
}
