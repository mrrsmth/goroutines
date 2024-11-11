package main

import (
	"fmt"
)

func main() {
	chanFactorial := make(chan int)
	intChan := make(chan int)
	iCh := make(chan int, 3)
	//1
	iCh <- 10
	iCh <- 3
	iCh <- 24
	fmt.Println(<-iCh)
	fmt.Println(<-iCh)
	fmt.Println(<-iCh)
	//2
	go factorial(5, chanFactorial)
	resultFactorial := <-chanFactorial
	fmt.Println(resultFactorial)
	//3
	go func() {
		fmt.Println("5 in in intChan")
		intChan <- 5
	}()

	fmt.Println(<-intChan)
	fmt.Println("finish")
}

func factorial(n int, ch chan<- int) {
	if n < 1 {
		fmt.Println("Invalid input number")
		return
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	ch <- result
}
