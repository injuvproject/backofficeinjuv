package handlers

import (
	"net/http"

	"github.com/injuvproject/backofficeinjuv/helpers/security"
	"github.com/injuvproject/backofficeinjuv/models/binding"
	"github.com/unrolled/render"
	"github.com/zenazn/goji/web"
)

func GetProfile(c web.C, w http.ResponseWriter, r *http.Request) {
	template := c.Env["render"].(*render.Render)
	bnd := binding.Binding{}

	cookie, _ := r.Cookie("injuv_auth")
	claims, _ := security.Decode(cookie.Value)

	bnd = binding.Binding{

		"PageTitle":  "Back office INJUV",
		"CurrentURL": r.URL.Path,
		"Name":       claims["name"],
	}

	template.HTML(w, http.StatusOK, "panel/profile", bnd, render.HTMLOptions{
		Layout: "panel/layout",
	})

}
