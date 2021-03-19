package option

import "testing"

func TestCheckCanSave(t *testing.T) {
	cases := []struct {
		save int
		in   int
		want bool
	}{
		{1, SaveVariableProj, true},
		{2, SaveVariableModi, true},
		{4, SaveVariableRead, true},
		{1, SaveVariableModi, false},
		{3, SaveVariableRead, false},
		{6, SaveVariableProj, false},
		{6, SaveVariableModi, true},
		{7, SaveVariableRead, true},
	}
	for _, c := range cases {
		Config.Save = c.save
		got := CheckCanSave(c.in)
		if got != c.want {
			t.Errorf("checkCanSave(%#v) == %#v, want %#v", c.in, got, c.want)
		}
	}
}
