package asrctl

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	watson "github.com/oshankfriends/go-examples/watson-go"
	"github.com/oshankfriends/go-examples/watson-go/authentication"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"io"
	"net/http"
	"os"
)

func NewAsrCommand() *cobra.Command {
	var logLevel string
	o := &AsrOption{}
	cmd := &cobra.Command{
		Use:   "asr",
		Short: "covert speech to text",
		Long:  "asrctl command is use to convert any audio file in text format",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if lvl, err := logrus.ParseLevel(logLevel); err == nil {
				logrus.SetLevel(lvl)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			o.Complete(cmd, args)
			if err := o.Run(); err != nil {
				logrus.Error(err)
			}
		},
	}
	cmd.Flags().StringVarP(&o.UserName, "username", "u", "2a602575-bbb6-4c07-a8e6-05913cc69b6d", "Username for accessing watson api")
	cmd.Flags().StringVarP(&o.PassWord, "password", "p", "cjCn4hOC1JKa", "Password for accessing watson api, for more info visit https://console.bluemix.net/catalog/services/speech-to-text")
	cmd.Flags().StringVarP(&logLevel, "loglevel", "l", "info", "logging level")
	cmd.Flags().StringVarP(&o.InputFile, "input", "i", "/home/oshank/Downloads/a04.pcm", "input audio file which will use to convert into text")
	cmd.Flags().StringVarP(&o.ContentType, "content-type", "c", "audio/L16; rate=16000; channels=1", "audio content type")
	cmd.Flags().IntVarP(&o.InactivityTimeout, "inactivity-timeout", "t", -1, "The time in seconds after which, if only silence (no speech) is detected in submitted audio,the connection is closed.")
	return cmd
}

type AsrOption struct {
	InputFile         string
	ContentType       string
	InactivityTimeout int
	Client            *watson.Client
	UserName          string
	PassWord          string
}

func (a *AsrOption) Complete(cmd *cobra.Command, args []string) error {
	a.Client = watson.NewClient(http.DefaultClient, authentication.Credentials{
		Url:      "https://stream.watsonplatform.net/speech-to-text/api",
		Username: a.UserName,
		Password: a.PassWord,
	})
	return nil
}

func (a *AsrOption) Run() error {
	eventStream, writer, err := a.Client.Asr.Stream(map[string]interface{}{"continuous": true, "interim_results": false, "timestamps": false},
		a.ContentType, "en-US_BroadbandModel", a.InactivityTimeout)
	if err != nil {
		return err
	}
	file, err := os.Open(a.InputFile)
	if err != nil {
		return err
	}
	if _, err = io.Copy(writer, file); err != nil {
		return err
	}
	table := tablewriter.NewWriter(os.Stdout)
	defer table.Render()
	table.SetHeader([]string{"Transcript", "Confidence"})
	for event := range eventStream {
		if len(event.Results) > 0 {
			table.Append([]string{event.Results[0].Alternatives[0].Transcript, fmt.Sprintf("%f", event.Results[0].Alternatives[0].Confidence)})
		}
		logrus.Debugf("%+v", event)
	}
	return nil
}
