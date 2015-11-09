package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/injuvproject/backofficeinjuv/config"
	"github.com/injuvproject/backofficeinjuv/helpers/getresponse"
	"github.com/injuvproject/backofficeinjuv/helpers/security"
	"github.com/injuvproject/backofficeinjuv/helpers/utils"
	"github.com/injuvproject/backofficeinjuv/models/binding"
	"github.com/injuvproject/backofficeinjuv/models/user"
	"github.com/jmoiron/sqlx"
	"github.com/unrolled/render"
	"github.com/zenazn/goji/web"
)

func GetLogin(c web.C, w http.ResponseWriter, r *http.Request) {
	template := c.Env["render"].(*render.Render)
	bnd := binding.GetDefault(r)

	cookie, _ := r.Cookie("injuv_auth")

	if cookie != nil{
		claims, _ := security.Decode(cookie.Value)
		id := int(claims["id"].(float64))
		http.Redirect(w, r, fmt.Sprintf("/panel/%d.html", id), http.StatusFound)
		return

		
	}


	template.HTML(w, http.StatusOK, "home/login", bnd)
}

func PostLogin(c web.C, w http.ResponseWriter, r *http.Request) {
	db := c.Env["mysql"].(*sqlx.DB)

	logInEmail := utils.GetAndTrim(r, "email")
	logInPassword := utils.GetAndTrim(r, "password")

	fmt.Printf("%s\n\n", logInEmail)

	newUser := &user.User{
		Email: logInEmail,
	}

	exist := newUser.EmailExist(db)
	if exist == false {
		http.Redirect(w, r, fmt.Sprintf("/login.html?false=%s", responseget.LogInErrEmalOrPassword), http.StatusFound)
		return
	}

	err := newUser.LoadID(db)
	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/login.html?false=%s", responseget.LogInErrEmalOrPassword), http.StatusFound)
		return
	}

	u, err := user.Get(db, newUser.ID)
	if err != nil {
		panic(err)
	}

	if logInEmail == u.Email && logInPassword == u.Password {

		claims := map[string]interface{}{
			config.ConstID:    u.ID,
			config.ConstName:  u.UserName,
			config.ConstAdmin: u.Admin,
		}

		// Esto tiene el contenido de la cookie
		contenidoCookie := security.Encode(claims)

		// Expiración del próximo año
		fechaExpiracion := time.Now().AddDate(0, 0, 1)

		// Creo la cookie
		cookie := http.Cookie{
			Name:       "injuv_auth",
			Value:      contenidoCookie,
			Path:       "/",
			Expires:    fechaExpiracion,
			RawExpires: fechaExpiracion.Format(time.UnixDate),
		}

		http.SetCookie(w, &cookie)
		http.Redirect(w, r, fmt.Sprintf("/panel/%d.html", newUser.ID), http.StatusFound)
		return
	}
	http.Redirect(w, r, fmt.Sprintf("/login.html?false=%s", responseget.LogInErrEmalOrPassword), http.StatusFound)
	return

}

func Logout(c web.C, w http.ResponseWriter, r *http.Request) {
	// Expiración del próximo año
	fechaExpiracion := time.Now().AddDate(0, 0, -1)

	// Creo la cookie
	cookie := http.Cookie{
		Name:       "injuv_auth",
		Value:      "",
		Path:       "/",
		Expires:    fechaExpiracion,
		RawExpires: fechaExpiracion.Format(time.UnixDate),
	}

	fmt.Printf("%s\n\n", cookie)
	fmt.Printf("%s\n\n", fechaExpiracion)

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/login.html", http.StatusFound)
	return
}
