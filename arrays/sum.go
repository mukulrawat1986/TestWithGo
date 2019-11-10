package main

// Sum takes a slice as input and returns the sum of the slice
func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {}
