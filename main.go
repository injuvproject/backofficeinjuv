package main

import (
	"net/http"

	"github.com/injuvproject/backofficeinjuv/handlers"
	"github.com/injuvproject/backofficeinjuv/middleware/renderer"
	"github.com/theosomefactory/goji-static"
	"github.com/zenazn/goji"
)

func main() {
	goji.Use(renderer.InjectRender)
	goji.Use(static.Static("assets"))
	goji.Get("/", http.RedirectHandler("/login.html", http.StatusFound))
	goji.Get("/login.html", handlers.GetLogin)

	goji.Serve()
}
