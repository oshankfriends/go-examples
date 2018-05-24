package main

import (
	"net/http"
	"fmt"
)

type CountHandler struct {
	count   int
}

func(c *CountHandler)ServeHTTP(w http.ResponseWriter,r *http.Request){
	c.count ++

	w.Header().Set("User-Agent","GoServer")
	w.WriteHeader(http.StatusFound)
	fmt.Fprint(w,c.count)
}

func main(){
	http.Handle("/count",&CountHandler{})
	http.ListenAndServe(":8000",nil)
}