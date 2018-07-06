package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
)

type CountHandler struct {
	count int
}

func (c *CountHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.count++
	w.Header().Set("User-Agent", "GoServer")
	w.WriteHeader(http.StatusFound)
	fmt.Fprint(w, c.count)
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func CreateUsers(w http.ResponseWriter, r *http.Request) {
	u := new(User)
	err := json.NewDecoder(r.Body).Decode(u)
	fmt.Printf("err : %v, body : %+v\n", err, u)
	return
}

func MiddlewareFunc(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		u := new(User)
		err := json.NewDecoder(request.Body).Decode(u)
		fmt.Printf("err : %v, body : %+v\n", err, u)
		next(writer, request)
	}
}

func main() {
	http.Handle("/count", &CountHandler{})
	http.HandleFunc("/user", MiddlewareFunc(CreateUsers))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
