package models

import (
	"net/url"
	"os"

	"github.com/garyburd/redigo/redis"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

var c redis.Conn

func init() {
	var (
		u   *url.URL
		err error
	)

	if redisURL := os.Getenv("REDISTOGO_URL"); redisURL != "" {
		u, err = url.Parse(redisURL)
		handle(err)
	} else {
		u = &url.URL{Host: "127.0.0.1:6379"}
	}

	c, err = redis.Dial("tcp", u.Host)
	handle(err)

	if u.User != nil {
		if pw, ok := u.User.Password(); ok {
			_, err = c.Do("AUTH", pw)
			handle(err)
		}
	}
}
