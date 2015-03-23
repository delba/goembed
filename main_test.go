package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func init() {

}

func TestEmbed(t *testing.T) {
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
	Embed(w, req)

	fmt.Printf("%d - %s", w.Code, w.Body.String())

}

func TestEmptyEmbed(t *testing.T) {

	req, err := http.NewRequest(
		"POST",
		"http://example.com/embed/",
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	w := httptest.NewRecorder()
	Embed(w, req)
	if w.Code != 406 {

	}
	fmt.Printf("%d - %s", w.Code, w.Body.String())
}

func TestIndex(t *testing.T) {

}
