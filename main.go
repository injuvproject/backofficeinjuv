package main

import (
	"fmt"
	"net/http"
	"os/user"

	_ "github.com/go-sql-driver/mysql"
	"github.com/injuvproject/backofficeinjuv/handlers"
	"github.com/injuvproject/backofficeinjuv/helpers/helperdb"
	"github.com/injuvproject/backofficeinjuv/middleware/database"
	"github.com/injuvproject/backofficeinjuv/middleware/renderer"
	"github.com/jmoiron/sqlx"
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
	goji.Get("/test", cji.Use(database.InjectDatabase(db)).On(func(c web.C, w http.ResponseWriter, r *http.Request) {
		render := c.Env["render"].(*render.Render)
		db := c.Env["mysql"].(*sqlx.DB)
		newuser := user.User{}

		err := db.Get(&newuser, "SELECT * FROM Usuario WHERE id=1", "newuser")
		if err != nil {
			fmt.Printf("error %#v\n")
		}
		fmt.Printf("%#v\n", newuser)

		render.JSON(w, http.StatusOK, map[string]interface{}{
		//"lastid":      num,
		//"lastidusers": newUser.ID,
		//"user":  user,
		//"Error": fmt.Sprintf("%s", err),
		})
	}))
	goji.Serve()
}
