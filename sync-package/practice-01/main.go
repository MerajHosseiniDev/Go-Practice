package main

import (
	"fmt"
	"sync"
)

func task1(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(id)
}

func task2(sentence string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(sentence)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go task1(10, &wg)
	go task2("Can you see me?", &wg)

	wg.Wait()

}