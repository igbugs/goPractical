package main

import (
	"net/http"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域的设置
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade websocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	// websocket.Conn
	for {
		_, data, err := conn.ReadMessage()
		if err != nil {
			goto ERR
		}
		err = conn.WriteMessage(websocket.TextMessage, data)
		if err != nil {
			goto ERR
		}
	}
	ERR:
		conn.Close()
}

func main()  {
	// http://localhost:7777/ws
	http.HandleFunc("/ws", wsHandler)

	http.ListenAndServe("0.0.0.0:7777", nil)
}
