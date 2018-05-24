package watson_go

import (
	"github.com/dghubble/sling"
	"github.com/oshankfriends/go-examples/watson-go/authentication"
	asr "github.com/oshankfriends/go-examples/watson-go/speech-to-text"
	tts "github.com/oshankfriends/go-examples/watson-go/text-to-speech"
	"net/http"
)

type Client struct {
	Sling *sling.Sling
	TTS   *tts.TTSService
	Asr   *asr.ASRService
	Auth  *authentication.Auth
}

func NewClient(client *http.Client, creds authentication.Credentials) *Client {
	baseSling := sling.New().Client(client).Base("https://stream.watsonplatform.net")
	auth := authentication.NewAuth(baseSling.New(), creds)
	return &Client{
		Sling: baseSling,
		TTS:   tts.NewTTSService(baseSling.New(), auth),
		Asr:   asr.NewASRService(baseSling, auth),
		Auth:  auth,
	}
}
