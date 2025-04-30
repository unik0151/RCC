package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

func main() {

	http.HandleFunc("/", handlerWs)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return
	}
}

/*
*
升级长连接 ， 如果断掉则err
*/
func handlerWs(writer http.ResponseWriter, request *http.Request) {
	conn, err := Up1.Upgrade(writer, request, nil)
	if err != nil {
		return
	}

	defer func() {
		_ = conn.Close()
		fmt.Println("close finished! ")
	}()
	/**
	读取发送的消息
	*/
	for {
		messageType, bytes, errcode := conn.ReadMessage()
		if errcode != nil {
			return
		}
		fmt.Printf("messageType=%d,bytes=%s\n", messageType, string(bytes))
	}

}

var Up1 = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
