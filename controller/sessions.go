package controller

import (
	"fmt"
	"net/http"
)

type Sessions struct{}

func (s *Sessions) Login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		s.New(w, r)
	case "POST":
		s.Create(w, r)
	default:
		s.New(w, r)
	}
}

func (s *Sessions) New(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sessions new")
}

func (s *Sessions) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sessions create")
}

func (s *Sessions) Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Sessions destroy")
}
