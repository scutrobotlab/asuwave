package datautil

import (
	"encoding/json"
	"testing"

	"github.com/scutrobotlab/asuwave/variable"
)

func TestMakeCmd(t *testing.T) {
	cases := []struct {
		act  uint8
		v    variable.T
		want []byte
	}{
		{
			ActModeWrite,
			variable.T{
				Board: 1,
				Name:  "a",
				Type:  "float",
				Addr:  0x80123456,
				Data:  -8.25,
				Tick:  0,
			},
			[]byte{0x01, 0x07, 0x04, 0x56, 0x34, 0x12, 0x80, 0x00, 0x00, 0x04, 0xc1, 0x00, 0x00, 0x00, 0x00, 0x0a},
		},
	}

	for _, c := range cases {
		got := MakeCmd(c.act, &c.v)
		if string(got) != string(c.want) {
			t.Errorf("makeCmd(%#v,%#v) == %#v, want %#v", c.act, c.v, got, c.want)
		}
	}
}

func TestFindValidPart(t *testing.T) {
	cases := []struct {
		data     []byte
		startIdx int
		endIdx   int
	}{
		{
			data: []byte{
				0x56, 0x34, 0x12, 0x80,
				0x00, 0x00, 0x04, 0xc1, 0x01, 0x02, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x01,
			},
			startIdx: 0,
			endIdx:   0,
		},

		{
			data: []byte{
				0x01, 0x02, 0x04,
				0x56, 0x34, 0x12, 0x80,
				0x00, 0x00, 0x04, 0xc1, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x01,
				0x0a,
			},
			startIdx: 0,
			endIdx:   20,
		},

		{
			data: []byte{
				0x01, 0x02, 0x04,
				0x56, 0x34, 0x12, 0x80,
				0x00, 0x00, 0x04, 0xc1, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x01,
				0x0a,
				0x01, 0x02, 0x04,
				0x56, 0x34, 0x12, 0x80,
				0x00, 0x00, 0x04, 0xc1, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x01,
				0x0a,
			},
			startIdx: 0,
			endIdx:   40,
		},

		{
			data: []byte{
				0x56, 0x34, 0x12, 0x80,
				0x00, 0x00, 0x04, 0xc1, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x01,
				0x0a,

				0x01, 0x02, 0x04,
				0x56, 0x34, 0x12, 0x80,
				0x00, 0x00, 0x04, 0xc1, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x01,
				0x0a,

				0x01, 0x02, 0x04,
			},
			startIdx: 17,
			endIdx:   37,
		},

		{
			data: []byte{
				0x00, 0x00, 0x01, 0x02,
				0x0a,

				0x01, 0x02, 0x04,
				0x56, 0x34, 0x12, 0x80,
				0x00, 0x00, 0x04, 0xc1, 0x00, 0x00, 0x00, 0x00,
				0x00, 0x00, 0x00, 0x01,
				0x0a,

				0x00, 0x0a, 0x01, 0x02,
			},
			startIdx: 5,
			endIdx:   25,
		},
	}
	for _, c := range cases {
		s, n := FindValidPart(c.data)
		if s != c.startIdx || n != c.endIdx {
			t.Errorf("VerifyBuff(%#v) == %#v,%#v want %#v,%#v", c.data, s, n, c.startIdx, c.endIdx)
		}
	}
}

func TestMakeChartPack(t *testing.T) {
	cases := []struct {
		in    []byte
		listV variable.ListT
		want  variable.ListChartT
	}{
		{
			in: []byte{
				0x01, 0x02, 0x04,
				0x56, 0x34, 0x12, 0x80,
				0x00, 0x00, 0x04, 0xc1, 0x00, 0x00, 0x00, 0x00,
				0x01, 0x00, 0x00, 0x00,
				0x0a,
			},
			listV: variable.ListT{
				Variables: []variable.T{
					{
						Board: 1,
						Name:  "a",
						Type:  "float",
						Addr:  0x80123456,
					},
				},
			},
			want: variable.ListChartT{
				Variables: []variable.ToChartT{
					{
						Board: 1,
						Name:  "a",
						Data:  -8.25,
						Tick:  1,
					},
				},
			},
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
			listV: variable.ListT{
				Variables: []variable.T{
					{
						Board: 1,
						Name:  "a",
						Type:  "float",
						Addr:  0x80123456,
					},
					{
						Board: 1,
						Name:  "b",
						Type:  "int",
						Addr:  0x80654321,
					},
				},
			},
			want: variable.ListChartT{
				Variables: []variable.ToChartT{
					{
						Board: 1,
						Name:  "a",
						Data:  -8.25,
						Tick:  1,
					},
					{
						Board: 1,
						Name:  "b",
						Data:  1,
						Tick:  1,
					},
				},
			},
		},
	}
	for _, c := range cases {
		var got variable.ListChartT
		MakeChartPack(&got, &c.listV, c.in)
		b1, _ := json.Marshal(got)
		b2, _ := json.Marshal(c.want)
		if string(b1) != string(b2) {
			t.Errorf("MakeChartPack(%#v,%#v) want %#v", got, c.in, c.want)
		}
	}
}

func TestGetProjectVariables(t *testing.T) {
	txt := `
address     size       variable name                            type
0x200055bc  0x4        Last_YawToSend                           float

address     size       variable name                            type
0x200055b8  0x4        YawToSend                                float

address     size       variable name                            type
0x200055d0  0x2        ax                                       short int
0x200055ea  0x2        ay                                       short int
0x200055c6  0x2        az                                       short int
`
	var x variable.ListProjectT
	y := variable.ListProjectT{
		Variables: []variable.ToProjectT{
			{
				Addr: "0x200055bc",
				Size: "0x4",
				Name: "Last_YawToSend",
				Type: "float",
			},
			{
				Addr: "0x200055b8",
				Size: "0x4",
				Name: "YawToSend",
				Type: "float",
			},
			{
				Addr: "0x200055d0",
				Size: "0x2",
				Name: "ax",
				Type: "short int",
			},
			{
				Addr: "0x200055ea",
				Size: "0x2",
				Name: "ay",
				Type: "short int",
			},
			{
				Addr: "0x200055c6",
				Size: "0x2",
				Name: "az",
				Type: "short int",
			},
		},
	}

	GetProjectVariables(&x, []byte(txt))
	b1, _ := json.Marshal(x)
	b2, _ := json.Marshal(y)
	if string(b1) != string(b2) {
		t.Errorf("TestGetProjectVariables(%#v,txt) want %#v", x, y)
	}
}
