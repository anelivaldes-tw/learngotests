package iteration

import (
	"fmt"
	"testing"
)

func ExampleRepeat() {
	repeated := Repeat("w", 7)
	fmt.Println(repeated)
	// Output: wwwwwww
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("r", 100)
	}
}
