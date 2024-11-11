package main

import (
	"fmt"
	"time"
)

func main() {
	resultFactorial := make(chan int)
	intChan := make(chan int)

	go factorial(5, resultFactorial)

	fmt.Println(<-resultFactorial)

	go func() {
		fmt.Println("5 in in intChan")
		intChan <- 5
	}()
	fmt.Println(<-intChan)
	fmt.Println("finish")
}

func factorial(n int, ch chan int) {
	if n < 1 {
		fmt.Println("Invalid input number")
		return
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	ch <- result
	time.Sleep(100 * time.Millisecond)
}
