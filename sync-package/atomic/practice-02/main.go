package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var count int32 = 0

	for i:=0; i<5; i++ {
		success := atomic.CompareAndSwapInt32(&count, 0, 1)
		if success {
			fmt.Println("Login allowed!")
		} else {
			fmt.Println("Already logged in!")
		}
	}
}