package main

import (
	"fmt"
	"sync"
	"time"
)

type BankAccount struct {
	Balance	int
	mu		sync.RWMutex
}

func (b *BankAccount) deposit(wg *sync.WaitGroup, depositRequest chan int) {
	defer wg.Done()

	for amount := range depositRequest{
		time.Sleep(500*time.Millisecond)

		b.mu.Lock()
		b.Balance += amount
		b.mu.Unlock()

		fmt.Printf("Deposit of %d$ successful!\n", amount)
	}
}

func (b *BankAccount) CheckBalance() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	
	return b.Balance
}


func main() {
	var wg sync.WaitGroup
	BankAccount := BankAccount{Balance: 200}
	depositRequest := make(chan int)
	amounts := []int{100,1500,50,40}

	wg.Add(1)
	go BankAccount.deposit(&wg, depositRequest)

	for _, amount := range amounts {
		depositRequest <- amount
	}

	close(depositRequest)
	wg.Wait()

	fmt.Printf("Current Balance: %d$\n", BankAccount.CheckBalance())

}