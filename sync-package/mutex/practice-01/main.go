package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	balance = 0
	mu sync.Mutex
)

func deposit(wg *sync.WaitGroup, ch <- chan int) {
	defer wg.Done()
	
	for amount := range ch {
		mu.Lock()

		balance += amount
		time.Sleep(1*time.Second)
		fmt.Printf("Deposit of %d$ successful!\nCurrent balance: %d$\n", amount, balance) 
	
		mu.Unlock()
		time.Sleep(500*time.Millisecond)
	}
}

func main() {
	var wg sync.WaitGroup
	amounts := []int{100, 200, 5, 150}
	depositRequest := make(chan int)

	wg.Add(1)
	go deposit(&wg, depositRequest)

	for _, amount := range amounts {
		depositRequest <- amount
	}

	close(depositRequest)
	wg.Wait()
	
	fmt.Println("Program completely weorked!")
}