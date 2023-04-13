package socket

import (
	"log"

	"github.com/PiamNaJa/CourseZ_Backend/models"
	"github.com/gofiber/websocket/v2"
)

type Server struct {
	conns map[*websocket.Conn]bool
}


func NewServer() *Server {
	return &Server{
		conns: make(map[*websocket.Conn]bool),
	}
}

func (s *Server) HandleConnection(c *websocket.Conn) {
	s.conns[c] = true
	for {
		var chat models.Conversation
		err := c.ReadJSON(&chat)
		if err != nil {
			s.conns[c] = false
			delete(s.conns, c)
			c.Close()
			break
		}
		log.Printf("Room: %d", chat.ChatRoom_id)
		log.Printf("User: %d", chat.Sender_id)
		log.Printf("Message: %s", chat.Message)
		s.Broadcast(chat)
	}
}

func (s *Server) Broadcast(msg models.Conversation) {
	for conn := range s.conns {
		err := conn.WriteJSON(msg)
		if err != nil {
			log.Printf("Error: %v", err)
			conn.Close()
			delete(s.conns, conn)
			break
		}
	}
}
