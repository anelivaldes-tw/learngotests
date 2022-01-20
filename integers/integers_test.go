package integers

import (
	"testing"
)

func TestIntegers(t *testing.T) {
	t.Run("add two integers", func(t *testing.T) {
		sum := Add(2, 3)
		expected := 5

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})

	t.Run("multiply two integers", func(t *testing.T) {
		sum := Multiply(2, 3)
		expected := 6

		if sum != expected {
			t.Errorf("expected '%d' but got '%d'", expected, sum)
		}
	})
}
