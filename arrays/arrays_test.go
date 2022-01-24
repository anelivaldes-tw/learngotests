package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	// These first test is redundant
	t.Run("Colletion of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		sum := Sum(numbers)
		expected := 15

		if sum != expected {
			t.Errorf("got %d want %d given, %v", sum, expected, numbers)
		}
	})
	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		sum := Sum(numbers)
		expected := 6

		if sum != expected {
			t.Errorf("got %d want %d given, %v", sum, expected, numbers)
		}
	})

	t.Run("SumAll", func(t *testing.T) {
		got := SumAll([]int{1, 2, 5}, []int{0, 9, 7}, []int{3, 4, 5})
		want := []int{8, 16, 12}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("SumAllTails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 5}, []int{0, 9, 7}, []int{3, 4, 5})
		want := []int{7, 16, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("SumAllTails with an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 5}, []int{}, []int{3, 4, 5})
		want := []int{7, 0, 9}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("SumAllFirstAndLastOne", func(t *testing.T) {
		got := SumAllFirstAndLastOne([]int{1, 2, 5}, []int{0, 9, 7}, []int{3, 4})
		want := []int{6, 7, 7}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

}
