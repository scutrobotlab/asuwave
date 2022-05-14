package option_test

import (
	"testing"

	"github.com/scutrobotlab/asuwave/internal/option"
)

func FuzzOption(f *testing.F) {
	f.Fuzz(func(t *testing.T, logLevel int, saveFilePath bool, saveVarList bool, updateByProj bool) {

		option.SetLogLevel(logLevel)
		option.SetSaveFilePath(saveFilePath)
		option.SetSaveVarList(saveVarList)
		option.SetUpdateByProj(updateByProj)

		got := option.Get()

		assertEQ(t, got.LogLevel, logLevel)
		assertEQ(t, got.SaveFilePath, saveFilePath)
		assertEQ(t, got.SaveVarList, saveVarList)
		assertEQ(t, got.UpdateByProj, updateByProj)
	})
}

func assertEQ[V int | bool](t *testing.T, a V, b V) {
	if a != b {
		t.Errorf("%v != %v", a, b)
	} else {
		t.Logf("%v == %v", a, b)
	}
}
