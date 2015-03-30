package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/delba/goembed/models"
	"github.com/julienschmidt/httprouter"
)

func init() {

}

func TestCreate(t *testing.T) {
	v := url.Values{}
	v.Set("url", "http://vimeo.com/18150336")

	req, err := http.NewRequest(
		"POST",
		"http://example.com/embed/",
		bytes.NewBuffer([]byte(v.Encode())),
	)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		log.Fatal(err)
	}
	w := httptest.NewRecorder()
	ps := httprouter.Params{}
	items.Create(w, req, ps)

	var item models.Item
	json.Unmarshal(w.Body.Bytes(), &item)
	if item.AuthorURL != "https://vimeo.com/phoenixfly" {
		t.Errorf("Incorrect author url")
	}
	if item.URI != "/videos/18150336" {
		t.Errorf("Incorrect uri")
	}
}

func TestEmptyCreate(t *testing.T) {

	req, err := http.NewRequest(
		"POST",
		"http://example.com/embed/",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	w := httptest.NewRecorder()
	ps := httprouter.Params{}
	items.Create(w, req, ps)
	if w.Code != 406 {
		t.Errorf("Wrong response code for invalid request")
	}
}

func TestIndex(t *testing.T) {
	req, err := http.NewRequest(
		"POST",
		"http://example.com/",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	w := httptest.NewRecorder()
	ps := httprouter.Params{}
	items.Index(w, req, ps)
	if w.Code != 200 {
		t.Errorf("Wrong response code for index request")
	}

	// Maybe parse the page content here and confirm that
	// it's outputting the right content.
}
