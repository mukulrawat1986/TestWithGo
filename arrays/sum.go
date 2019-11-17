package main

// Sum takes a slice as input and returns the sum of the slice
func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll takes varying number of slices and returns a new slice
// containing the total for each slice passed in
func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func main() {}
