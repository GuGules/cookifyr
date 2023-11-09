package main

import (
	"log"
	"net/http"

	"github.com/GuGules/cookifyr/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
	Addr   string
	Router *chi.Mux
}

func NewServer(addr string) *Server {
	return &Server{
		Addr:   addr,
		Router: chi.NewRouter(),
	}
}

func (s *Server) Serve() error {
	log.Printf("server will listen on %s", s.Addr)
	return http.ListenAndServe(s.Addr, s.Router)
}

func configureServer(s *Server) {
	s.Router.Use(middleware.Logger)
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
}

func main() {
	var config, errConfig = config.NewFromEnv()

	if errConfig != nil {
		log.Fatalf("cookifyr: %s", errConfig)
	}

	s := NewServer(config.HttpAddress)
	configureServer(s)
	s.Serve()
}
