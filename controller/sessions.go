package controller

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Sessions struct{}

func (s *Sessions) New(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Sessions new")
}

func (s *Sessions) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Sessions create")
}

func (s *Sessions) Destroy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("Sessions destroy")
}
