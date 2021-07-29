package server

import (
	"encoding/json"
	"fmt"
	"github.com/scutrobotlab/asuwave/datautil"
	"github.com/scutrobotlab/asuwave/fromelf"
	"github.com/scutrobotlab/asuwave/serial"
	"github.com/scutrobotlab/asuwave/variable"
	"io"
	"net/http"
	"os"
	"sort"
)

func makeVariableCtrl(vList *variable.ListT, isVToRead bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer variable.Refresh()
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json")
		var err error
		switch r.Method {
		case http.MethodGet:
			b, _ := json.Marshal(vList)
			io.WriteString(w, string(b))

		case http.MethodPost:
			var newVariable variable.T
			postData, _ := io.ReadAll(r.Body)
			err = json.Unmarshal(postData, &newVariable)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				io.WriteString(w, errorJson("Invaild json"))
				return
			}

			for _, v := range vList.Variables {
				if v.Addr == newVariable.Addr {
					w.WriteHeader(http.StatusBadRequest)
					io.WriteString(w, errorJson("Address already used"))
					return
				}
			}
			if isVToRead && serial.SerialCur.Name != "" {
				err = serial.SendCmd(datautil.ActModeSubscribe, newVariable)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					io.WriteString(w, errorJson(err.Error()))
					return
				}
			}
			vList.Variables = append(vList.Variables, newVariable)
			w.WriteHeader(http.StatusNoContent)
			io.WriteString(w, "")

		case http.MethodPut:
			if isVToRead {
				w.WriteHeader(http.StatusMethodNotAllowed)
				io.WriteString(w, errorJson(http.StatusText(http.StatusMethodNotAllowed)))
				return
			}
			var modVariable variable.T
			postData, _ := io.ReadAll(r.Body)
			err = json.Unmarshal(postData, &modVariable)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				io.WriteString(w, errorJson("Invaild json"))
				return
			}
			if serial.SerialCur.Name != "" {
				err = serial.SendCmd(datautil.ActModeWrite, modVariable)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					io.WriteString(w, errorJson(err.Error()))
					return
				}
			}
			w.WriteHeader(http.StatusNoContent)
			io.WriteString(w, "")

		case http.MethodDelete:
			var oldVariable variable.T
			postData, _ := io.ReadAll(r.Body)
			err = json.Unmarshal(postData, &oldVariable)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				io.WriteString(w, errorJson("Invaild json"))
				return
			}
			for i, v := range vList.Variables {
				if v.Addr == oldVariable.Addr {
					if isVToRead && serial.SerialCur.Name != "" {
						err = serial.SendCmd(datautil.ActModeUnSubscribe, oldVariable)
						if err != nil {
							w.WriteHeader(http.StatusInternalServerError)
							io.WriteString(w, errorJson(err.Error()))
							return
						}
					}

					vList.Variables = append(vList.Variables[:i], vList.Variables[i+1:]...)
					w.WriteHeader(http.StatusNoContent)
					io.WriteString(w, "")
					return
				}
			}
			w.WriteHeader(http.StatusBadRequest)
			io.WriteString(w, errorJson("No such address"))

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			io.WriteString(w, errorJson(http.StatusText(http.StatusMethodNotAllowed)))
		}
	}
}

func variableToProjCtrl(w http.ResponseWriter, r *http.Request) {
	defer variable.Refresh()
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		b, _ := json.Marshal(variable.ToProj)
		io.WriteString(w, string(b))

	case http.MethodPost:
		r.ParseMultipartForm(32 << 20)
		file, _, err := r.FormFile("file")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, errorJson(err.Error()))
			return
		}
		defer file.Close()

		tempFile, err := os.CreateTemp("", "elf")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, errorJson(err.Error()))
			return
		}
		defer os.Remove(tempFile.Name())

		io.Copy(tempFile, file)

		f, err := fromelf.Check(tempFile)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, errorJson(err.Error()))
			return
		}
		defer f.Close()

		err = fromelf.ReadVariable(&variable.ToProj, f)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			io.WriteString(w, errorJson(err.Error()))
			return
		}

		w.WriteHeader(http.StatusNoContent)
		io.WriteString(w, "")
	case http.MethodDelete:
		fmt.Println("good")
		w.WriteHeader(http.StatusNoContent)
		io.WriteString(w, "")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, errorJson(http.StatusText(http.StatusMethodNotAllowed)))
	}
}

func variableTypeCtrl(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		var types struct{ Types []string }
		for k := range variable.TypeLen {
			types.Types = append(types.Types, k)
		}
		sort.Strings(types.Types)
		b, _ := json.Marshal(types)
		io.WriteString(w, string(b))

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		io.WriteString(w, errorJson(http.StatusText(http.StatusMethodNotAllowed)))
	}
}
