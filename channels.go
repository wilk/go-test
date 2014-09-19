package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("starting")

	var message string

	// channels are the same type of what they transport
	channel := make(chan string)

	fmt.Println("send something into the channel asynchronousely")
	go func () {
		time.Sleep(time.Second * 2)
		// sending something to the channel with '<-' syntax
		channel <- "hello world"
	}()

	fmt.Println("receiving something from the channel")
	// receiving something is a blocking operation
	// so the main thread will be blocked until something arrives down the channel
	message = <- channel
	fmt.Println(message)

	fmt.Println("finishing")
}
