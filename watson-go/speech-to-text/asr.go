package speech_to_text

import (
	"errors"
	"github.com/dghubble/sling"
	"github.com/gorilla/websocket"
	"github.com/oshankfriends/go-examples/watson-go/authentication"
	"io"
	"net/http"
	"net/url"
	"github.com/sirupsen/logrus"
)

const defaultVersion = "v1"

type ASRService struct {
	Version string
	sling   *sling.Sling
	AuthSet *authentication.Auth
	dialer  *websocket.Dialer
}

func NewASRService(baseSling *sling.Sling, authSet *authentication.Auth) *ASRService {
	return &ASRService{
		Version: defaultVersion,
		sling:   baseSling,
		AuthSet: authSet,
		dialer:  websocket.DefaultDialer,
	}
}

func (asr *ASRService) getDefaultHeaders() http.Header {
	var h = http.Header{
		"Content-Type": []string{"application/json"},
		"User-Agent":   []string{"watson-go/1.0.0"},
	}
	return h
}

func (asr *ASRService) Stream(message map[string]interface{}, contentType string, model string, inactivityTimeout int) (<-chan SpeechRecognizeResponse, io.WriteCloser, error) {
	token, err := asr.AuthSet.GetToken()
	if err != nil {
		return nil, nil, err
	}
	uri, err := url.Parse(asr.AuthSet.Creds.Url)
	if err != nil {
		return nil, nil, err
	}
	headers := asr.getDefaultHeaders()
	headers["Origin"] = []string{uri.String()}

	uri.Scheme = "wss"
	uri.Path += "/" + asr.Version + "/recognize"
	queryValues := &url.Values{}
	queryValues.Set("watson-token", token)
	queryValues.Set("model", model)
	uri.RawQuery = queryValues.Encode()

	logrus.Debug("dialing...", uri)
	conn, _, err := asr.dialer.Dial(uri.String(), headers)
	if err != nil {
		return nil, nil, err
	}
	logrus.Debugf("Websocket Conn created Local addr : %s, Remote Addr : %s \n", conn.LocalAddr(), conn.RemoteAddr())
	respChan := make(chan SpeechRecognizeResponse, 10)
	s := &Stream{
		EventChan:         respChan,
		Conn:              conn,
		Message:           message,
		InactivityTimeout: inactivityTimeout,
		ContentType:       contentType,
	}
	return respChan, s, nil
}

type Stream struct {
	EventChan         chan<- SpeechRecognizeResponse
	Conn              *websocket.Conn
	Message           map[string]interface{}
	ContentType       string
	InactivityTimeout int
	Started           bool
	Stopped           bool
}

func (s *Stream) Write(b []byte) (int, error) {
	if s.Stopped {
		return 0, errors.New("can not write on stopped stream")
	}

	if !s.Started {
		msg := make(map[string]interface{})
		for k, v := range s.Message {
			msg[k] = v
		}
		msg["action"] = "start"
		msg["content-type"] = s.ContentType
		msg["interim_results"] = true
		msg["inactivity_timeout"] = s.InactivityTimeout

		if err := s.Conn.WriteJSON(msg); err != nil {
			return 0, err
		}
		s.Started = true
		go s.readResponse()
	}
	err := s.Conn.WriteMessage(websocket.BinaryMessage, b)
	return len(b), err
}

func (s *Stream) readResponse() {
	for {
		var speechResp SpeechRecognizeResponse
		if err := s.Conn.ReadJSON(&speechResp); err != nil {
			logrus.Errorf("Stream.readResponse err : %v",err)
			if wsErr,ok := err.(*websocket.CloseError); ok && wsErr.Code == 1011 {
				close(s.EventChan)
				return
			}
			continue
		}
		if s.Stopped {
			close(s.EventChan)
			return
		}
		s.EventChan <- speechResp
	}
}

func (s *Stream) Close() error {
	if s.Stopped {
		return errors.New("can not close a stopped stream")
	}
	s.Stopped = true
	m := map[string]interface{}{
		"action": "stop",
	}
	return s.Conn.WriteJSON(m)
}
