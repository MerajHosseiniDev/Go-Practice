package main

import (
	"fmt"
	"time"
)

func worker(done chan struct{}) {
	fmt.Println("Working...")
	time.Sleep(2*time.Second)
	fmt.Println("Done working")
	done <- struct{}{}
}

func main() {
	done := make(chan struct{})

	go worker(done)

	<- done

	fmt.Println("Main function is done!")
}