package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-gomail/gomail"
	"github.com/injuvproject/backofficeinjuv/config"
	"github.com/injuvproject/backofficeinjuv/helpers/security"
	"github.com/injuvproject/backofficeinjuv/helpers/utils"
	"github.com/injuvproject/backofficeinjuv/models/activity"
	"github.com/injuvproject/backofficeinjuv/models/binding"
	"github.com/injuvproject/backofficeinjuv/models/user"
	"github.com/jmoiron/sqlx"
	"github.com/unrolled/render"
	"github.com/zenazn/goji/web"
)

var (
	StrUserEmpty        = "El nombre esta vacio"
	StrDescriptionEmpty = "La descripcion es vacia"
	StrDateExpireEmpty  = "Fecha de expiracion invalida"
	StrRecursosEmpty    = "No se asigno recurso"
	StrEstadoEmpty      = "No se asigno estado"
	StrPioridadEmpty    = "No se asigno pioridad"
	errStrUnableToSend  = "No se a podido enviar el email intentelo mas tarde."

	strSenderFormat = "%s <%s>"
	strEmailFormat  = "%s\n\n--\n%s"
	errBadMail      = errors.New(errStrUnableToSend)
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
	db := c.Env["mysql"].(*sqlx.DB)
	bnd := binding.GetDefault(r)

	name := utils.GetAndTrim(r, "name")
	description := utils.GetAndTrim(r, "description")
	dateExpire := utils.GetAndTrim(r, "fechaExpiracion")
	recursos := utils.GetAndTrim(r, "recurso")
	estado := utils.GetAndTrim(r, "estado")
	pioridad := utils.GetAndTrim(r, "pioridad")
	userid, _ := strconv.Atoi(recursos)

	if name == "" {
		bnd["Error"] = StrUserEmpty
	}

	if description == "" {
		bnd["Error"] = StrDescriptionEmpty
	}

	activityNew := &activity.Activity{
		Title:       name,
		Description: description,
		ExpireDate:  dateExpire,
		User:        userid,
		Estate:      estado,
		Adjuntos:    0,
		Pioridad:    pioridad,
	}

	activityNew.Create(db)

	userAdd, _ := user.Get(db, userid)

	//TO_DO LOAD EMAIL FOR USER=user

	emailSender := fmt.Sprintf(strSenderFormat, "injuv araucania", config.SendTo)
	emailFullSubject := fmt.Sprintf("%s", "Notificacion nueva tarea")
	emailMessage := fmt.Sprintf("Se a creado una nueva Tarea %s fecha termino tarea %s", name, dateExpire)
	emailFullName := fmt.Sprintf("Asignada a %s %s", userAdd.FirstName, userAdd.LastName)
	emailBody := fmt.Sprintf(
		strEmailFormat,
		emailMessage,
		emailFullName,
	)

	err := sendEmail(emailBody, emailSender, emailFullSubject, userAdd.Email)
	if err == errBadMail {
		bnd["Error"] = errStrUnableToSend
		template.JSON(w, http.StatusOK, bnd)
		return
	}

	template.JSON(w, http.StatusOK, bnd)
}

func GetActivitiesByUser(c web.C, w http.ResponseWriter, r *http.Request) {

}

func sendEmail(Body, Sender, Subject, sentTo string) error {

	msg := gomail.NewMessage()
	msg.SetHeader("From", Sender)
	msg.SetHeader("To", sentTo)
	msg.SetAddressHeader("Cco", config.SendCopy, config.SendName)
	msg.SetHeader("Subject", Subject)
	msg.SetBody("text/plain", Body)

	mailer := gomail.NewPlainDialer(config.SmtpServer, config.SmptPort, config.SmtpUser, config.SmtpToken)

	if err := mailer.DialAndSend(msg); err != nil {
		return errBadMail
	}

	return nil
}
