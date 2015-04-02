package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/delba/goembed/models"
	"github.com/julienschmidt/httprouter"
)

type Users struct{}

func (u *Users) New(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	layoutFile := layoutPath("application.html")
	viewFile := viewPath("users", "new.html")
	t, err := template.ParseFiles(layoutFile, viewFile)
	handle(err)

	err = t.Execute(w, nil)
	handle(err)
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	user, err := models.CreateUser(username, password)
	handle(err)

	fmt.Println(user.ID, user.Username, user.Password)

	http.Redirect(w, r, "/", 302)
}
