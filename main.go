package main

import (
	"net/http"
	"os"

	"github.com/delba/goembed/controllers"
	"github.com/julienschmidt/httprouter"
)

var (
	items    controllers.Items
	users    controllers.Users
	sessions controllers.Sessions
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := httprouter.New()
	router.GET("/", items.Index)
	router.GET("/items/:id", items.Show)
	router.POST("/embed", items.Create)
	router.GET("/register", users.New)
	router.POST("/register", users.Create)
	router.GET("/login", sessions.New)
	router.POST("/login", sessions.Create)
	router.DELETE("/logout", sessions.Destroy)

	fs := http.FileServer(http.Dir("public"))
	router.Handler("GET", "/public/*filepath", http.StripPrefix("/public/", fs))
	router.Handler("GET", "/favicon.ico", fs)

	http.ListenAndServe(":"+port, router)
}
