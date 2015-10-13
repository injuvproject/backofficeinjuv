package main

import (
	"github.com/injuvproject/backofficeinjuv/handlers"
	"github.com/injuvproject/backofficeinjuv/middleware/renderer"
	"github.com/zenazn/goji"
)

func main() {
	goji.Use(renderer.InjectRender)

	goji.Get("/", handlers.GetHome)
	goji.Serve()
}
