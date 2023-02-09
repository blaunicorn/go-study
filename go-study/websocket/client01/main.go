package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	fmt.Println("client01")
	dialer := websocket.Dialer{}
	conn, _, err := dialer.Dial("ws://127.0.0.1:8888", nil)
	if err != nil {
		log.Println(err)
		return
	}
	conn.WriteMessage(websocket.TextMessage, []byte("I will back."))
	go send(conn)
	for {
		messageType, p, e := conn.ReadMessage()
		if e != nil {
			break
		}
		fmt.Println(messageType, string(p), e)
	}

}

func send(conn *websocket.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		line, _, _ := reader.ReadLine()
		conn.WriteMessage(websocket.TextMessage, line)
	}
}
