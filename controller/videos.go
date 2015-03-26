package controller

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path"
	"strconv"

	"github.com/delba/goembed/model"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

type Videos struct{}

func (v *Videos) Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(path.Join("views", "index.html"))
	handle(err)

	items, err := model.AllItems()
	handle(err)

	err = t.Execute(w, items)
	handle(err)
}

func (v *Videos) Show(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query()["id"][0])
	handle(err)

	item, err := model.FindItem(id)
	handle(err)

	data, err := json.Marshal(item)
	handle(err)

	w.Write(data)
}

func (v *Videos) Create(w http.ResponseWriter, r *http.Request) {
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
