package main

import (
	"fmt"
	"time"
)

// a sample log function
func log(from string) {
	for i := 0; i < 4; i++ {
		fmt.Println(from, i)
	}
}

func main() {
	// it's executed directly on the main thread
	log("first")

	// it's executed in a parallel thread
	go log("second")

	// so this anonymous function
	go func (msg string) {
		time.Sleep(time.Second * 2)
		fmt.Println(msg)
	}("third")

	var input string
	fmt.Scanln(&input)
}
