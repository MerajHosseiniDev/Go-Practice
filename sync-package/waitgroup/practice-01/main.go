package main

import (
	"fmt"
	"sync"
	"time"
)

func task1(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task1 started!")
	time.Sleep(2*time.Second)
	fmt.Println("Task1 result: ",id)
	fmt.Println("Task1 finished!")
}

func task2(sentence string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Task1 started!")
	time.Sleep(1*time.Second)
	fmt.Println("Task2 result: ", sentence)
	fmt.Println("Task2 finished!")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go task1(10, &wg)
	go task2("Can you see me?", &wg)

	wg.Wait()
	fmt.Println("All tasks Completed!")

}