package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
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

var (
	strFirstNameEmpty = errors.New("El nombre no fue ingresado.")
	strLastNameEmpty  = errors.New("El apellido no fue ingresado.")
	strUserNameEmpty  = errors.New("El usuario no fue ingresado.")
	strPasswordEmpty  = errors.New("El password no fue ingresado.")
	strErrorPassword  = errors.New("Las password no coinciden.")
	strDeleteSuccess  = "User was removed successfully."
	strEditSuccess    = "User updated successfully."

	URLPanelUserNew    = "/panel/%d/cuentas.html?add=%s"
	URLPanelUserByID   = "/panel/%d/cuenta/%d/editar.html?edit=%s"
	URLProfileUserByID = "/panel/%d/perfil/%d/editar.html?edit=%s"
	URLPanelDelete     = "/panel/%d/cuentas.html?delete=%s"
)

func GetAccounts(c web.C, w http.ResponseWriter, r *http.Request) {
	template := c.Env["render"].(*render.Render)
	db := c.Env["mysql"].(*sqlx.DB)
	bnd := binding.Binding{}
	numberPage, _ := strconv.Atoi(c.URLParams["page"])
	cookie, _ := r.Cookie("injuv_auth")
	claims, _ := security.Decode(cookie.Value)

	var (
		init   int
		page   []int
		Status string
	)

	if numberPage == 0 || numberPage == 1 {
		init = 1
	} else {
		init = numberPage
	}

	users, _ := user.Range(db, init, config.NumberResultPage)
	total := user.CountAll(db)
	totalPage := (total / config.NumberResultPage)

	for i := 1; i <= totalPage; i++ {
		page = append(page, i)
	}

	if r.FormValue("delete") == getresponse.ComparableValue {
		Status = strDeleteSuccess
	}

	bnd = binding.Binding{

		"PageTitle":  "Back office INJUV",
		"CurrentURL": r.URL.Path,
		"ID":         claims["id"],
		"ADMIN":      claims["guuid"],
		"Page":       page,
		"Users":      users,
		"TotalPage":  totalPage,
		"Success":    Status,
	}

	template.HTML(w, http.StatusOK, "panel/accounts", bnd, render.HTMLOptions{
		Layout: "panel/layout",
	})
}

func GetAccount(c web.C, w http.ResponseWriter, r *http.Request) {
	template := c.Env["render"].(*render.Render)
	db := c.Env["mysql"].(*sqlx.DB)
	bnd := binding.GetDefault(r)
	id, _ := strconv.Atoi(c.URLParams["uid"])
	cookie, _ := r.Cookie("injuv_auth")
	claims, _ := security.Decode(cookie.Value)
	u, err := user.Get(db, id)

	if err != nil {
		panic(err)
	}

	if r.FormValue("edit") == getresponse.ComparableValue {
		bnd["Success"] = strEditSuccess
	}

	bnd["User"] = u
	bnd["ID"] = claims["id"]
	bnd["ADMIN"] = claims["guuid"]

	template.HTML(w, http.StatusOK, "panel/show", bnd, render.HTMLOptions{
		Layout: "panel/layout",
	})
}

func GetNewAccount(c web.C, w http.ResponseWriter, r *http.Request) {
	template := c.Env["render"].(*render.Render)
	bnd := binding.GetDefault(r)
	cookie, _ := r.Cookie("injuv_auth")
	claims, _ := security.Decode(cookie.Value)

	bnd["ID"] = claims["id"]
	bnd["ADMIN"] = claims["guuid"]

	template.HTML(w, http.StatusOK, "panel/edit", bnd, render.HTMLOptions{
		Layout: "panel/layout",
	})
}

func PostNewAccount(c web.C, w http.ResponseWriter, r *http.Request) {
	bnd := binding.GetDefault(r)
	db := c.Env["mysql"].(*sqlx.DB)
	template := c.Env["render"].(*render.Render)
	id, _ := strconv.Atoi(c.URLParams["uid"])
	cookie, _ := r.Cookie("injuv_auth")
	claims, _ := security.Decode(cookie.Value)

	firstName := utils.GetAndTrim(r, "firstname")
	lastName := utils.GetAndTrim(r, "lastname")
	userName := utils.GetAndTrim(r, "userName")
	email := utils.GetAndTrim(r, "email")
	password := utils.GetAndTrim(r, "password")
	rePassword := utils.GetAndTrim(r, "repassword")
	userAdmin, _ := strconv.ParseBool(utils.GetAndTrim(r, "admin"))

	fmt.Println("%t\n\n", userAdmin)

	if firstName == "" {
		bnd["Error"] = strFirstNameEmpty
	}

	if lastName == "" {
		bnd["Error"] = strLastNameEmpty
	}

	if userName == "" {
		bnd["Error"] = strUserNameEmpty
	}

	if password == "" {
		bnd["Error"] = strPasswordEmpty
	}
	if rePassword == "" {
		bnd["Error"] = strPasswordEmpty
	}
	if password != rePassword {
		bnd["Error"] = strErrorPassword
	}

	if bnd["Error"] != nil {
		template.HTML(w, http.StatusOK, "panel/edit", bnd)
		return
	}

	newUser := &user.User{
		ID:         id,
		FirstName:  firstName,
		LastName:   lastName,
		UserName:   userName,
		Email:      email,
		Password:   password,
		Admin:      userAdmin,
		SignupDate: time.Now(),
	}
	ids := claims["id"].(float64)
	bnd["ADMIN"] = claims["guuid"]
	bnd["User"] = newUser

	newUser.Create(db)

	http.Redirect(w, r, fmt.Sprintf(URLPanelUserNew, int(ids), getresponse.ComparableValue), http.StatusFound)
	return

}

