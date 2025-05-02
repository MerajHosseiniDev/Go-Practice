package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.Tick(2*time.Second)
	timeout := time.After(10*time.Second)

	for {
		select {
		case t := <- ticker:
			fmt.Println("Tick received at ", t)
		case <- timeout:
			fmt.Println("Timeout reached, pragram will stop!")
			return
		}
	}
}