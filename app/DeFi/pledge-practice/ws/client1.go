package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/gorilla/websocket"
	"os"
	"time"
)

func main() {
	dialer := websocket.Dialer{}
	deadline, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	conn, _, err := dialer.DialContext(deadline, "ws://localhost:8080", nil)
	if err != nil {
		return
	}
	go wr(conn)
	for {

		_, p, err := conn.ReadMessage()
		if err != nil {
			return
		}

		fmt.Println(string(p) + " 你呢")
	}

}

/*
*
读出当前输入的值
*/
func wr(conn *websocket.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		//var data = make([]byte, 1024)
		TextMessage, _, err := reader.ReadLine()
		if err != nil {
			return
		}
		err = conn.WriteMessage(websocket.TextMessage, TextMessage)

	}
}
