package controllers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/delba/goembed/models"
	"github.com/julienschmidt/httprouter"
)

type Items struct{}

func (i *Items) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	layoutFile := layoutPath("application.html")
	viewFile := viewPath("items", "index.html")
	t, err := template.ParseFiles(layoutFile, viewFile)
	handle(err)

	items, err := models.AllItems()
	handle(err)

	err = t.Execute(w, items)
	handle(err)
}

func (i *Items) Show(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	handle(err)

	item, err := models.FindItem(id)
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

	item, err := models.CreateItem(url)
	handle(err)

	data, err := json.Marshal(item)
	handle(err)

	w.Write(data)
}
