package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path"
	"strconv"

	"github.com/delba/goembed/model"
	"github.com/julienschmidt/httprouter"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

type Items struct{}

func (i *Items) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, err := template.ParseFiles(path.Join("views", "items", "index.html"))
	handle(err)

	items, err := model.AllItems()
	handle(err)

	err = t.Execute(w, items)
	handle(err)
}

func (i *Items) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	handle(err)

	item, err := model.FindItem(id)
	handle(err)

	data, err := json.Marshal(item)
	handle(err)

	w.Write(data)
}

func (i *Items) Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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
