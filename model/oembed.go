package model

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/garyburd/redigo/redis"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

var c redis.Conn

func init() {
	var err error

	c, err = redis.Dial("tcp", "127.0.0.1:6379")
	handle(err)
}

type OEmbed struct {
	AuthorName      string `json:"author_name"`
	AuthorURL       string `json:"author_url"`
	Description     string `json:"description"`
	Duration        int    `json:"duration"`
	Height          int    `json:"height"`
	HTML            string `json:"html"`
	IsPlus          string `json:"is_plus"`
	ProviderName    string `json:"provider_name"`
	ProviderURL     string `json:"provider_url"`
	ThumbnailHeight int    `json:"thumbnail_height"`
	ThumbnailURL    string `json:"thumbnail_url"`
	ThumbnailWidth  int    `json:"thumbnail_width"`
	Title           string `json:"title"`
	Type            string `json:"type"`
	URI             string `json:"uri"`
	Version         string `json:"version"`
	VideoID         int    `json:"video_id"`
	Width           int    `json:"width"`
}

func (o *OEmbed) RawHTML() template.HTML {
	return template.HTML(o.HTML)
}

func AllOEmbeds() ([]OEmbed, error) {
	var oembeds []OEmbed
	var err error

	urls, err := redis.Strings(c.Do("LRANGE", "myurls", "0", "-1"))
	if err != nil {
		return oembeds, err
	}

	c.Send("MULTI")

	for _, url := range urls {
		c.Send("HGETALL", url)
	}

	values, err := redis.Values(c.Do("EXEC"))
	if err != nil {
		return oembeds, err
	}

	var oembed OEmbed

	for _, v := range values {
		values, err = redis.Values(v, nil)
		if err != nil {
			return oembeds, err
		}

		err = redis.ScanStruct(values, &oembed)
		if err != nil {
			return oembeds, err
		}

		oembeds = append(oembeds, oembed)
	}

	return oembeds, err
}

func CreateOEmbed(url string) (OEmbed, error) {
	var oembed OEmbed
	var err error

	res, err := http.Get("https://vimeo.com/api/oembed.json?url=" + url)
	if err != nil {
		return oembed, err
	}
	defer res.Body.Close()

	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return oembed, err
	}

	err = json.Unmarshal(contents, &oembed)
	if err != nil {
		return oembed, err
	}

	c.Send("MULTI")
	c.Send("HMSET", redis.Args{}.Add(url).AddFlat(oembed)...)
	c.Send("SADD", "urls", url)
	c.Send("LPUSH", "myurls", url)
	_, err = c.Do("EXEC")
	if err != nil {
		return oembed, err
	}

	return oembed, err
}
