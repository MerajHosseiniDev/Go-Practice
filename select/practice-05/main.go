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

	for {
		select {
		case msg := <-ch:
			fmt.Println("Recieved: ", msg)
			return
		default:
			fmt.Println("No data yet.")
			time.Sleep(5008*time.Millisecond)
		}
	}
}
