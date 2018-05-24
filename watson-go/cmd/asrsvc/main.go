package main

import (
	"fmt"
	"github.com/oshankfriends/go-examples/watson-go"
	"github.com/oshankfriends/go-examples/watson-go/authentication"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	client := watson_go.NewClient(http.DefaultClient, authentication.Credentials{
		Url:      "https://stream.watsonplatform.net/speech-to-text/api",
		Username: "2a602575-bbb6-4c07-a8e6-05913cc69b6d",
		Password: "cjCn4hOC1JKa",
	})
	eventStream, writer, err := client.Asr.Stream(map[string]interface{}{"continuous": true, "interim_results": false, "timestamps": false},
		"audio/wav", "en-US_BroadbandModel", 1)
	if err != nil {
		log.Print("stream error :", err)
		return
	}
	file, err := os.Open("/home/oshank/Downloads/male.wav")
	if err != nil {
		log.Println("file open err :", err)
		return
	}
	if _, err = io.Copy(writer, file); err != nil {
		log.Println("io.Copy() failed to copy audio file to API :", err)
		return
	}

	for event := range eventStream {
		if len(event.Results) > 0 && event.Results[0].Alternatives[0].Confidence > 0.8 {
			writer.Close()
		}
		fmt.Printf("%+v\n", event.Results)
	}

}
