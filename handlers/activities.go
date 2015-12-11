package handlers

import (
	"net/http"

	"github.com/injuvproject/backofficeinjuv/helpers/security"
	"github.com/injuvproject/backofficeinjuv/helpers/utils"
	"github.com/injuvproject/backofficeinjuv/models/binding"
	"github.com/injuvproject/backofficeinjuv/models/user"
	"github.com/jmoiron/sqlx"
	"github.com/unrolled/render"
	"github.com/zenazn/goji/web"
)

func GetActivities(c web.C, w http.ResponseWriter, r *http.Request) {
	template := c.Env["render"].(*render.Render)
	db := c.Env["mysql"].(*sqlx.DB)
	cookie, _ := r.Cookie("injuv_auth")
	claims, _ := security.Decode(cookie.Value)
	users, _ := user.GetAll(db)
	bnd := binding.GetDefault(r)
	bnd["Users"] = users
	bnd["ID"] = claims["id"]
	bnd["ADMIN"] = claims["guuid"]
	template.HTML(w, http.StatusOK, "panel/activities", bnd, render.HTMLOptions{
		Layout: "panel/layout",
	})
}

func NewActivitie(c web.C, w http.ResponseWriter, r *http.Request) {
	template := c.Env["render"].(*render.Render)
	name := utils.GetAndTrim(r, "name")
	description := utils.GetAndTrim(r, "description")
	dateExpire := utils.GetAndTrim(r, "fechaExpiracion")
	recursos := utils.GetAndTrim(r, "recurso")
	estado := utils.GetAndTrim(r, "estado")
	bnd := binding.GetDefault(r)

	bnd["User"] = name
	bnd["Description"] = description

	template.JSON(w, http.StatusOK, bnd)
}
