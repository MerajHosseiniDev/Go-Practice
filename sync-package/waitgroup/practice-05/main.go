package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	
	go func() {
		defer wg.Done()
		fmt.Println("message from goroutine.")
	}()

	wg.Wait()
	fmt.Println("Program Completely worked!")
}