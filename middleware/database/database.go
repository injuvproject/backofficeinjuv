package database

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/zenazn/goji/web"
)

func InjectDatabase(db *sqlx.DB) func(c *web.C, h http.Handler) http.Handler {
	return func(c *web.C, h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			c.Env["mysql"] = db
			h.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
