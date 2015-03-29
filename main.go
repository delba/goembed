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

	var (
		videos   controller.Videos
		users    controller.Users
		sessions controller.Sessions
	)

	var routes = map[string]http.HandlerFunc{
		"/":         videos.Index,
		"/items/":   videos.Show,
		"/embed":    videos.Create,
		"/register": users.Register,
		"/login":    sessions.Login,
		"/logout":   sessions.Logout,
	}

	for path, handlerFunc := range routes {
		http.HandleFunc(path, handlerFunc)
	}

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.Handle("/favicon.ico", fs)

	http.ListenAndServe(":"+port, nil)
}
