package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Validater interface {
	Vaildate() []error
}

type Handler[T any] func(Context[T]) error

func makeHTTPHandler[T any](h Handler[T]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var reqData T
		if err := json.NewDecoder(r.Body).Decode(&reqData); err != nil {
			// todo
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		errs := Validate(reqData)
		if len(errs) > 0 {
			panic(errs)
		}
		h(Context[T]{
			r:            r,
			w:            w,
			RequestParam: reqData,
		})
	}
}

func Validate(data any) []error {
	if v, ok := data.(Validater); ok {
		return v.Vaildate()
	}
	return nil
}

func POST[T any](route string, h Handler[T]) {
	http.HandleFunc(route, makeHTTPHandler(h))
}

func main() {

	POST("/user", handlerCreateUser[CreateUserParams])
	http.ListenAndServe(":9091", nil)
}

type Context[T any] struct {
	r            *http.Request
	w            http.ResponseWriter
	RequestParam T
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
	Email            string `json:"email"`
	Password         string `json:"password"`
	Username         string `json:"username"`
	VerificationCode int    `json:"code"`
}

func (c CreateUserParams) Validate() []error {
	return []error{fmt.Errorf("Damn")}
}

func handlerCreateUser[T any](c Context[T]) error {
	userParams := c.RequestParam
	fmt.Println(userParams)
	var user User
	return Page("user", user)
}

func Page(file string, v any) error { return nil }

func JSON(code int, v any) error {
	return nil
}
