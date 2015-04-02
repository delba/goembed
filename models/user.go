package models

import (
	"errors"
	"strings"

	"github.com/garyburd/redigo/redis"
)

type User struct {
	ID       int
	Username string
	Password string
}

func CreateUser(username string, password string) (user User, err error) {
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	isMember, err := redis.Bool(c.Do("SISMEMBER", "users:usernames", username))
	if err != nil {
		return
	}

	if isMember {
		err = errors.New("Username already taken")
		return
	}

	id, err := redis.Int(c.Do("INCR", "users:uid"))
	if err != nil {
		return
	}

	user = User{ID: id, Username: username, Password: password}

	c.Send("MULTI")
	c.Send("HMSET", redis.Args{}.Add("users:"+string(id)).AddFlat(user)...)
	c.Send("SADD", "users:usernames", username)
	c.Send("LPUSH", "users:ids", id)
	c.Send("SET", "users:id:"+username, id)
	_, err = c.Do("EXEC")

	return user, err
}
