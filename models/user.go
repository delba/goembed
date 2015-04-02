package models

import (
	"errors"
	"strings"

	"github.com/garyburd/redigo/redis"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Username string
	Password []byte
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

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return
	}

	user = User{Username: username, Password: hashedPassword}

	id, err := redis.Int(c.Do("INCR", "users:uid"))
	if err != nil {
		return
	}

	user.ID = id

	c.Send("MULTI")
	c.Send("HMSET", redis.Args{}.Add("users:"+string(id)).AddFlat(user)...)
	c.Send("SADD", "users:usernames", username)
	c.Send("LPUSH", "users:ids", id)
	c.Send("SET", "users:id:"+username, id)
	_, err = c.Do("EXEC")

	return user, err
}

func FindUser(id int) (user User, err error) {
	values, err := redis.Values(c.Do("HGETALL", "users:"+string(id)))
	if err != nil {
		return
	}

	err = redis.ScanStruct(values, &user)

	return user, err
}

func FindUserByUsername(username string) (user User, err error) {
	id, err := redis.Int(c.Do("GET", "users:id:"+username))
	if err != nil {
		return
	}

	user, err = FindUser(id)

	return user, err
}

func AuthenticateUser(username string, password string) (user User, err error) {
	user, err = FindUserByUsername(username)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	if err != nil {
		return
	}

	return user, err
}
