package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isGreaterThan100(n	int, ch	chan bool) {
	if n > 100 {
		ch <- true
	} else {
		ch <- false
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Enter a number: ")

	input, _ := reader.ReadString('\n')

	number, err := strconv.Atoi(input[:len(input)-1])
	if err != nil {
		fmt.Println("Invalid input! Please input a valid number!")
		return
	}

	resultChan := make(chan bool)

	go isGreaterThan100(number, resultChan)

	isGreaterThan100 := <- resultChan

	if isGreaterThan100 {
		fmt.Println("The number is greater than 100.")
	} else {
		fmt.Println("The number isn't greater than 100.")
	}

}