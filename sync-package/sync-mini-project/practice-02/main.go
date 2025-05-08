package main

import (
	"fmt"
	"sync"
	"time"
)

var	Balance int = 1000


func processWithdrawals(wg *sync.WaitGroup, ch <- chan int) {
	defer wg.Done()

	for amount := range ch {
		if Balance >= amount {
			fmt.Printf("Withdrawal %d$ ...\n", amount)
			Balance -= amount
			time.Sleep(1 * time.Second)
			fmt.Printf("withdrawal of %d$ successful.\nRemaining balance: %d$\n", amount, Balance)
		} else {
			fmt.Printf("Insufficient funds for withdrawal of %d$.\nCurrent balance: %d$\n", amount, Balance)
		}
		time.Sleep(500 * time.Millisecond)		
	}
}

func main() {
	amounts := []int{100, 200, 200, 500, 50}
	withdrawalrequests := make(chan int)
	var wg sync.WaitGroup


	fmt.Printf("Hello and welcome\nBalance: %d$\n", Balance)

	wg.Add(1)
	go processWithdrawals(&wg, withdrawalrequests)

	for _, amount := range amounts {
		withdrawalrequests <- amount
	}

	close(withdrawalrequests)

	wg.Wait()
	fmt.Println("Final Amount balance: ", Balance)

}
