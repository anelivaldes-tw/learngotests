package arrays

import "fmt"

func ExampleSumAllFirstAndLastOne() {
	a := []int{2, 3, 45, 5}
	b := []int{8, 3, 45, 9}
	c := []int{4, 1}
	d := []int{9}
	e := []int{}
	r := SumAllFirstAndLastOne(a, b, c, d, e)
	fmt.Println(r)
	// Output: [7 17 5 18 0]
}
