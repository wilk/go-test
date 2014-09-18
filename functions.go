package main

import "fmt"

// func <func_name>(<arg1> <type>, <arg2> <type>, ...) <return_type>
func sum(a int, b int) int {
	return a + b
}

// multiple return
// func <func_name>(<arg1> <type>, <arg2> <type>, ...) (<return_type_1> <return_type_2> <return_type_3> ...)
func values() (int, int, int) {
	return 2, 4, 6
}

func main() {
	s := sum(10, 20)
	fmt.Println(s)

	// it's possible to use the blank identifier to avoid some values
	_, _, a := values()
	fmt.Println(a)
}
