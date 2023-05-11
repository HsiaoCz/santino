package main

import (
	"encoding/json"
	"net/http"
)

type Handler func(Context) error

func POST(route string, h Handler) {
	h(Context{})
}

func main() {

	POST("/user", handlerCreateUser)
}

type Context struct {
	r *http.Request
	w http.ResponseWriter
}

// DATA
type User struct {
	ID                int
	FirstName         string
	LastName          string
	EncryptedPassword string
	Email             string
	Posts             []any
	IsAdmin           bool
	VerificationCode  int
}

type CreateUserParams struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
}

func handlerCreateUser(c Context) error {
	var params CreateUserParams
	if err := json.NewDecoder(c.r.Body).Decode(&params); err != nil {
		return err
	}
	var user any
	return JSON(http.StatusOK, user)
}

func JSON(code int, v any) error {
	return nil
}
