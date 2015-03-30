package models

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/garyburd/redigo/redis"
)

type Item struct {
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
	ID              int    `json:"id"`
}

func (i *Item) RawHTML() template.HTML {
	return template.HTML(i.HTML)
}

func FindItem(id int) (item Item, err error) {
	values, err := redis.Values(c.Do("HGETALL", "items:"+string(id)))
	if err != nil {
		return
	}

	err = redis.ScanStruct(values, &item)

	return item, err
}

func FindItemByURL(url string) (item Item, err error) {
	id, err := redis.Int(c.Do("GET", "items:id:"+url))
	if err != nil {
		return
	}

	item, err = FindItem(id)

	return item, err
}

func AllItems() (items []Item, err error) {
	ids, err := redis.Ints(c.Do("LRANGE", "items:ids", "0", "-1"))
	if err != nil {
		return
	}

	c.Send("MULTI")

	for _, id := range ids {
		c.Send("HGETALL", "items:"+string(id))
	}

	values, err := redis.Values(c.Do("EXEC"))
	if err != nil {
		return
	}

	var item Item

	for _, v := range values {
		values, err = redis.Values(v, nil)
		if err != nil {
			return
		}

		err = redis.ScanStruct(values, &item)
		if err != nil {
			return
		}

		items = append(items, item)
	}

	return items, err
}

func CreateItem(url string) (item Item, err error) {
	url = strings.TrimSpace(url)

	isMember, err := redis.Bool(c.Do("SISMEMBER", "items:urls", url))
	if err != nil {
		return
	}

	if isMember {
		item, err = FindItemByURL(url)
		return item, err
	}

	item, err = OEmbed(url)
	if err != nil {
		return
	}

	id, err := redis.Int(c.Do("INCR", "items:uid"))
	if err != nil {
		return
	}

	item.ID = id

	c.Send("MULTI")
	c.Send("HMSET", redis.Args{}.Add("items:"+string(id)).AddFlat(item)...)
	c.Send("SADD", "items:urls", url)
	c.Send("LPUSH", "items:ids", id)
	c.Send("SET", "items:id:"+url, id)
	_, err = c.Do("EXEC")

	return item, err
}

func OEmbed(url string) (item Item, err error) {
	res, err := http.Get("https://vimeo.com/api/oembed.json?url=" + url)
	if err != nil {
		return
	}

	defer res.Body.Close()
	contents, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(contents, &item)

	return item, err
}
