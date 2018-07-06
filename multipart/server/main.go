package main

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"net/http"
	"os"
	"time"
)

const chunkSize = 32

type WebHandler struct {
}

func (wb *WebHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	file, _ := os.Create("test.wav")
	defer file.Close()
	mr, err := r.MultipartReader()
	if err != nil {
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	part, err := mr.NextPart()
	for err == nil {
		log.Infof("Part = %+v", part)
		if part.Header.Get("Content-Type") == "audio/L16; rate=16000; channels=1" {
			var buf [chunkSize]byte
			part.Read(buf[:])
			file.Write(buf[:])
		}
		part, err = mr.NextPart()
	}
	log.Error(err)
}

type MyHandler struct {
	Body map[string]interface{}
}

func (m *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	json.NewDecoder(r.Body).Decode(&m.Body)
	fmt.Fprint(w, m.Body)
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("%s %s ", r.Method, r.URL)
		defer func(begin time.Time) {
			log.Infof("%s %s resp time: %s", r.Method, r.URL, time.Since(begin))
		}(time.Now())
		next.ServeHTTP(w, r)
	})
}

func main() {

	http.DefaultServeMux.Handle("/multi", &WebHandler{})
	http.DefaultServeMux.Handle("/pipe", &MyHandler{})
	s := &http.Server{
		Addr:    ":8080",
		Handler: Logging(http.DefaultServeMux),
	}
	log.Info("listening on 8080 ")
	log.Fatal(s.ListenAndServe())
}
