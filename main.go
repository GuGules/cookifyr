package main

import (
	"log"
	"net/http"

	"github.com/GuGules/cookifyr/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

const defaultAddr = ":3000"

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

func main() {
	godotenv.Load()

	addr := utils.GetEnvOr("HTTP_ADDRESS", defaultAddr)

	s := NewServer(addr)
	s.Router.Use(middleware.Logger)
	s.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	s.Serve()
}
