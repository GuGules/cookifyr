package main

import (
	"log"
	"net/http"

	"github.com/GuGules/cookifyr/internal/apperror"
	"github.com/GuGules/cookifyr/internal/config"
	custommiddleware "github.com/GuGules/cookifyr/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/go-sql-driver/mysql"
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

func configureServer(s *Server, db *sqlx.DB) {
	s.Router.Use(middleware.Logger)
	s.Router.Use(custommiddleware.InjectDatabase(db))
}

func ToHttpHandler(handlerWithError func(w http.ResponseWriter, r *http.Request) error) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		errHandle := handlerWithError(w, r)

		if errHandle != nil {
			appError, isAppError := errHandle.(apperror.AppError)
			if !isAppError {
				appError = apperror.NewInternalServerError()
			}

			http.Error(w, appError.Msg, appError.Status)
		}
	}
}

func main() {
	config, errConfig := config.NewFromEnv()

	if errConfig != nil {
		log.Fatalf("cookifyr: %s", errConfig)
	}

	db, errDb := sqlx.Connect(config.DbDriver, config.DbConnString)

	if errDb != nil {
		log.Fatalf("cookifyr: %s", errDb)
	}

	s := NewServer(config.HttpAddress)
	configureServer(s, db)
	s.Serve()
}
