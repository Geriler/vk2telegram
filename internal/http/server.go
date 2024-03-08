package http

import "net/http"

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handlers http.Handler) error {
	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handlers,
	}

	return s.httpServer.ListenAndServe()
}
