package main

import "fmt"

// three types of for
func main() {
	// one single condition
	i := 0
	for i < 4 {
		fmt.Println(i)
		i++
	}

	// classical initial/condition/after loop
	for i := 4; i < 10; i++ {
		fmt.Println(i)
	}

	// infinite loop
	for {
		fmt.Println("loop")
		break
	}
}
