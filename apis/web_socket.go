package apis

import (
	"log"
	"net/http"
	"github.com/labstack/echo"
	"github.com/gorilla/websocket"
)

type Message struct {
	Id int64 `json:"id"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

var _clients = make(map[*websocket.Conn]bool)
var _broadcast = make(chan Message)

var _upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func InitWebSocket(e *echo.Echo){
	e.GET("/ws", handleConnections)
}

func handleConnections(c echo.Context) error {
	ws, err := _upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	_clients[ws] = true

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(_clients, ws)
			break
		}
		_broadcast <- msg        
	}
	
	return err
}

func handleMessages() {
	for {
		msg := <- _broadcast
		for client := range _clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(_clients, client)
			}
		}
	}
}