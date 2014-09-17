package main

import "fmt"

func main() {
	// basic if condition
	// always open a block {}
	if 10 < 11 {
		fmt.Println("true")
	}

	// if-else condition
	if 10 > 9 {
		fmt.Println("true")
	// else needs to be on the same line
	} else {
		fmt.Println("false")
	}

	// if-else-if condition
	if 10 == 10 {
		fmt.Println("true")
	} else if 10 != 10 {
		fmt.Println("false")
	} else {
		fmt.Println("whatever")
	}

	// it's possible to declare different vars in the same if-else block
	// they will appear with the same value in every branches
	if num := 10; num < 11 {
		fmt.Println(num, "is 10")
	} else if num > 9 {
		fmt.Println(num, "is still 10")
	} else {
		fmt.Println(num, "either here")
	}
}
