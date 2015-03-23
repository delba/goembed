package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"path"

	"github.com/delba/goembed/model"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(path.Join("views", "index.html"))
	handle(err)

	oembeds, err := model.AllOEmbeds()
	handle(err)

	err = t.Execute(w, oembeds)
	handle(err)
}

func Create(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Url param can't be blank"))
		return
	}

	oembed, err := model.CreateOEmbed(url)
	handle(err)

	data, err := json.Marshal(oembed)
	handle(err)

	w.Write(data)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", Index)
	http.HandleFunc("/embed", Create)

	http.ListenAndServe(":"+port, nil)
}
