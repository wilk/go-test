package main

import "fmt"

// a function that returns an anonymous function that returns an int
func closure() func() int {
	i := 0

	// here's the recursive definition
	return func() int {
		i = i + 1
		return i
	}
}

func main() {
	c1 := closure()
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c1())
	fmt.Println(c1())

	c2 := closure()
	fmt.Println(c2())
	fmt.Println(c2())
}
