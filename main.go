package main

import (
	"net/http"
	"os"

	"github.com/delba/goembed/controller"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	var videos controller.Videos
	http.HandleFunc("/", videos.Index)
	http.HandleFunc("/items/", videos.Show)
	http.HandleFunc("/embed", videos.Create)

	var users controller.Users
	http.HandleFunc("/register", users.Register)

	var sessions controller.Sessions
	http.HandleFunc("/login", sessions.Login)
	http.HandleFunc("/logout", sessions.Logout)

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.Handle("/favicon.ico", fs)

	http.ListenAndServe(":"+port, nil)
}
