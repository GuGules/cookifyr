package util

import (
	"net/http"

	"github.com/GuGules/cookifyr/internal/apperror"
	custommiddleware "github.com/GuGules/cookifyr/internal/middleware"
	"github.com/jmoiron/sqlx"
)

func GetDbFromRequest(r *http.Request) (*sqlx.DB, error) {
	db, ok := r.Context().Value(custommiddleware.Database).(*sqlx.DB)

	if !ok || db == nil {
		return nil, apperror.NewInternalServerError()
	}

	return db, nil
}
