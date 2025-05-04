package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	count int
	mu	  sync.Mutex
}

func (c *Counter) increment(wg *sync.WaitGroup){
	defer wg.Done()
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *Counter) decrement(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count--
}

func main() {
	var wg sync.WaitGroup
	counter := &Counter{}

	for i:=0; i < 100; i++ {
		wg.Add(1)
		go counter.increment(&wg)
	}

	for i:=0; i<5; i++ {
		wg.Add(1)
		go counter.decrement(&wg)
	}

	wg.Wait()
	fmt.Println("Final Count: ", counter.count)

}