package slip_test

import (
	"testing"

	"github.com/scutrobotlab/asuwave/pkg/slip"
)

func FuzzSlip(f *testing.F) {
	for s := 0; s < 40; s++ {
		seed := []byte{}
		for i := 0; i < s; i++ {
			if i%2 == 0 {
				seed = append(seed, slip.END)
			}
			if i%3 == 0 {
				seed = append(seed, slip.ESC)
			}
			if i%4 == 0 {
				seed = append(seed, slip.ESC_END)
			}
			if i%5 == 0 {
				seed = append(seed, slip.ESC_ESC)
			}
		}
		f.Add(seed)
	}
	f.Fuzz(func(t *testing.T, PDU []byte) {
		SDU := slip.Pack(PDU)

		if SDU[0] != slip.END {
			t.Fatal("SDU[0] not END, but: ", SDU[0])
		}
		if SDU[len(SDU)-1] != slip.END {
			t.Fatal("SDU[-1] not END, but: ", SDU[len(SDU)-1])
		}
		for i, s := range SDU {
			if i != 0 && i != len(SDU)-1 && s == slip.END {
				t.Fatalf("SDU[%d] got END", i)
			}
		}

		nPDU, err := slip.Unpack(SDU)
		if err != nil {
			t.Fatal(err.Error())
		}

		t.Logf("PDU %v -> SDU %v -> nPDU %v", PDU, SDU, nPDU)

		if len(PDU) != len(nPDU) {
			t.Fatalf("PDU %v (len: %d) != nPDU %v (len: %d)", PDU, len(PDU), nPDU, len(nPDU))
		}
		for i := range PDU {
			if PDU[i] != nPDU[i] {
				t.Fatalf("PDU[%d](%d) != nPDU[%d](%d)", i, PDU[i], i, nPDU[i])
			}
		}
		t.Log("Success: ", PDU)
	})
}
