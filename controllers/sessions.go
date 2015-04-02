package controllers

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/delba/goembed/models"
	"github.com/julienschmidt/httprouter"
)

type Sessions struct{}

func (s *Sessions) New(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	layoutFile := layoutPath("application.html")
	viewFile := viewPath("sessions", "new.html")
	t, err := template.ParseFiles(layoutFile, viewFile)
	handle(err)

	err = t.Execute(w, nil)
	handle(err)
}

func (s *Sessions) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	fmt.Println(username, password)

	user, err := models.AuthenticateUser(username, password)
	if err != nil {
		if err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
			err = renderTemplate("new.html", w)
			handle(err)
			return
		}

		handle(err)
	}

	fmt.Println("User is authenticated.")
	fmt.Println(user.Username, user.Password)

	http.Redirect(w, r, "/login", 302)
}

func (s *Sessions) Destroy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Sessions destroy")
}

func renderTemplate(p string, w http.ResponseWriter) (err error) {
	layoutFile := layoutPath("application.html")
	viewFile := viewPath("sessions", p)
	t, err := template.ParseFiles(layoutFile, viewFile)
	if err != nil {
		return
	}

	err = t.Execute(w, nil)

	return err
}
