package server

import (
	"net/http"

	"github.com/golang/glog"
	"github.com/gorilla/websocket"

	"github.com/scutrobotlab/asuwave/pkg/elffile"
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
			glog.Errorln("upgrade:", err)
			return
		}
		defer c.Close()
		for {
			b := <-ch
			err = c.WriteMessage(websocket.TextMessage, []byte(b))
			if err != nil {
				glog.Errorln("write:", err)
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
		glog.Errorln("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		select {
		case file := <-elffile.ChFileWrite:
			glog.Infoln("ws got modified file:", file)
			err = c.WriteMessage(websocket.TextMessage, []byte(file))
			if err != nil {
				glog.Errorln("write:", err)
				break
			}
		case err, ok := <-elffile.ChFileError:
			if !ok {
				return
			}
			glog.Errorln("error:", err)
		}
	}
}
