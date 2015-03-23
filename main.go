package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/delba/goembed/models"
	"github.com/garyburd/redigo/redis"
)

var c redis.Conn

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	var err error

	c, err = redis.Dial("tcp", "127.0.0.1:6379")
	handle(err)
}

func Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(path.Join("views", "index.html"))
	handle(err)

	res, err := redis.Values(c.Do("HGETALL", "https://vimeo.com/28018829"))
	var oembed models.OEmbed
	redis.ScanStruct(res, &oembed)

	data := struct {
		HTML template.HTML
	}{
		HTML: template.HTML(oembed.HTML),
	}

	err = t.Execute(w, data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Url param can't be blank"))
		return
	}
	res, err := http.Get("https://vimeo.com/api/oembed.json?url=" + url)
	handle(err)
	defer res.Body.Close()

	contents, err := ioutil.ReadAll(res.Body)
	handle(err)

	var oembed models.OEmbed
	json.Unmarshal(contents, &oembed)
	c.Send("HMSET", redis.Args{}.Add(url).AddFlat(oembed)...)

	var buf bytes.Buffer
	err = json.Indent(&buf, contents, "", "  ")
	if err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(buf.Bytes())
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
