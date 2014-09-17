package main

import "fmt"

func main() {
	// classical switch
	i := 2
	switch i {
	case 1:
		fmt.Println("i is first")
	case 2:
		fmt.Println("i is second")
	case 3:
		fmt.Println("i is third")
	}

	// default case and more than one option on the same case
	i = 10
	switch i {
	case 1, 2:
		fmt.Println("first and second")
	default:
		fmt.Println("greater")
	}

	// switch can be used like an if-else statement if no condition is specified
	switch {
	case i < 10:
		fmt.Println("lower then 10")
	case i > 10:
		fmt.Println("greater then 10")
	default:
		fmt.Println("actually is 10")
	}
}
