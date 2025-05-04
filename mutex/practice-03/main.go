package main

import (
	"fmt"
	"sync"
)

var (
	count int
	mu	  sync.Mutex
)

func increament(wg *sync.WaitGroup){
	defer wg.Done()
	mu.Lock()
	defer mu.Unlock()
	count++
}

func main() {
	var wg sync.WaitGroup

	for i:=0; i < 100; i++ {
		wg.Add(1)
		go increament(&wg)
	}

	wg.Wait()
	fmt.Println("Final Count: ", count)

}