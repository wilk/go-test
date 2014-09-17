package main // it needs to be defined as main

import "fmt" // import fmt library... fmt stands for?

func hello_world() string {
	return "hello world"
}

func main() { // declaring main function. it executes when go runs the program
	fmt.Println("Hello world")

	fmt.Println(hello_world())
}
