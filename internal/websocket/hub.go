package websocket

import(
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool{
		return true
	},
}

func HandleConnections(w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil{
		fmt.Println("Websocker upgrade error:", err)
		return
	}

	clients[conn] = true 
	fmt.Println("Client connected")

	for {
		_, _, err := conn.ReadMessage()

		if err != nil{
			fmt.Println("Client disconnected")

			delete(clients, conn)

			conn.Close()

			break
		}
	}
}

func Broadcast(message []byte){
	for client := range clients {
		err := client.WriteMessage(
			websocket.TextMessage,
			message,
		)

		if err != nil{
			fmt.Println("Broadcast error:", err)

			client.Close()

			delete(clients, client)
		}
	}
}