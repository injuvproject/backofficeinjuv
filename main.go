package main

import (
	"net/http"

	"github.com/injuvproject/backofficeinjuv/handlers"
	"github.com/injuvproject/backofficeinjuv/helpers/helperdb"
	"github.com/injuvproject/backofficeinjuv/middleware/database"
	"github.com/injuvproject/backofficeinjuv/middleware/renderer"
	"github.com/pressly/cji"
	"github.com/theosomefactory/goji-static"
	"github.com/zenazn/goji"
)

func main() {

	db := helperdb.GetDatabase()
	defer db.Close()

	goji.Use(renderer.InjectRender)
	goji.Use(static.Static("assets"))
	goji.Get("/", http.RedirectHandler("/login.html", http.StatusFound))
	goji.Get("/login.html", handlers.GetLogin)
	goji.Post("/login.html", cji.Use(database.InjectDatabase(db)).On(handlers.PostLogin))

	goji.Serve()
}
