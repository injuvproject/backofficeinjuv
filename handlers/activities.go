package handlers

import (
	"net/http"

	"github.com/injuvproject/backofficeinjuv/helpers/security"
	"github.com/injuvproject/backofficeinjuv/models/binding"
	"github.com/unrolled/render"
	"github.com/zenazn/goji/web"
)

func GetActivities(c web.C, w http.ResponseWriter, r *http.Request) {
	template := c.Env["render"].(*render.Render)
	cookie, _ := r.Cookie("injuv_auth")
	claims, _ := security.Decode(cookie.Value)
	bnd := binding.GetDefault(r)
	bnd["ID"] = claims["id"]
	bnd["ADMIN"] = claims["guuid"]
	template.HTML(w, http.StatusOK, "panel/activities", bnd, render.HTMLOptions{
		Layout: "panel/layout",
	})
}
