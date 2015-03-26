package main

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"path"
	"strconv"

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

	items, err := model.AllItems()
	handle(err)

	err = t.Execute(w, items)
	handle(err)
}

func Show(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query()["id"][0])
	handle(err)

	item, err := model.FindItem(id)
	handle(err)

	data, err := json.Marshal(item)
	handle(err)

	w.Write(data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Url param can't be blank"))
		return
	}

	item, err := model.CreateItem(url)
	handle(err)

	data, err := json.Marshal(item)
	handle(err)

	w.Write(data)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", Index)
	http.HandleFunc("/items/", Show)
	http.HandleFunc("/embed", Create)

	fs := http.FileServer(http.Dir("public"))
	http.Handle("/public/", http.StripPrefix("/public/", fs))
	http.Handle("/favicon.ico", fs)

	http.ListenAndServe(":"+port, nil)
}
