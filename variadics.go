package main

import "fmt"

// three dots to explode the input into an array
func total(nums ...int) int {
	result := 0

	for _, num := range nums {
		result = result + num
	}

	return result
}

func main() {
	t1 := total(1,2,3)
	fmt.Println(t1)

	// event array and slices can be passed as variadic
	// just use three dots notation at the end
	t2 := total([]int{1,2,3,4}...)
	fmt.Println(t2)
}
