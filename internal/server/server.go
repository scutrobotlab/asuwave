package server

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"mime"
	"net"
	"net/http"
	"strconv"

	"github.com/golang/glog"
	"github.com/scutrobotlab/asuwave/internal/option"
	"github.com/scutrobotlab/asuwave/internal/variable"
)

// Start server
func Start(c chan string, fsys *fs.FS) {
	port := ":" + strconv.Itoa(option.Config.Port)
	glog.Infoln("Listen on " + port)

	fmt.Println("asuwave running at:")
	fmt.Println("- Local:   http://localhost" + port + "/")
	ips := getLocalIP()
	for _, ip := range ips {
		fmt.Println("- Network: http://" + ip + port + "/")
	}
	fmt.Println("Don't close this before you have done")

	variableToReadCtrl := makeVariableCtrl(variable.Read, true)
	variableToWriteCtrl := makeVariableCtrl(variable.Write, false)
	websocketCtrl := makeWebsocketCtrl(c)

	mime.AddExtensionType(".js", "application/javascript")
	http.Handle("/", http.FileServer(http.FS(*fsys)))

	http.Handle("/serial", logs(serialCtrl))
	http.Handle("/serial_cur", logs(serialCurCtrl))
	http.Handle("/variable_read", logs(variableToReadCtrl))
	http.Handle("/variable_write", logs(variableToWriteCtrl))
	http.Handle("/variable_proj", logs(variableToProjCtrl))
	http.Handle("/variable_type", logs(variableTypeCtrl))
	http.Handle("/file/upload", logs(fileUploadCtrl))
	http.Handle("/file/path", logs(filePathCtrl))
	http.Handle("/option", logs(optionCtrl))
	http.Handle("/ws", logs(websocketCtrl))
	http.Handle("/filews", logs(fileWebsocketCtrl))
	glog.Fatalln(http.ListenAndServe(port, nil))
}

func logs(f func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		glog.Infoln("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		http.HandlerFunc(f).ServeHTTP(w, r)
	})
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
