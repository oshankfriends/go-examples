package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"time"
)

const (
	boundry   = "---------------------------9051914041544843365972754266"
	chunkSize = 32
)

func main() {
	begin := time.Now()
	pr, pw := io.Pipe()
	fileContent, _ := ioutil.ReadFile("/home/oshank/Downloads/audio-file.wav")

	go func() {
		mltwr := multipart.NewWriter(pw)
		mltwr.SetBoundary(boundry)

		for start, end := 0, chunkSize; start < len(fileContent); start += chunkSize {
			if end > len(fileContent) {
				end = len(fileContent)
			}
			fmt.Println("creating part", start/chunkSize)
			h := make(textproto.MIMEHeader)
			h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "musicfile", "audio-file.wav"))
			h.Set("Content-Type", "audio/L16; rate=16000; channels=1")
			fileWr, _ := mltwr.CreatePart(h)
			fileWr.Write(fileContent[start:end])
			end += chunkSize
		}
		pw.Close()
	}()

	reqBegin := time.Now()
	resp, err := http.Post("http://localhost:8080/multi", "multipart/form-data; boundary="+boundry, pr)
	defer resp.Body.Close()
	fmt.Println(err, resp.Status)
	log.Infof("took total: %s, req time : %s", time.Since(begin), time.Since(reqBegin))
}
