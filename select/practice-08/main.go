package main

import (
	"fmt"
	"time"
)

func task1(ch1 chan string) {
	time.Sleep(1 * time.Second)
	ch1 <- "goroutine 1 finished first.."
}

func task2(ch2 chan string) {
	time.Sleep(2 * time.Second)
	ch2 <- "goroutine 2 finished first.."
}

func task3(ch3 chan string) {
	time.Sleep(3 * time.Second)
	ch3 <- "goroutine 3 finished first.."
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)

	go task1(ch1)
	go task2(ch2)
	go task3(ch3)

	recieved := 0
	for recieved < 3 {
		select {
		case msg1 := <-ch1:
			fmt.Println("Recieved: ", msg1)
			recieved ++
		case msg2 := <-ch2:
			fmt.Println("Recieved: ", msg2)
			recieved ++
		case msg3 := <-ch3:
			fmt.Println("Recieved: ", msg3)
			recieved ++
		}
	}
}
