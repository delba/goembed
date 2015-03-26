package main

import (
	"net/http"
	"os"

	"github.com/delba/goembed/controller"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/items/", controller.Show)
	http.HandleFunc("/embed", controller.Create)

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.Handle("/favicon.ico", fs)

	http.ListenAndServe(":"+port, nil)
}
