// The mini project was practiced here once again. Go-Practice/sync-package/sync-mini-project/main.go

package main

import (
	"fmt"
	"sync"
	"time"
)

var balance int = 1000

func withdrawalprocess(wg *sync.WaitGroup, ch <- chan int) {
	defer wg.Done()
	for amount := range ch {
		if balance >= amount {
			fmt.Printf("Withdrawal %d$ ...", amount)
			balance -= amount
			time.Sleep(1*time.Second)
			fmt.Printf("Withdrawal of %d$ successful!\nRemaining balance: %d$\n", amount, balance)
		} else {
			fmt.Printf("Insufficient funds for withdrawal of %d$.\nCurrent balance: %d$\n", amount, balance)
		}
		time.Sleep(500*time.Millisecond)
	}
}


func main() {
	amounts := []int{100, 200, 200, 500, 50}
	var wg sync.WaitGroup
	withdrawalRequest := make(chan int)

	fmt.Printf("Hello and Welcome!\nBalance: %d$\n", balance)

	wg.Add(1)
	
	go withdrawalprocess(&wg, withdrawalRequest)

	for _, amount := range amounts {
		withdrawalRequest <- amount
	}

	close(withdrawalRequest)
	wg.Wait()

	fmt.Printf("Final amount balance: %d$\n", balance)



}