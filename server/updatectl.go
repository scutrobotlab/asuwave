package server

import (
	"io"
	"net/http"

	"github.com/scutrobotlab/asuwave/helper"
)

func updateCtrl(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	// var err error

	switch r.Method {
	case http.MethodPost:
		helper.CheckUpdate(true)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, errorJson(http.StatusText(http.StatusMethodNotAllowed)))
	}
}
