package variable_test

import (
	"fmt"
	"testing"

	"github.com/scutrobotlab/asuwave/internal/variable"
)

func TestFilt(t *testing.T) {
	cases := []struct {
		in        []byte
		listV     map[uint32]variable.T
		wantChart []variable.ChartT
		wantAdd   []variable.CmdT
		wantDel   []variable.CmdT
	}{
		{
			in: []byte{
				0x01, 0x02, 0x04,
				0x56, 0x34, 0x12, 0x80,
				0x00, 0x00, 0x04, 0xc1, 0x00, 0x00, 0x00, 0x00,
				0x01, 0x00, 0x00, 0x00,
				0x0a,
			},
			listV: map[uint32]variable.T{
				0x80123456: {
					Board:      1,
					Name:       "a",
					Type:       "float",
					Addr:       0x80123456,
					SignalGain: 1,
				},
				0x80654321: {
					Board:      1,
					Name:       "b",
					Type:       "int",
					Addr:       0x80654321,
					SignalGain: 1,
				},
			},
			wantChart: []variable.ChartT{
				{
					Board: 1,
					Name:  "a",
					Data:  -8.25,
					Tick:  1,
				},
			},
			wantAdd: []variable.CmdT{
				{
					Board:   1,
					TypeLen: 4,
					Addr:    0x80654321,
				},
			},
			wantDel: []variable.CmdT{},
		},

		{
			in: []byte{
				0x01, 0x02, 0x04,
				0x56, 0x34, 0x12, 0x80,
				0x00, 0x00, 0x04, 0xc1, 0x00, 0x00, 0x00, 0x00,
				0x01, 0x00, 0x00, 0x00,
				0x0a,
				0x01, 0x02, 0x04,
				0x21, 0x43, 0x65, 0x80,
				0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
				0x01, 0x00, 0x00, 0x00,
				0x0a,
			},
			listV: map[uint32]variable.T{
				0x80123456: {
					Board:      1,
					Name:       "a",
					Type:       "float",
					Addr:       0x80123456,
					SignalGain: 1,
				},
			},
			wantChart: []variable.ChartT{
				{
					Board: 1,
					Name:  "a",
					Data:  -8.25,
					Tick:  1,
				},
			},
			wantAdd: []variable.CmdT{},
			wantDel: []variable.CmdT{
				{
					Board:   1,
					TypeLen: 4,
					Addr:    0x80654321,
				},
			},
		},
	}
	for _, c := range cases {
		variable.SetAll(variable.Read, c.listV)
		gotChart, gotAdd, gotDel := variable.Filt(c.in)
		if ok, msg := assertEQ(gotChart, c.wantChart); !ok {
			t.Errorf("Chart list " + msg)
		}
		if ok, msg := assertEQ(gotAdd, c.wantAdd); !ok {
			t.Errorf("Add list " + msg)
		}
		if ok, msg := assertEQ(gotDel, c.wantDel); !ok {
			t.Errorf("Del list " + msg)
		}
	}
}

func assertEQ[T variable.ChartT | variable.CmdT](a []T, b []T) (bool, string) {
	if len(a) != len(b) {
		return false, fmt.Sprint("Different length", len(a), len(b))
	}
	for i := range a {
		if a[i] != b[i] {
			return false, fmt.Sprintf("Different at %d: %v != %v", i, a[i], b[i])
		}
	}
	return true, ""
}
