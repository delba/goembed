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

	urls, err := redis.Strings(c.Do("LRANGE", "myurls", "0", "-1"))

	var data []models.OEmbed
	for _, url := range urls {
		res, _ := redis.Values(c.Do("HGETALL", url))
		var oembed models.OEmbed
		redis.ScanStruct(res, &oembed)

		oembed.RawHTML = template.HTML(oembed.HTML)

		data = append(data, oembed)
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

	c.Send("MULTI")
	c.Send("HMSET", redis.Args{}.Add(url).AddFlat(oembed)...)
	c.Send("SADD", "urls", url)
	c.Send("LPUSH", "myurls", url)
	_, err = c.Do("EXEC")
	handle(err)

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
