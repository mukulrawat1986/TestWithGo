package main

import "fmt"


// Hello function returns a string
func Hello(name string) string {
	return "Hello, " + name
}

func main() {
	fmt.Printf(Hello(""))
}
