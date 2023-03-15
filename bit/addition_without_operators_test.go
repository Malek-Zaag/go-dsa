package bit

import "testing"

func TestAdd(t *testing.T) {
	var tests = []struct {
		a, b int
	}{
		{0, 1},
		{20, 2},
		{20, 4},
		{20, 3},
	}

	for i, test := range tests {
		got := Add(test.a, test.b)
		want := test.a + test.b
		if got != want {
			t.Fatalf("Failed test case #%d. Want %#v got %#v", i, want, got)
		}
	}
}
