package server

import (
	"net/http"
	"testing"

	"github.com/scutrobotlab/asuwave/internal/serial"
	"github.com/scutrobotlab/asuwave/internal/variable"
)

func TestVariableToReadCtrl(t *testing.T) {
	cases := casesT{
		{
			http.MethodGet,
			"/variable_read",
			nil,
			http.StatusOK,
		},
		{
			http.MethodPost,
			"/variable_read",
			struct {
				Board uint8
				Name  string
				Type  string
				Addr  uint32
			}{
				Board: 1,
				Name:  "a",
				Type:  "int",
				Addr:  0x20123456,
			},
			http.StatusNoContent,
		},
		{
			http.MethodGet,
			"/variable_read",
			nil,
			http.StatusOK,
		},
		{
			http.MethodPut,
			"/variable_read",
			nil,
			http.StatusMethodNotAllowed,
		},
		{
			http.MethodDelete,
			"/variable_read",
			struct {
				Board uint8
				Name  string
				Type  string
				Addr  uint32
			}{
				Board: 1,
				Name:  "a",
				Type:  "int",
				Addr:  0x20123456,
			},
			http.StatusNoContent,
		},
	}

	variableToReadCtrl := makeVariableCtrl(variable.Read)

	ctrlerTest(variableToReadCtrl, cases, t)
}

func TestVariableToWriteCtrl(t *testing.T) {
	serial.Open("Test port", 115200)
	cases := casesT{
		{
			http.MethodGet,
			"/variable_write",
			nil,
			http.StatusOK,
		},
		{
			http.MethodPost,
			"/variable_write",
			struct {
				Board uint8
				Name  string
				Type  string
				Addr  uint32
			}{
				Board: 1,
				Name:  "a",
				Type:  "double",
				Addr:  0x20123456,
			},
			http.StatusNoContent,
		},
		{
			http.MethodGet,
			"/variable_write",
			nil,
			http.StatusOK,
		},
		{
			http.MethodPut,
			"/variable_write",
			struct {
				Board uint8
				Name  string
				Type  string
				Addr  uint32
				Data  float64
			}{
				Board: 1,
				Name:  "a",
				Type:  "double",
				Addr:  0x20123456,
				Data:  100,
			},
			http.StatusNoContent,
		},
		{
			http.MethodDelete,
			"/variable_write",
			struct {
				Board uint8
				Name  string
				Type  string
				Addr  uint32
			}{
				Board: 1,
				Name:  "a",
				Type:  "double",
				Addr:  0x20123456,
			},
			http.StatusNoContent,
		},
	}

	variableToWriteCtrl := makeVariableCtrl(variable.Write)

	ctrlerTest(variableToWriteCtrl, cases, t)
}

func TestVariableTypeCtrl(t *testing.T) {
	cases := casesT{
		{
			http.MethodGet,
			"/serials",
			nil,
			http.StatusOK,
		},
		{
			http.MethodPost,
			"/serials",
			nil,
			http.StatusMethodNotAllowed,
		},
	}

	ctrlerTest(variableTypeCtrl, cases, t)
}
