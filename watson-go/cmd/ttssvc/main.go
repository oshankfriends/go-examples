package main

import (
	"github.com/oshankfriends/go-examples/watson-go"
	"github.com/oshankfriends/go-examples/watson-go/authentication"
	"io/ioutil"
	"net/http"
	"fmt"
)

func main() {
	client := watson_go.NewClient(http.DefaultClient, authentication.Credentials{
		"e8fd718e-34b0-47b8-8849-66194f3a9ed7",
		"Zt4J1WVmhbDO",
		"https://stream.watsonplatform.net/text-to-speech/api",
	})
	_, body, err := client.TTS.Synthesize("hello how can I help you", "en-US_AllisonVoice", "audio/wav")
	if err == nil {
		ioutil.WriteFile("sample.wav", body, 0666)
	}
	fmt.Println("Error : ",err)
}
