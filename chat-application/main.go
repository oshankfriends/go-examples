package main

import (
	"fmt"
	"github.com/oshankfriends/go-examples/chat-application/routers"
	"github.com/oshankfriends/go-examples/chat-application/websocket"
	"github.com/prometheus/common/log"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/gomniauth/providers/google"
	"github.com/stretchr/signature"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	router := routers.New()
	errc := make(chan error)

	ws := websocket.NewServer(router)
	ws.Listen()
	defer ws.Close()

	go func() {
		sigChan := make(chan os.Signal)
		signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)
		errc <- fmt.Errorf("%s", <-sigChan)
	}()

	go func() {
		log.Info("http server listening on port 8080")
		errc <- http.ListenAndServe(":8080", router)
	}()

	log.Error(<-errc)
}

func init() {
	gomniauth.SetSecurityKey(signature.RandomKey(64))
	gomniauth.WithProviders(
		google.New(
			"489037682004-73m4673uk1ngt92bshh63dk6n7aisddr.apps.googleusercontent.com",
			"29PEoFO_Pd-ajYw5eAn5nD23",
			"http://localhost:8080/auth/callback/google",
		),
	)
}
