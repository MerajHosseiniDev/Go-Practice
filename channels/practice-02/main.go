package main

import (
	"fmt"
	"sync"
)

func doubleAndSend(n int, ch chan int) {
	result := n * 2
	ch <- result
}


func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 5)

	for i:=1; i<=5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			doubleAndSend(n, ch)
		}(i)
	}

	wg.Wait()

	close(ch)

	var sum int

	for result := range ch {
		sum += result
	}

	fmt.Println("Total sum: ", sum)
}