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
	logger.Log.Println("Listen on " + port)

	fmt.Println("asuwave running at:")
	fmt.Println("- Local:   http://localhost" + port + "/")
	ips := getLocalIP()
	for _, ip := range ips {
		fmt.Println("- Network: http://" + ip + port + "/")
	}
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
	http.HandleFunc("/file/upload", fileUploadCtrl)
	http.HandleFunc("/file/path", filePathCtrl)
	http.HandleFunc("/option", optionCtrl)
	http.HandleFunc("/ws", websocketCtrl)
	http.HandleFunc("/filews", fileWebsocketCtrl)
	logger.Log.Fatal(http.ListenAndServe(port, nil))
}

func errorJson(s string) string {
	j := struct{ Error string }{s}
	b, _ := json.Marshal(j)
	return string(b)
}

func getLocalIP() []string {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}

	return ips
}
