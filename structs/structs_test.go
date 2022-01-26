package structs

import "testing"

func TestShapes(t *testing.T) {

	t.Run("Rectangle Perimeter", func(t *testing.T) {
		rectangle := Rectangle{Height: 10.0, Width: 10.0}
		got := Perimeter(rectangle)
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("Rectangle Area", func(t *testing.T) {
		rectangle := Rectangle{Height: 10.0, Width: 10.0}
		got := rectangle.Area()
		want := 100.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("Circle Area", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
