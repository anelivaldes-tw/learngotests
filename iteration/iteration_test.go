package iteration

import "testing"

func TestIteration(t *testing.T) {
	times := 5
	repeated := Repeat("a", times)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
