package main

import (
	"fmt"
	"net/http"
	"time"
)

func Cookie(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{Name: "testcookie", Value: "testValue", Expires: time.Now().Add(time.Second * 30), HttpOnly: true}
	http.SetCookie(w, cookie)
	fmt.Fprint(w, cookie.String())
}
func main() {
	http.HandleFunc("/cookie", Cookie)
	http.ListenAndServe(":9090", nil)
}
