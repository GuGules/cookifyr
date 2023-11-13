package middleware

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
)

type CtxKey int

const (
	Database CtxKey = iota
)

func InjectDatabase(db *sqlx.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), Database, db)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
