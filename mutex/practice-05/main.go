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

	for depositAmount := range depositRequest{
		time.Sleep(500*time.Millisecond)

		b.mu.Lock()
		b.Balance += depositAmount
		b.mu.Unlock()

		fmt.Printf("Deposit of %d$ successful! New balance: %d$\n", depositAmount, b.Balance)
	}
}

func (b *BankAccount) CheckBalance() int {
	b.mu.RLock()
	defer b.mu.RUnlock()
	
	return b.Balance
}

func (b *BankAccount) withdraw(wg *sync.WaitGroup, amount int) bool {
	defer wg.Done()
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.Balance >= amount {
		time.Sleep(508*time.Millisecond)
		b.Balance -= amount
		fmt.Printf("Withdrawal of %d$ successful! New balance: %d$\n", amount, b.Balance)
		return true
	} else {
		fmt.Printf("Insufficient funds: Tried to withdraw %d$, But only %d$ available\n", amount, b.Balance)	
		return false
	}
	
}


func main() {
	var wg sync.WaitGroup
	BankAccount := BankAccount{Balance: 200}
	depositRequest := make(chan int)
	depositAmounts := []int{100,1500,50,40}
	withdrawAmounts := []int{1000,200,400,100,100}

	wg.Add(1)
	go BankAccount.deposit(&wg, depositRequest)

	for _, amount := range depositAmounts {
		depositRequest <- amount
	}
	
	close(depositRequest)
	wg.Wait()

	for _, wAmount := range withdrawAmounts {
		wg.Add(1) 
		go BankAccount.withdraw(&wg, wAmount)
	}

	wg.Wait()

	fmt.Printf("Final Current Balance: %d$\n", BankAccount.CheckBalance())

}