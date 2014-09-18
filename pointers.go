package main

import "fmt"

// accept an int passed by value
func vint(i int) {
	i = 10
}

// accept an address, so an int passed by reference
func pint(i *int) {
	*i = 20
}

func main() {
	i := 1
	fmt.Println(i)

	// it doesn't change anything
	vint(i)
	fmt.Println(i)

	// it passes the reference so it will be changed
	pint(&i)
	fmt.Println(i)
}
