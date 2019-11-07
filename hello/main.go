package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// Hello function takes as argument a name and returns a greeting
// for that person
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return spanishHelloPrefix + name
	} else if language == "French" {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Printf(Hello("", ""))
}
