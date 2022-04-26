package server

import (
	"net/http"
	"testing"

	"github.com/scutrobotlab/asuwave/variable"
)

func TestVariableToReadCtrl(t *testing.T) {
	// serial.Open("Test port")
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

	variableToReadCtrl := makeVariableCtrl(&variable.ToRead, true)

	ctrlerTest(variableToReadCtrl, "variableToReadCtrl", cases, t)
}

func TestVariableToModiCtrl(t *testing.T) {
	cases := casesT{
		{
			http.MethodGet,
			"/variable_modi",
			nil,
			http.StatusOK,
		},
		{
			http.MethodPost,
			"/variable_modi",
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
			"/variable_modi",
			nil,
			http.StatusOK,
		},
		{
			http.MethodPut,
			"/variable_modi",
			struct {
				Board uint8
				Name  string
				Type  string
				Addr  uint32
				Data  float64
			}{
				Board: 1,
				Name:  "a",
				Type:  "int",
				Addr:  0x20123456,
				Data:  100,
			},
			http.StatusInternalServerError,
		},
		{
			http.MethodDelete,
			"/variable_modi",
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

	variableToModiCtrl := makeVariableCtrl(&variable.ToModi, false)

	ctrlerTest(variableToModiCtrl, "variableToModiCtrl", cases, t)
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

	ctrlerTest(variableTypeCtrl, "variableTypeCtrl", cases, t)
}