func GetEditAccount(c web.C, w http.ResponseWriter, r *http.Request) {
	template := c.Env["render"].(*render.Render)
	db := c.Env["mysql"].(*sqlx.DB)
	bnd := binding.GetDefault(r)
	id, _ := strconv.Atoi(c.URLParams["uid"])
	cookie, _ := r.Cookie("injuv_auth")
	claims, _ := security.Decode(cookie.Value)
	u, err := user.Get(db, id)

	if err != nil {
		panic(err)
	}

	bnd["User"] = u
	bnd["ID"] = claims["id"]
	bnd["ADMIN"] = claims["guuid"]

	template.HTML(w, http.StatusOK, "panel/edit", bnd, render.HTMLOptions{
		Layout: "panel/layout",
	})
}

func PostEditAccount(c web.C, w http.ResponseWriter, r *http.Request) {
	bnd := binding.GetDefault(r)
	db := c.Env["mysql"].(*sqlx.DB)
	template := c.Env["render"].(*render.Render)
	id, _ := strconv.Atoi(c.URLParams["uid"])
	cookie, _ := r.Cookie("injuv_auth")
	claims, _ := security.Decode(cookie.Value)

	firstName := utils.GetAndTrim(r, "firstname")
	lastName := utils.GetAndTrim(r, "lastname")
	userName := utils.GetAndTrim(r, "userName")
	email := utils.GetAndTrim(r, "email")
	password := utils.GetAndTrim(r, "password")
	rePassword := utils.GetAndTrim(r, "repassword")
	userAdmin, _ := strconv.ParseBool(utils.GetAndTrim(r, "admin"))

	fmt.Println("%t\n\n", userAdmin)

	if firstName == "" {
		bnd["Error"] = strFirstNameEmpty
	}

	if lastName == "" {
		bnd["Error"] = strLastNameEmpty
	}

	if userName == "" {
		bnd["Error"] = strUserNameEmpty
	}

	if password == "" {
		bnd["Error"] = strPasswordEmpty
	}
	if rePassword == "" {
		bnd["Error"] = strPasswordEmpty
	}
	if password != rePassword {
		bnd["Error"] = strErrorPassword
	}

	if bnd["Error"] != nil {
		template.HTML(w, http.StatusOK, "panel/edit", bnd)
		return
	}

	newUser := &user.User{
		ID:         id,
		FirstName:  firstName,
		LastName:   lastName,
		UserName:   userName,
		Email:      email,
		Password:   password,
		Admin:      userAdmin,
		SignupDate: time.Now(),
	}
	ids := claims["id"].(float64)
	bnd["ADMIN"] = claims["guuid"]
	bnd["User"] = newUser

	newUser.Save(db)

	if claims["guuid"].(bool) == true {
		http.Redirect(w, r, fmt.Sprintf(URLPanelUserByID, int(ids), id, getresponse.ComparableValue), http.StatusFound)
		return
	}

	http.Redirect(w, r, fmt.Sprintf(URLProfileUserByID, int(ids), id, getresponse.ComparableValue), http.StatusFound)
	return

}

func GetDeleteProfile(c web.C, w http.ResponseWriter, r *http.Request) {
	template := c.Env["render"].(*render.Render)
	db := c.Env["mysql"].(*sqlx.DB)
	bnd := binding.GetDefault(r)
	cookie, _ := r.Cookie("injuv_auth")
	claims, _ := security.Decode(cookie.Value)

	ids := claims["id"].(float64)

	id, _ := strconv.Atoi(c.URLParams["uid"])
	err := user.Delete(db, id)
	bnd["ID"] = ids
	bnd["ADMIN"] = claims["guuid"]

	if err == user.ErrUserNotExist {
		bnd["Error"] = user.ErrUserNotExist
		template.HTML(w, http.StatusOK, "/panel/profile", bnd)
	}

	http.Redirect(w, r, fmt.Sprintf(URLPanelDelete, int(ids), getresponse.ComparableValue), http.StatusFound)
	return

}
