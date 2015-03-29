package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Users struct{}

func (u *Users) New(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("UsersNew")
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("UsersCreate")
}
