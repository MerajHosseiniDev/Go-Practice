package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "mwssage from channel."
	}()

	select {
	case msg := <-ch:
		fmt.Println("Recieved: ", msg)
	default:
		fmt.Println("No data yet.")
	}
}
