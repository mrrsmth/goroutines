package main

import (
	"fmt"
	m "goroutines/mutex"
	wait "goroutines/wait"
	"sync"
	"time"
)

func main() {
	chanFactorial := make(chan int)
	intChan := make(chan int)
	iCh := make(chan int, 3)
	closeCh := make(chan int, 3)
	results := make(map[int]int)
	structCh := make(chan struct{})
	streamCh := make(chan int)
	workCh := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2) // в группе две горутины

	// 1
	iCh <- 10
	iCh <- 3
	iCh <- 24
	fmt.Println(<-iCh)
	fmt.Println(<-iCh)
	fmt.Println(<-iCh)

	// 2
	go factorial(5, chanFactorial)
	resultFactorial := <-chanFactorial
	fmt.Println(resultFactorial)

	// 3
	go func() {
		fmt.Println("5 in intChan")
		intChan <- 5
	}()

	fmt.Println(<-intChan)
	fmt.Println("finish")

	//4
	closeCh <- 10
	closeCh <- 3
	close(closeCh)
	// closeCh <- 24       // chan close
	fmt.Println(<-closeCh) // 10
	fmt.Println(<-closeCh) // 3
	fmt.Println(<-closeCh) // 0
	for value := range closeCh {
		fmt.Printf("Value: %d\n", value)
	}

	for i := 0; i < cap(closeCh); i++ {
		if val, opened := <-closeCh; opened {
			fmt.Println(val)
		} else {
			fmt.Println("Channel closed!")
		}
	}

	// 5
	go chFactorial(7, structCh, results)

	<-structCh

	for i, v := range results {
		fmt.Println(i, " - ", v)
	}

	// 6
	go streamFactorial(7, streamCh)

	for {
		num, opened := <-streamCh
		if !opened {
			break
		}
		fmt.Println(num)
	}

	for num := range streamCh {
		fmt.Println(num)
	}

	// Mutex
	m.Mutex()
	for i := 1; i < 5; i++ {
		go m.Work(i, workCh)
	}
	// ожидаем завершения всех горутин
	for i := 1; i < 5; i++ {
		<-workCh
	}
	fmt.Println("The End")

	var mutex sync.Mutex
	for i := 1; i < 5; i++ {
		go m.WorkMutex(i, workCh, &mutex)
	}
	// ожидаем завершения всех горутин
	for i := 1; i < 5; i++ {
		<-workCh
	}
	fmt.Println("The End Mutex")

	// Wait
	wait.Wait()

	work := func(id int) {
		defer wg.Done()
		fmt.Printf("Горутина %d начала выполнение \n", id)
		time.Sleep(2 * time.Second)
		fmt.Printf("Горутина %d завершила выполнение \n", id)
	}
	// вызываем горутины
	go work(1)
	go work(2)

	wg.Wait() // ожидаем завершения обоих горутин
	fmt.Println("Горутины завершили выполнение")
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
	close(ch)
}

func chFactorial(n int, ch chan struct{}, results map[int]int) {
	defer close(ch)
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		results[i] = result
	}
}

func streamFactorial(n int, ch chan int) {
	defer close(ch)
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
		ch <- result // посылаем по числу
	}
}
