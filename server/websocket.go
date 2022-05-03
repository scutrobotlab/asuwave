package server

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/scutrobotlab/asuwave/fromelf"
	"github.com/scutrobotlab/asuwave/logger"
)

func makeWebsocketCtrl(ch chan string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var upgrader = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
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

func fileWebsocketCtrl(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.Log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		select {
		case file := <-fromelf.ChFileModi:
			log.Println("ws got modified file:", file)
			err = c.WriteMessage(websocket.TextMessage, []byte(file))
			if err != nil {
				logger.Log.Println("write:", err)
				break
			}
		case err, ok := <-fromelf.ChFileError:
			if !ok {
				return
			}
			logger.Log.Println("error:", err)
		}
	}
}
