package boostrap

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
}

var clients = make(map[string]*Client)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func HandleWebSocket(c *gin.Context) {
	userID := c.Param("userID")
	fmt.Println("userID: ", userID)
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	client := &Client{ID: userID, Conn: conn}

	clients[userID] = client

	defer func() {
		conn.Close()
		delete(clients, userID)
	}()

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}

		var message struct {
			To      string `json:"to"`
			Content string `json:"content"`
		}

		if err := json.Unmarshal(p, &message); err != nil {
			continue
		}

		if targetClient, ok := clients[message.To]; ok {
			targetClient.Conn.WriteMessage(messageType, p)
		}
	}
}
