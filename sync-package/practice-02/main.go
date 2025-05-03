package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i:=1; i<=5; i++ {
		wg.Add(1)
		go func (i int)  {
			defer wg.Done()

			time.Sleep(1*time.Second)
			fmt.Println(i)
		}(i)
	}

	wg.Wait()

}
