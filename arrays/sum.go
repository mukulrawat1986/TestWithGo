package main

// Sum takes an array of 5 integers and returns their sum
func Sum(numbers [5]int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {}
