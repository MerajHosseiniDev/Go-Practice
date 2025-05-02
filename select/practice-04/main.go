package main

import (
	"fmt"
	"time"
)

func main() {

	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		chan1 <- "Message from channel 1."
	}()

	go func() {
		time.Sleep(2 * time.Second)
		chan2 <- "Message from channel 2."
	}()

	select {
	case msg1 := <-chan1:
		fmt.Println("Recieved: ", msg1)
	case msg2 := <-chan2:
		fmt.Println("Recieved: ", msg2)
	}
}
