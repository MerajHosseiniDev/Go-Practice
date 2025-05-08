package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wallet int32 = 5000
	var wg sync.WaitGroup

	for i:=0; i<10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&wallet, -300)
		}()
	}

	wg.Wait()

	fmt.Println("Final balance: ", atomic.LoadInt32(&wallet))
}