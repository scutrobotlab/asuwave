package server

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"net"
	"net/http"
	"strconv"

	"github.com/scutrobotlab/asuwave/logger"
	"github.com/scutrobotlab/asuwave/option"
	"github.com/scutrobotlab/asuwave/variable"
)

// Start server
func Start(c chan string, fsys *fs.FS) {
	port := ":" + strconv.Itoa(option.Config.Port)

	fmt.Println("asuwave running at:")
	fmt.Println("- Local:   http://localhost" + port + "/")
	fmt.Println("- Network: http://" + getLocalIP() + port + "/")
	fmt.Println("Don't close this before you have done")

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
	logger.Log.Fatal(http.ListenAndServe(port, nil))
}

func errorJson(s string) string {
	j := struct{ Error string }{s}
	b, _ := json.Marshal(j)
	return string(b)
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			return ipnet.IP.String()
		}
	}

	return "127.0.0.1"
}
