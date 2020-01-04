package main

import "fmt"

// Hello returns a string greeting
func Hello(name string) string {
	return "Hello, World"
}

func main() {
	fmt.Println(Hello(""))
}
