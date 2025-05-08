package main

import (
	"fmt"
	"sync"
	"time"
)

type WalletSystem struct {
	Users	map[string]int
	mu		sync.RWMutex	  
}

type TransactionType int

const (
	Deposit TransactionType = iota
	Withdraw
)

type Request struct {
	UserID	string
	Amount	int
	Type 	TransactionType
}

func (t TransactionType) String() string {
	switch t {
		case Deposit:
			return "Deposit"
		case Withdraw:
			return "Withdraw"
		default:
			return "Unknown"
	}
}

func (w *WalletSystem) ProcessTransaction(req Request) {
	w.mu.Lock()
	defer w.mu.Unlock()

	balance, ok := w.Users[req.UserID]
	if !ok {
		fmt.Printf("User %s not found!\n", req.UserID)
		return
	}

	switch req.Type {
		case Deposit:
			balance += req.Amount
			w.Users[req.UserID] = balance
			fmt.Printf("Card %s: Deposit %d$, New Balance: %d$\n", req.UserID, req.Amount, balance)
		case Withdraw:
			if balance >= req.Amount {
				balance -= req.Amount
				w.Users[req.UserID] = balance
				fmt.Printf("Card %s: Withdrew %d$, New Balance: %d$\n", req.UserID, req.Amount, balance)
			} else {
				fmt.Printf("Card %s: Insufficient funds to withdraw %d$, Balance: %d$\n", req.UserID, req.Amount, balance)
			}
		default:
			fmt.Println("Invalid Transication type!")	
	}
}

func main() {
	walletSystem := WalletSystem{
		Users: map[string]int{
			"user1": 800,
			"user2": 1000,
			"user3": 1200,
		},
	}

	requestsList := []Request {
		{UserID: "user1", Amount: 200, Type: Deposit},
		{UserID: "user2", Amount: 300, Type: Withdraw},
		{UserID: "user3", Amount: 100, Type: Deposit},
		{UserID: "user1", Amount: 100, Type: Withdraw},
		{UserID: "user2", Amount: 200, Type: Deposit},
		{UserID: "user3", Amount: 100, Type: Withdraw},
		{UserID: "user1", Amount: 200, Type: Deposit},
		{UserID: "user2", Amount: 1500, Type: Deposit},
	}

	requests := make(chan Request)
	var wg sync.WaitGroup

	for i:=1; i<=3; i++ {
		wg.Add(1)
		go func (id int) {
			defer wg.Done()
			for req := range requests {
			fmt.Printf("Goroutine %d Processing %s for %s: %d$\n", id, req.Type, req.UserID, req.Amount)
			walletSystem.ProcessTransaction(req)
			time.Sleep(50*time.Millisecond)
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

	fmt.Println("All transication completed!\n Final Balances: ")
	for cardID, balance := range walletSystem.Users {
		fmt.Printf("%s: %d$\n", cardID, balance)
	}
}