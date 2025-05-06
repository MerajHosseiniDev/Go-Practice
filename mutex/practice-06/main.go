package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	Balance  int
	mu       sync.RWMutex
}

type TracsactionType int

const (
	Deposit TracsactionType = iota
	Withdraw 
)

type Request struct {
	Amount int
	Type   TracsactionType
}

func (b *BankAccount) processRequest(req Request) {
	b.mu.Lock()
	defer b.mu.Unlock()

	switch req.Type {
		case Deposit:
			b.Balance += req.Amount
			fmt.Printf("Deposited %d$, New Balance: %d$\n", req.Amount, b.Balance)
		case Withdraw:
			if b.Balance>= req.Amount {
				b.Balance -= req.Amount
				fmt.Printf("Withdrew %d$, New Balance: %d$\n", req.Amount, b.Balance)
			} else {
				fmt.Printf("Insufficient funds to withdraw %d$\n", req.Amount)
			}
		default:
			fmt.Println("Invalid action")
	}
}


func main() {
	bankAccount := BankAccount{Balance: 1000}
	requestsList := []Request{
		{Amount: 100, Type: Deposit},
		{Amount: 50, Type: Withdraw},
		{Amount: 200, Type: Withdraw},
		{Amount: 300, Type: Deposit},
		{Amount: 150, Type: Withdraw},
	}
	requests := make(chan Request)
	var wg sync.WaitGroup

	for i:=1; i<=3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for req := range requests {
				fmt.Printf("Goroutine %d processing request..\n", id)
				bankAccount.processRequest(req)
			}
		}(i)
	}

	go func() {
		for _, req := range requestsList {
			requests <- req
		}
		close(requests)
	}()

	wg.Wait()

	fmt.Printf("Final balance: %d$\n", bankAccount.Balance)





}



