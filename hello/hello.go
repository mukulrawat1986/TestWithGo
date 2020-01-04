package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
)

// Hello returns a string greeting
func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello(""))
}
