package controller

import (
	"fmt"
	"net/http"
)

type Users struct{}

func (u *Users) Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		u.New(w, r)
	case "POST":
		u.Create(w, r)
	default:
		u.New(w, r)
	}
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UsersNew")
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UsersCreate")
}
