package loaduser

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/injuvproject/backofficeinjuv/config"
	"github.com/injuvproject/backofficeinjuv/helpers/getresponse"
	"github.com/injuvproject/backofficeinjuv/models/user"
	"github.com/jmoiron/sqlx"
	"github.com/zenazn/goji/web"
)

var (
	formatURLlogin = fmt.Sprintf("/login.html?false=%s", responseget.LogInErrEmalOrPassword)
)

func LoadUser(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		db := c.Env["mysql"].(*sqlx.DB)
		userID := c.Env[config.ConstID].(float64)
		id, _ := strconv.Atoi(c.URLParams["id"])
		u, err := user.Get(db, id)

		if err != nil {
			http.Redirect(w, r, formatURLlogin, http.StatusFound)
			return
		}

		if int(userID) != u.ID {
			http.Redirect(w, r, formatURLlogin, http.StatusFound)
			return
		}

		c.Env[config.ConstID] = user.User{
			ID:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
		}

		h.ServeHTTP(w, r)
		return
	}
	return http.HandlerFunc(fn)
}


