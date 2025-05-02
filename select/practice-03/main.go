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

	pinEntered := make(chan int)

	go func() {
		time.Sleep(2*time.Second)
		pinEntered <- 1234
	}()

	select {
	case pin := <- pinEntered:
		fmt.Printf("PIN entered: [%d]\nWelcome!", pin)
	case <- time.After(3*time.Second):
		fmt.Println("No PIN entered. Session ended.")
	}
}