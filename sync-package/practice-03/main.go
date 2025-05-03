package main

import (
	"fmt"
	"sync"
	"time"
)

func file1(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Start downloading file1 by plus service!(%0.0)")
	time.Sleep(1 * time.Second)
	fmt.Println("Downloading file1.. (32.42%)")
	time.Sleep(1 * time.Second)
	fmt.Println("Download file1 completed! (100%)")
}

func file2(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Start downloading file2 by normal service! (0.0%)")
	time.Sleep(1 * time.Second)
	fmt.Println("Downloading file2.. (24.99%)")
	time.Sleep(1 * time.Second)
	fmt.Println("Downloading file2.. (65.12%)")
	time.Sleep(1 * time.Second)
	fmt.Println("Download file2 completed! (100%)")
}

func file3(wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Start downloading file3 by VIP service!(0.0%)")
	time.Sleep(1 * time.Second)
	fmt.Println("Download file3 completed!(100%)")
}

func main() {
	var wg sync.WaitGroup

	fmt.Println("Starting all downloads...")
	
	wg.Add(3)
	go file1(&wg)
	go file2(&wg)
	go file3(&wg)

	wg.Wait()
	fmt.Println("All files downloaded!")
	
}
