package server

import (
	"net/http"
	"testing"
)

func TestOptionCtrl(t *testing.T) {
	cases := casesT{
		{
			http.MethodGet,
			"/options",
			nil,
			http.StatusOK,
		},
		{
			http.MethodPost,
			"/options",
			nil,
			http.StatusMethodNotAllowed,
		},
		{
			http.MethodPut,
			"/options",
			struct {
				Key   string
				Value int
			}{
				Key:   "LogLevel",
				Value: 3,
			},
			http.StatusOK,
		},
		{
			http.MethodDelete,
			"/options",
			nil,
			http.StatusMethodNotAllowed,
		},
		{
			http.MethodGet,
			"/options",
			nil,
			http.StatusOK,
		},
	}
	ctrlerTest(optionCtrl, cases, t)
}
