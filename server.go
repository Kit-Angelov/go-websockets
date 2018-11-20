package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// var clients = make(map[*websocket.Conn]bool)
var clients = make(map[string]*websocket.Conn)

// var broadcast = make(chan Message)
var message_queue = make(chan Message)

var upgrader = websocket.Upgrader{}

type Message struct {
	UserId string
}

type NewMessage struct {
	Message string
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// defer ws.Close()

	// clients[ws] = true
	println("connect")

	var msg Message
	err = ws.ReadJSON(&msg)
	println(msg.UserId)
	if err != nil {
		log.Printf("error: %v", err)
		// delete(clients, ws)
	}
	clients[msg.UserId] = ws
}

func handleMessages() {
	for {
		println("WAIT")
		time.Sleep(2 * time.Second)
		message := NewMessage{Message: "HEELLOO"}

		userId := "2"

		client, ok := clients[userId]

		if ok == false {
			continue
		}

		err := client.WriteJSON(message)
		if err != nil {
			println("error write json")
			client.Close()
			// delete(clients, userId)
		}
		// msg := <-broadcast
		// for client := range clients {
		// 	err := client.WriteJSON(msg)
		// 	if err != nil {
		// 		log.Printf("error: %v", err)
		// 		client.Close()
		// 		delete(clients, client)
		// 	}
		// }
	}
}

func main() {
	fs := http.FileServer(http.Dir(""))
	http.Handle("/", fs)
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()
	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
