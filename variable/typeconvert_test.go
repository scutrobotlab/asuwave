package variable

import "testing"

func TestAnyToBytes(t *testing.T) {
	cases := []struct {
		in   interface{}
		want []byte
	}{
		{uint8(10), []byte{0x0a}},
		{uint16(65534), []byte{0xfe, 0xff}},
		{uint32(42949672), []byte{0x28, 0x5c, 0x8f, 0x02}},
		{uint64(18446744073709551615), []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}},
		{int8(-2), []byte{0xfe}},
		{int16(-134), []byte{0x7a, 0xff}},
		{int32(942949672), []byte{0x28, 0x45, 0x34, 0x38}},
		{int64(8446744073709551615), []byte{0xff, 0xff, 0x17, 0x76, 0xfb, 0xdc, 0x38, 0x75}},
		{float32(-8.25), []byte{0x00, 0x00, 0x04, 0xc1}},
		{float64(13278.125), []byte{0x0, 0x0, 0x0, 0x0, 0x10, 0xef, 0xc9, 0x40}},
	}
	for _, c := range cases {
		got := AnyToBytes(c.in)
		if string(got) != string(c.want) {
			t.Errorf("AnyToBytes(%#v) == %#v, want %#v", c.in, got, c.want)
		}
	}
}
