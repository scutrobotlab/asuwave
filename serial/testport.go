package serial

import (
	"errors"
	"fmt"

	"go.bug.st/serial"
)

type testPort struct{}

func newTestPort() serial.Port { return &testPort{} }

func (tp *testPort) SetMode(mode *serial.Mode) error { return nil }

func (tp *testPort) Read(p []byte) (n int, err error) {
	s := ""
	fmt.Scanln(&s)
	copy(p, []byte(s))
	fmt.Printf("Test port read: %v\n", p[:len(s)])
	return len(s), nil
}

func (tp *testPort) Write(p []byte) (n int, err error) {
	fmt.Printf("Test port write: %v\n", p)
	return len(p), nil
}

func (tp *testPort) ResetInputBuffer() error { return nil }

func (tp *testPort) ResetOutputBuffer() error { return nil }

func (tp *testPort) SetDTR(dtr bool) error { return errors.New("not supported") }

func (tp *testPort) SetRTS(rts bool) error { return errors.New("not supported") }

func (tp *testPort) GetModemStatusBits() (*serial.ModemStatusBits, error) {
	return nil, errors.New("not supported")
}

func (tp *testPort) Close() error { return nil }
