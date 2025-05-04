package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{1, 2, 3, 4, 5}
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		sum := 0
		for _, num := range nums {
			sum += num
		}
		ch <- sum
		close(ch)
	}()

	go func() {
		defer wg.Done()
		sum := <-ch
		fmt.Println(sum)
	}()

	wg.Wait()
	fmt.Println("Program completely worked!")

}
