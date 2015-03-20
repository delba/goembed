package main

import (
	"html/template"
	"net/http"
	"path"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(path.Join("views", "index.html"))
	handle(err)

	err = t.Execute(w, nil)
}

func Embed(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/embed", Embed)
}
