package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name != "" {
		return englishHelloPrefix + name
	}
	return "Hello, World"
}

func main() {
	fmt.Println(Hello("Jack"))
}
