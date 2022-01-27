package structs

import "testing"

func TestShapes(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

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
		want := 100.0
		checkArea(t, rectangle, want)
	})
	t.Run("Circle Area", func(t *testing.T) {
		circle := Circle{10}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})
}
