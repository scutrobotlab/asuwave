package server

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"

	"github.com/scutrobotlab/asuwave/internal/serial"
	"github.com/scutrobotlab/asuwave/internal/variable"
)

// vList 要控制的参数列表；
// isVToRead 为true代表只读变量，为false代表可写变量
func makeVariableCtrl(m variable.Mod, isVToRead bool) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		w.Header().Set("Content-Type", "application/json")
		var err error
		switch r.Method {
		// 获取变量列表
		case http.MethodGet:
			b, _ := variable.GetAll(m)
			io.WriteString(w, string(b))
		// 新增变量
		case http.MethodPost:
			var newVariable variable.T
			postData, _ := io.ReadAll(r.Body)
			err = json.Unmarshal(postData, &newVariable)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				io.WriteString(w, errorJson("Invaild json"))
				return
			}
			if newVariable.Addr < 0x20000000 || newVariable.Addr >= 0x80000000 {
				w.WriteHeader(http.StatusBadRequest)
				io.WriteString(w, errorJson("Address out of range"))
				return
			}
			if _, ok := variable.Get(m, newVariable.Addr); ok {
				w.WriteHeader(http.StatusBadRequest)
				io.WriteString(w, errorJson("Address already used"))
				return
			}
			variable.Set(m, newVariable.Addr, newVariable)
			w.WriteHeader(http.StatusNoContent)
			io.WriteString(w, "")
		// 为变量赋值
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
			if serial.SerialCur.Name == "" {
				w.WriteHeader(http.StatusInternalServerError)
				io.WriteString(w, "Not allow when serial port closed.")
				return
			}
			err = serial.SendWriteCmd(modVariable)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				io.WriteString(w, errorJson(err.Error()))
				return
			}
			w.WriteHeader(http.StatusNoContent)
			io.WriteString(w, "")
		// 删除变量
		case http.MethodDelete:
			var oldVariable variable.T
			postData, _ := io.ReadAll(r.Body)
			err = json.Unmarshal(postData, &oldVariable)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				io.WriteString(w, errorJson("Invaild json"))
				return
			}

			// 我认为不必再检查是否存在这个变量
			// if _, ok := vList.Variables[oldVariable.Addr]; !ok {
			// 	w.WriteHeader(http.StatusBadRequest)
			// 	io.WriteString(w, errorJson("No such address"))
			// }

			// 从 vList.Variables 中删除地址为 oldVariable.Addr 的变量
			variable.Delete(m, oldVariable.Addr)
			w.WriteHeader(http.StatusNoContent)
			io.WriteString(w, "")
			return

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			io.WriteString(w, errorJson(http.StatusText(http.StatusMethodNotAllowed)))
		}
	}
}

// 工程变量
func variableToProjCtrl(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case http.MethodGet:
		b, _ := variable.GetAllProj()
		io.WriteString(w, string(b))
	case http.MethodDelete:
		variable.SetAllProj(variable.Projs{})

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
