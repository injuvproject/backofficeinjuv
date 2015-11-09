package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/injuvproject/backofficeinjuv/handlers"
	"github.com/injuvproject/backofficeinjuv/helpers/helperdb"
	"github.com/injuvproject/backofficeinjuv/helpers/security"
	"github.com/injuvproject/backofficeinjuv/middleware/authenticate"
	"github.com/injuvproject/backofficeinjuv/middleware/database"
	"github.com/injuvproject/backofficeinjuv/middleware/loaduser"
	"github.com/injuvproject/backofficeinjuv/middleware/renderer"
	"github.com/injuvproject/backofficeinjuv/models/binding"

	"github.com/pressly/cji"
	"github.com/theosomefactory/goji-static"
	"github.com/unrolled/render"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func main() {

	db := helperdb.GetDatabase()
	defer db.Close()

	goji.Use(renderer.InjectRender)
	goji.Use(static.Static("assets"))
	goji.Get("/", http.RedirectHandler("/login.html", http.StatusFound))
	goji.Get("/login.html", handlers.GetLogin)
	goji.Post("/login.html", cji.Use(database.InjectDatabase(db)).On(handlers.PostLogin))
	goji.Get("/panel/:id.html", cji.Use(database.InjectDatabase(db), authenticate.InjectAuthenticate, loaduser.LoadUser).On(handlers.GetProfile))

	goji.Get("/test", cji.Use(authenticate.InjectAuthenticate).On(func(c web.C, w http.ResponseWriter, r *http.Request) {

		renderer := c.Env["render"].(*render.Render)

		/*render.JSON(w, http.StatusOK, map[string]interface{}{

		})*/
		bnd := binding.Binding{}

		cookie, _ := r.Cookie("injuv_auth")
		claims, _ := security.Decode(cookie.Value)

		bnd = binding.Binding{

			"PageTitle":  "Back office INJUV",
			"CurrentURL": r.URL.Path,
			"Name":       claims["name"],
		}

		renderer.HTML(w, http.StatusOK, "panel/profile", bnd, render.HTMLOptions{
			Layout: "panel/layout",
		})

	}))
	goji.Serve()
}
