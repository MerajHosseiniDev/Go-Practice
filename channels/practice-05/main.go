package main

import (
	"fmt"
	"sync"
)


func main() {
	nuums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	part1 := nuums[:5]
	part2 := nuums[5:]
	ch := make(chan int)
	var wg sync.WaitGroup

	wg.Add(3)

	go func ()  {
		defer wg.Done()
		sum1 := 0
		for _, num := range part1 {
			sum1 += num
		}
		ch <- sum1
	}()

	go func ()  {
		defer wg.Done()
		sum2 := 0
		for _, num := range part2 {
			sum2+=num
		}
		ch <- sum2
	}()

	go func ()  {
		defer wg.Done()
		sum1 := <- ch
		sum2 := <- ch
		fmt.Println("Final total: ", sum1 + sum2)
	}()

	wg.Wait()
	close(ch)
	
	fmt.Println("Program completely worked!")
}