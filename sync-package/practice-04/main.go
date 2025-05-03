package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range nums {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			nums[i] *= 2
		}(i)
	}
	wg.Wait()
	fmt.Println("Result: ", nums)
}
