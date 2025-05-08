package main

import (
	"fmt"
	"sync"
)

type ATM struct {
	balance	int
	mu 		sync.Mutex
}

func (a *ATM) deposit(depositRequest chan int, wg *sync.WaitGroup){
	defer wg.Done()
	for amount := range depositRequest {
		a.mu.Lock()
		a.balance += amount
		fmt.Printf("Deposited %d$\nNew balance: %d$\n", amount, a.balance)
		a.mu.Unlock()
	}
}

func main() {
	atm := &ATM{balance: 500}
	var wg sync.WaitGroup
	depositRequest := make(chan int)
	amounts := []int{100, 200, 5, 150}

	wg.Add(1)
	go atm.deposit(depositRequest, &wg)

	for _, amount := range amounts {
		depositRequest <- amount
	}

	close(depositRequest)
	wg.Wait()

	fmt.Printf("Final balance: %d$\n", atm.balance)

}