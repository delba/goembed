package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/delba/goembed/models"
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
	url := r.FormValue("url")
	res, err := http.Get("https://vimeo.com/api/oembed.json?url=" + url)
	handle(err)
	defer res.Body.Close()

	contents, err := ioutil.ReadAll(res.Body)
	handle(err)

	var oembed models.OEmbed
	json.Unmarshal(contents, &oembed)

	fmt.Println(oembed.HTML)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", Index)
	http.HandleFunc("/embed", Embed)

	http.ListenAndServe(":"+port, nil)
}
