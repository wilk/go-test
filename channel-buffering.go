package main

import (
	"fmt"
	"time"
)

func main() {
	// a buffered channel that buffers two incoming messages
	messages := make(chan string, 2)

	go func () {
		time.Sleep(time.Second * 2)
		messages <- "hello"
		time.Sleep(time.Second * 2)
		messages <- "world"
		// and the third one??
		time.Sleep(time.Second * 2)
		messages <- "!!!!"
	}()

	fmt.Println(<- messages)
	fmt.Println(<- messages)
	// and the third one??
	fmt.Println(<- messages)
}
