package text_to_speech

import (
	"github.com/dghubble/sling"
	"github.com/gorilla/websocket"
	"github.com/oshankfriends/go-examples/watson-go/authentication"
	"log"
	"net/http"
	"time"
)

const defaultVersion = "v1"

type TTSService struct {
	Version string
	Sling   *sling.Sling
	Authset *authentication.Auth
}

func NewTTSService(baseSling *sling.Sling, authset *authentication.Auth) *TTSService {
	return &TTSService{
		Version: defaultVersion,
		Sling:   baseSling.New().Base("wss://stream.watsonplatform.net"),
		Authset: authset,
	}
}

func (tts *TTSService) Synthesize(text string, voice string, accept string) (*websocket.Conn, []byte, error) {
	var (
		body    []byte
		textMsg []string
	)
	tts.Sling.Path("/text-to-speech/api/" + tts.Version + "/synthesize")
	token, err := tts.Authset.GetToken()
	if err != nil {
		return nil, nil, err
	}
	tts.Sling.QueryStruct(&TTSQueryParams{voice, token})
	req, err := tts.Sling.Request()
	if err != nil {
		return nil, nil, err
	}
	log.Println("Dialing...", req.URL)
	conn, _, err := websocket.DefaultDialer.Dial(req.URL.String(), http.Header{"Origin": []string{"wss://stream.watsonplatform.net"}})
	if err != nil {
		return nil, nil, err
	}
	log.Println("Writing Json to Websocket...")
	if err = conn.WriteJSON(&TTSRequest{text, accept}); err != nil {
		return conn, nil, err
	}
	log.Println("Json written on WS")
	for {
		msgType, msgBody, err := conn.ReadMessage()
		if err != nil {
			time.Sleep(time.Second)
			log.Println(err)
			continue
		}
		if msgType == websocket.BinaryMessage {
			body = append(body, msgBody...)
		} else if msgType == websocket.TextMessage {
			textMsg = append(textMsg, string(msgBody))
		}
	}
}
