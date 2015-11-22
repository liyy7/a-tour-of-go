package pill

import "testing"

func TestString(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{0, "Placebo"},
		{1, "Aspirin"},
		{2, "Ibuprofen"},
		{3, "Paracetamol"},
		{int(Acetaminophen), "Paracetamol"},
		{5, "Pill(5)"},
		{10, "Pill(10)"},
		{10000, "Pill(10000)"},
		{-10000, "Pill(-10000)"},
	}
	for _, c := range cases {
		var p Pill = Pill(c.in)
		got := p.String()
		if got != c.want {
			t.Errorf("Pill(%d).String() == %q, want %q", c.in, got, c.want)
		} else {
			t.Logf("Pill(%d).String() == %q", c.in, got)
		}
	}
}
