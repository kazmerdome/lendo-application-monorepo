package server

import "net/http"

type server struct {
	httpPort string
	handler  http.Handler
}

func NewServer(httpPort string, handler http.Handler) *server {
	return &server{
		httpPort: httpPort,
		handler:  handler,
	}
}

func (r *server) ListenAndServe() error {
	err := http.ListenAndServe(":"+r.httpPort, r.handler)
	if err != nil {
		return err
	}
	return nil
}
