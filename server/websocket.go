package server

import (
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/scutrobotlab/asuwave/logger"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func makeWebsocketCtrl(ch chan string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			logger.Log.Print("upgrade:", err)
			return
		}
		defer c.Close()
		for {
			b := <-ch
			err = c.WriteMessage(websocket.TextMessage, []byte(b))
			if err != nil {
				logger.Log.Println("write:", err)
				break
			}
		}
	}
}
