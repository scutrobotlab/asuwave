package server

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/scutrobotlab/asuwave/internal/option"
)

func optionCtrl(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	switch r.Method {
	case http.MethodGet:
		b, _ := json.Marshal(option.GetAll())
		io.WriteString(w, string(b))

	case http.MethodPut:
		j := struct {
			Key   string
			Value string
		}{}
		postData, _ := io.ReadAll(r.Body)
		if err := json.Unmarshal(postData, &j); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, errorJson("Invaild json"))
			return
		}
		switch j.Key {
		case "LogLevel":
			if v, err := strconv.Atoi(j.Value); err == nil && v >= 1 && v <= 5 {
				option.SetLogLevel(v)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				io.WriteString(w, errorJson("Invaild value"))
				return
			}
		case "SaveVarList":
			if v, err := strconv.ParseBool(j.Value); err == nil {
				option.SetSaveVarList(v)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				io.WriteString(w, errorJson("Invaild value"))
				return
			}
		case "SaveFilePath":
			if v, err := strconv.ParseBool(j.Value); err == nil {
				option.SetSaveFilePath(v)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				io.WriteString(w, errorJson("Invaild value"))
				return
			}
		case "UpdateByProj":
			if v, err := strconv.ParseBool(j.Value); err == nil {
				option.SetUpdateByProj(v)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				io.WriteString(w, errorJson("Invaild value"))
				return
			}
		default:
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, errorJson("Option Key Unfound."))
			return
		}
		w.WriteHeader(http.StatusNoContent)
		io.WriteString(w, "")

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, errorJson(http.StatusText(http.StatusMethodNotAllowed)))
	}
}
