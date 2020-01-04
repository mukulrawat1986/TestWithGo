package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
)

// Hello returns a string greeting
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}
