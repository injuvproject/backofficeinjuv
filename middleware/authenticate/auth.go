package authenticate

import (
	"fmt"
	"net/http"

	"github.com/injuvproject/backofficeinjuv/helpers/getresponse"
	"github.com/injuvproject/backofficeinjuv/helpers/security"
	"github.com/zenazn/goji/web"
)

var (
	formatURLlogin = fmt.Sprintf("/login.html?false=%s", getresponse.LogInErrEmalOrPassword)
)

func InjectAuthenticate(c *web.C, h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		cookie, err := r.Cookie("injuv_auth")

		if err != nil {
			http.Redirect(w, r, formatURLlogin, http.StatusFound)
			return
		}

		claims, ok := security.Decode(cookie.Value)

		if !ok {
			http.Redirect(w, r, formatURLlogin, http.StatusFound)
			return
		}

		for key, value := range claims {

			c.Env[key] = value
		}

		h.ServeHTTP(w, r)
		return
	}
	return http.HandlerFunc(fn)
}
