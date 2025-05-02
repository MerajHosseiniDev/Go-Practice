package main

import (
	"fmt"
	"time"
)

func main() {

	messageChannel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		messageChannel <- "Message from channel"
	}()

	select {
	case msg := <-messageChannel:
		fmt.Println("Recieved: ", msg)
	case <-time.After(3 * time.Second):
		fmt.Println("No message recieved! session ended.")
		return
	}
}
