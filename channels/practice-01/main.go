package main

import (
	"fmt"
)

func task1(ch chan int) {

	ch<- 21
	ch<- 23
	ch<- 45
	
	close(ch)
}


func main() {
	ch := make(chan int)

	go task1(ch)

	for {
		value, ok := <- ch
		if !ok {
			break
		}
		fmt.Println(value)
		fmt.Println("i finished first channel practice")
	}
}