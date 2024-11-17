package websocket_service

import (
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
}

type WebSocketService struct {
	clients    map[string]*Client
	clientsMux sync.RWMutex
	broadcast  chan Message
}

type Message struct {
	SenderID   string `json:"sender_id"`
	ReceiverID string `json:"receiver_id"`
	Content    string `json:"content"`
}

func NewWebSocketService() *WebSocketService {
	return &WebSocketService{
		clients:   make(map[string]*Client),
		broadcast: make(chan Message),
	}
}

func (s *WebSocketService) AddClient(id string, conn *websocket.Conn) {
	s.clientsMux.Lock()
	defer s.clientsMux.Unlock()
	s.clients[id] = &Client{ID: id, Conn: conn}
}

func (s *WebSocketService) RemoveClient(id string) {
	s.clientsMux.Lock()
	defer s.clientsMux.Unlock()
	delete(s.clients, id)
}

func (s *WebSocketService) SendMessage(msg Message) {
	s.broadcast <- msg
}

func (s *WebSocketService) HandleMessages() {
	for {
		msg := <-s.broadcast
		s.clientsMux.RLock()
		if client, ok := s.clients[msg.ReceiverID]; ok {
			err := client.Conn.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Conn.Close()
				s.RemoveClient(msg.ReceiverID)
			}
		}
		s.clientsMux.RUnlock()
	}
}
