package server

import (
	"net/http"
	"testing"

	"github.com/scutrobotlab/asuwave/internal/serial"
)

func TestSerialCtrl(t *testing.T) {
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
		{
			http.MethodPut,
			"/serials",
			nil,
			http.StatusMethodNotAllowed,
		},
		{
			http.MethodDelete,
			"/serials",
			nil,
			http.StatusMethodNotAllowed,
		},
	}
	ctrlerTest(serialCtrl, cases, t)
}

func TestSerialCurCtrl(t *testing.T) {
	go serial.GrReceive()
	go serial.GrTransmit()
	go serial.GrRxPrase()
	cases := casesT{
		{
			http.MethodGet,
			"/serial_cur",
			nil,
			http.StatusOK,
		},
		{
			http.MethodPost,
			"/serial_cur",
			struct{ Serial string }{Serial: "Test port"},
			http.StatusOK,
		},
		{
			http.MethodGet,
			"/serial_cur",
			nil,
			http.StatusOK,
		},
		{
			http.MethodPut,
			"/serial_cur",
			struct{ Serial string }{Serial: "Test port"},
			http.StatusMethodNotAllowed,
		},
		{
			http.MethodDelete,
			"/serial_cur",
			nil,
			http.StatusNoContent,
		},
	}
	ctrlerTest(serialCurCtrl, cases, t)
}
