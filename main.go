package main

import (
	apis "imp-slipped-away-server/apis"
)

// var _clients = make(map[*websocket.Conn]bool)
// var _broadcast = make(chan Message)

// var _upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func main(){
	//go startWebSocket()

	apis.Init()
}

// func startWebSocket() {
// 	http.HandleFunc("/ws", handleConnections)
// 	go handleMessages()

// 	log.Println("http server started on :9000")
// 	err := http.ListenAndServe(":9000", nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }

// func handleConnections(w http.ResponseWriter, r *http.Request) {
// 	ws, err := _upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer ws.Close()

// 	_clients[ws] = true

// 	for {
// 		var msg Message
// 		err := ws.ReadJSON(&msg)
// 		if err != nil {
// 			log.Printf("error: %v", err)
// 			delete(_clients, ws)
// 			break
// 		}
// 		_broadcast <- msg        
// 	}
// }

// func handleMessages() {
// 	for {
// 		msg := <- _broadcast
// 		for client := range _clients {
// 			err := client.WriteJSON(msg)
// 			if err != nil {
// 				log.Printf("error: %v", err)
// 				client.Close()
// 				delete(_clients, client)
// 			}
// 		}
// 	}
// }