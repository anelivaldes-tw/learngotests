package arrays

// Simple a for loop
/*func Sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}*/

// Sum Using a range to iterate over an array
/*func Sum(numbers [5]int) (sum int) {
	sum = 0
	for _, n := range numbers {
		sum += n
	}
	return
}*/
// Sum Receiving a slice instead of an array,to handle collection of numbers of any size
func Sum(numbers []int) (sum int) {
	sum = 0
	for _, n := range numbers {
		sum += n
	}
	return
}

// SumAll which will take a varying number of slices, returning a new slice containing the totals for each slice passed in.
/*func SumAll(numbersToSum ...[]int) (sums []int) {
	lengthOfNumbers := len(numbersToSum)
	sums = make([]int, lengthOfNumbers)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return
}*/

// SumAll refactored version, do not worry about the slice capacity
func SumAll(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

// SumAllTails where it will calculate the totals of the "tails" of each slice. The tail of a collection is all items in the collection except the first one (the "head")
// Need a refactor to handle empty slice, because: panic: runtime error: slice bounds out of range [1:0]
/*func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		tail := numbers[1:] // Slices can be sliced! slice[low:high], eg: take from 1 to end
		sums = append(sums, Sum(tail))
	}
	return
}*/

func SumAllTails(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:] // Slices can be sliced! slice[low:high], eg: take from 1 to end
			sums = append(sums, Sum(tail))
		}
	}
	return
}

func SumAllFirstAndLastOne(numbersToSum ...[]int) (sums []int) {
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			first := numbers[0] // Slices can be sliced! slice[low:high], eg: take from 1 to end
			lastOne := numbers[(len(numbers) - 1)]
			values := []int{first, lastOne}
			sums = append(sums, Sum(values))
		}
	}
	return
}
