package variable

import (
	"fmt"
	"testing"
)

func TestFilt(t *testing.T) {
	cases := []struct {
		in        []byte
		listV     map[uint32]T
		wantChart []ChartT
		wantAdd   []CmdT
		wantDel   []CmdT
	}{
		{
			in: []byte{
				0x01, 0x02, 0x04,
				0x56, 0x34, 0x12, 0x80,
				0x00, 0x00, 0x04, 0xc1, 0x00, 0x00, 0x00, 0x00,
				0x01, 0x00, 0x00, 0x00,
				0x0a,
			},
			listV: map[uint32]T{
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
			wantChart: []ChartT{
				{
					Board: 1,
					Name:  "a",
					Data:  -8.25,
					Tick:  1,
				},
			},
			wantAdd: []CmdT{
				{
					Board:  1,
					Length: 4,
					Addr:   0x80654321,
				},
			},
			wantDel: []CmdT{},
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
			listV: map[uint32]T{
				0x80123456: {
					Board:      1,
					Name:       "a",
					Type:       "float",
					Addr:       0x80123456,
					SignalGain: 1,
				},
			},
			wantChart: []ChartT{
				{
					Board: 1,
					Name:  "a",
					Data:  -8.25,
					Tick:  1,
				},
			},
			wantAdd: []CmdT{},
			wantDel: []CmdT{
				{
					Board:  1,
					Length: 4,
					Addr:   0x80654321,
				},
			},
		},
	}
	for _, c := range cases {
		SetAll(RD, c.listV)
		gotChart, gotAdd, gotDel := Filt(c.in)
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

func assertEQ[T ChartT | CmdT](a []T, b []T) (bool, string) {
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
