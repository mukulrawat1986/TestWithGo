package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
)

// Hello function takes as argument a name and returns a greeting
// for that person
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return spanishHelloPrefix + name
	}
	return  englishHelloPrefix + name
}

func main() {
	fmt.Printf(Hello("", ""))
}
