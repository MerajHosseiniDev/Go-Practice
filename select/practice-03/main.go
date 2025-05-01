package main

import (
	"fmt"
	"time"
)

func main() {
	cardInserted := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		cardInserted <- "Card inserted!"
	}()

	select {
	case msg := <-cardInserted:
		fmt.Println("Card detected: ", msg)
	case <-time.After(5 * time.Second):
		fmt.Println("No card inserted. session timed out.")
	}
}
