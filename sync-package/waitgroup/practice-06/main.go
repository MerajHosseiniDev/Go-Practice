package main

import (
	"fmt"
	"sync"
)


func writer(wg *sync.WaitGroup, ch chan<- int) {
	defer wg.Done()
	ch <- 10
}

func reader(wg *sync.WaitGroup, ch <- chan int) {
	defer wg.Done()
	number := <- ch
	fmt.Println("Reader received: ", number)
}

func writeAndRead(wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	ch <- 11
	number := <- ch
	fmt.Println("Write and read recieved: ", number)
}


func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	ch2 := make(chan int, 1)

	wg.Add(3)

	go writer(&wg, ch)
	go reader(&wg, ch)
	go writeAndRead(&wg, ch2)

	wg.Wait()

}