package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
)

// Hello function takes as argument a name and returns a greeting
// for that person
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return  englishHelloPrefix + name
}

func main() {
	fmt.Printf(Hello(""))
}
