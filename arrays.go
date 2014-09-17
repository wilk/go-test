package main

import "fmt"

func main() {
	// basic array definition
	// var <var_name> [len]<var_type>
	var arr [5]int
	fmt.Println(arr)

	// get and set like normal arrays
	arr[2] = 10
	fmt.Println(arr)
	fmt.Println(arr[4])

	// inline definition
	// <var_name> := [len]<var_type>{el1, el2, ...}
	arr2 := [5]int{0,1,2,3,4}
	fmt.Println(arr2)

	// multidimensional arrays
	// var <var_name> [len][len]...<var_type>
	var multiArr [5][5]int
	fmt.Println(multiArr)

	// lenght
	fmt.Println(len(multiArr))
}
