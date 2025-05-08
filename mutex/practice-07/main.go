package main

import (
	"fmt"
	"sync"
)


type ATM struct {
	Cards map[string]int
	mu 		sync.RWMutex
}

type TransactionType int

const (
	Deposit TransactionType = iota
	Withdraw
)

type Request struct {
	CardID	string
	Amount	int
	Type	TransactionType
}

func (atm *ATM) processRequest(req Request) {
	atm.mu.Lock()
	defer atm.mu.Unlock()

	balance, ok := atm.Cards[req.CardID]
	if !ok {
		fmt.Printf("Card %s not found\n", req.CardID)
		return
	}

	switch req.Type {
		case Deposit:
			balance += req.Amount
			fmt.Printf("Card %s: Deposit %d$, New Balance: %d$\n",req.CardID, req.Amount, balance)
		case Withdraw:
			if balance >= req.Amount {
				balance -= req.Amount
				atm.Cards[req.CardID] = balance
				fmt.Printf("Card %s: Withdraw %d$, New Balance: %d$\n", req.CardID, req.Amount, balance)
			} else {
				fmt.Printf("Card %s: Insufficient Funds to Withdraw %d$, Balance: %d$\n", req.CardID, req.Amount, balance)
			}
		default:
			fmt.Println("Invalid Transaction type!")
	}
}


func main() {
	atm := ATM{
		Cards: map[string]int{
			"card1": 1000,
			"card2": 800,
			"card3": 1200,
		},
	}

	requestList := []Request{
		{CardID: "card1", Amount: 200, Type: Deposit},
		{CardID: "card2", Amount: 150, Type: Withdraw},
		{CardID: "card3", Amount: 50, Type: Withdraw},
		{CardID: "card1", Amount: 300, Type: Deposit},
		{CardID: "card2", Amount: 450, Type: Withdraw},
		{CardID: "card3", Amount: 300, Type: Deposit},
		{CardID: "card1", Amount: 50, Type: Withdraw},
		{CardID: "card2", Amount: 100, Type: Deposit},
	}

	requests := make(chan Request)
	var wg sync.WaitGroup

	for i:=1; i<=3; i++ {
		wg.Add(1)
		go func (id int)  {
			defer wg.Done()
			for req := range requests {
				fmt.Printf("Goroutine %d Processing Request..\n", id)
				atm.processRequest(req)
			}
		}(i)
	}

	go func() {
		for _, req := range requestList {
			requests <- req
		}
		close(requests)
	}()

	wg.Wait()

	fmt.Println("All transactions completed.\nFinal Balances:")
	for CardID, balance := range atm.Cards {
		fmt.Printf("Card %s: %d$\n", CardID, balance)
	}

}