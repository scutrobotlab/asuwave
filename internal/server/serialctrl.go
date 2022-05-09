package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/scutrobotlab/asuwave/internal/serial"
)

type SerialSetting struct {
	Serial string
	Baud   int
}

func serialCtrl(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		j := struct{ Serials []string }{Serials: serial.Find()}
		b, _ := json.Marshal(j)
		io.WriteString(w, string(b))

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, errorJson(http.StatusText(http.StatusMethodNotAllowed)))
	}
}

func serialCurCtrl(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	var err error

	switch r.Method {
	case http.MethodGet:
		j := SerialSetting{
			Serial: serial.SerialCur.Name,
			Baud:   serial.SerialCur.Mode.BaudRate,
		}
		b, _ := json.Marshal(j)
		io.WriteString(w, string(b))

	case http.MethodPost:
		j := SerialSetting{}
		postData, _ := io.ReadAll(r.Body)
		err = json.Unmarshal(postData, &j)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, errorJson("Invaild json"))
			return
		}

		err = serial.Open(j.Serial, j.Baud)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, errorJson(err.Error()))
			return
		}
		io.WriteString(w, string(postData))

	case http.MethodDelete:
		err = serial.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, errorJson(err.Error()))
			return
		}
		w.WriteHeader(http.StatusNoContent)
		io.WriteString(w, "")

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, errorJson(http.StatusText(http.StatusMethodNotAllowed)))
	}
}
