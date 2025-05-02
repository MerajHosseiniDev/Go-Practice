package main

import (
	"fmt"
	"time"
)

func worker(done chan struct{}) {
	for i:=1; i<=5; i++ {
		select {
		case <- done:
			fmt.Println("Sropped!")
			return
		default:
			fmt.Println("Counting: ", i)
			time.Sleep(1*time.Second)
		}
	}
	fmt.Println("Sounting Finished!")
}


func main() {
	done := make(chan struct{})

	go worker(done)

	time.Sleep(2500*time.Millisecond)
	done <- struct{}{}

	time.Sleep(1*time.Second)
}
