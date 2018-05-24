package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/prometheus/common/log"
	"github.com/stretchr/objx"
	"net/http"
)

var upgrader = &websocket.Upgrader{}

type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

type Server struct {
	Router  *http.ServeMux
	message chan *Message
	errc    chan chan error
	clients map[*Client]bool
}

type Client struct {
	conn  *websocket.Conn
	Name  string
	Email string
}

func NewServer(router *http.ServeMux) *Server {
	s := &Server{
		Router:  router,
		message: make(chan *Message),
		errc:    make(chan chan error),
		clients: make(map[*Client]bool),
	}
	go s.process()
	return s
}

func (s *Server) Listen() {
	log.Info("websocket server listening")
	s.Router.HandleFunc("/chat", s.HandleChat)
}

func (s *Server) Close() error {
	errch := make(chan error)
	s.errc <- errch
	return <-errch
}

func (s *Server) process() {
	var err error
	for {
		select {
		case msg := <-s.message:
			log.Infof("msg recieved %+v", msg)
			for client := range s.clients {
				if err = client.conn.WriteJSON(msg); err != nil {
					log.Info("got error in writing to webSocket", err)
				}
			}
		case errch := <-s.errc:
			for client := range s.clients {
				log.Info("closing", client.conn.RemoteAddr())
				err = client.conn.Close()
			}
			errch <- err
		}
	}
}

func (s *Server) HandleChat(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Infof("upgrade err : %s", err)
		return
	}
	log.Info("connection upgraded", conn.LocalAddr())

	client := &Client{conn: conn}
	if cookie, err := r.Cookie("auth"); err == nil {
		userData := objx.MustFromBase64(cookie.Value)
		client.Name = userData.Get("name").String()
		client.Email = userData.Get("email").String()
	}
	s.clients[client] = true

	for {
		var msg = &Message{}
		if err = conn.ReadJSON(msg); err != nil {
			log.Infof("err in reading msg : %s, deleting the client %v", err, client.Name)
			delete(s.clients, client)
			return
		}
		s.message <- msg
	}
}
