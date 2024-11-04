package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type webSocketServer struct {
	Upgrader  *websocket.Upgrader
	Clients   map[*websocket.Conn]bool
	Broadcast chan []byte
}

func newWebSocketServer() *webSocketServer {
	return &webSocketServer{
		Upgrader: &websocket.Upgrader{},
		Clients:  make(map[*websocket.Conn]bool),
		Broadcast: make(chan []byte),
	}
}

func (w *webSocketServer) handleConnections( writer http.ResponseWriter, r *http.Request) {
	for {

		conn, err := w.Upgrader.Upgrade(writer, r, nil)
		if err != nil {
			fmt.Println("Error upgrading to websocket: ", err)
			return
		}
		fmt.Println("Client connected")


	
		w.Clients[conn] = true
		fmt.Println("Number of clients connected: ", len(w.Clients))
		go w.handleMessages(conn)
	}
}

func (w *webSocketServer) handleMessages(conn *websocket.Conn) {
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message: ", err)
			conn.Close()
			delete(w.Clients, conn)
			return
		}
		w.Broadcast <- msg
	}
}

func (w *webSocketServer) broadcastMessages() {
	for {


		msg := <-w.Broadcast

		for client := range w.Clients {

			err := client.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				fmt.Print("Error: ", err)
				client.Close()
				delete(w.Clients, client)
			}
		}
	}
}



func main() {
	server := newWebSocketServer()
	http.HandleFunc("/ws", server.handleConnections)
	fmt.Println("Server started on :8080")
	go server.broadcastMessages()
	err:=http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("Error starting server: ", err)
		return
	}

}