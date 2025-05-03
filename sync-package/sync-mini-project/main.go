package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	Balance int = 1000
	mu      sync.Mutex
)

func withdraw(wg *sync.WaitGroup, amount int) {
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()

	if Balance >= amount {
		fmt.Printf("Withdrawal %d$ ...", amount)
		Balance -= amount
		time.Sleep(1 * time.Second)
		fmt.Printf("withdrawal of %d$ successful.\nRemaining balance: %d$\n", amount, Balance)
	} else {
		fmt.Printf("Insufficient funds for withdrawal of %d$.\nCurrent balance: %d$", amount, Balance)
	}
	time.Sleep(500 * time.Millisecond)
}

func main() {
	var wg sync.WaitGroup
	amounts := []int{100, 200, 200, 500, 50}

	fmt.Printf("Hello and welcome\nBalance: %d$\n", Balance)

	wg.Add(len(amounts))
	for _, amount := range amounts {
		go withdraw(&wg, amount)
	}

	wg.Wait()
	fmt.Println("Final Amount balance: ", Balance)

}
