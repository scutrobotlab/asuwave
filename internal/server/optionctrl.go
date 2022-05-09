package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/scutrobotlab/asuwave/internal/option"
)

func optionCtrl(w http.ResponseWriter, r *http.Request) {
	defer option.Refresh()
	defer option.Save()
	defer r.Body.Close()
	var err error

	switch r.Method {
	case http.MethodGet:
		j := struct{ Save int }{Save: option.Config.Save}
		b, _ := json.Marshal(j)
		io.WriteString(w, string(b))

	case http.MethodPut:
		j := struct {
			Save int
		}{}
		postData, _ := io.ReadAll(r.Body)
		err = json.Unmarshal(postData, &j)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, errorJson("Invaild json"))
			return
		}
		option.Config.Save = j.Save
		io.WriteString(w, string(postData))

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, errorJson(http.StatusText(http.StatusMethodNotAllowed)))
	}
}
