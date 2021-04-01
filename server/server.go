package server

import (
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"strconv"

	"github.com/scutrobotlab/asuwave/option"
	"github.com/scutrobotlab/asuwave/variable"
)

// Init server
func Init(c chan string, fsys *fs.FS) {
	host := "0.0.0.0:" + strconv.Itoa(option.Config.Port)
	log.Println("Listen on " + host)
	log.Println("Don't close this before you have done")

	variableToReadCtrl := makeVariableCtrl(&variable.ToRead, true)
	variableToModiCtrl := makeVariableCtrl(&variable.ToModi, false)
	websocketCtrl := makeWebsocketCtrl(c)

	http.Handle("/", http.FileServer(http.FS(*fsys)))

	http.HandleFunc("/serial", serialCtrl)
	http.HandleFunc("/serial_cur", serialCurCtrl)
	http.HandleFunc("/variable_read", variableToReadCtrl)
	http.HandleFunc("/variable_modi", variableToModiCtrl)
	http.HandleFunc("/variable_proj", variableToProjCtrl)
	http.HandleFunc("/variable_type", variableTypeCtrl)
	http.HandleFunc("/option", optionCtrl)
	http.HandleFunc("/ws", websocketCtrl)
	log.Fatal(http.ListenAndServe(host, nil))
}

func errorJson(s string) string {
	j := struct{ Error string }{s}
	b, _ := json.Marshal(j)
	return string(b)
}
