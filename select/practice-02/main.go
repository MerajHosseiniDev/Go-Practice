package main

import (
	"fmt"
	"time"
)

func main() {
	fast := make(chan string)
	slow := make(chan string)

	go func() {
		time.Sleep(1*time.Second)
		message1 := "Message from fast channel!"
		fast <- message1
	}()

	go func() {
		time.Sleep(3*time.Second)
		message2 := "Message from slow channel!"
		slow <- message2
	}()

	
	for i:=0; i<2; i++ {
		select {
		case msg1 := <- slow:
			fmt.Println("Recieved message from slow channel: ", msg1)
		case msg2 := <- fast:
			fmt.Println("Recieved message from fast channel: ", msg2)
		}
	}
}