package mutex

import (
	"fmt"
	"sync"
)

var counter int = 0 //  общий ресурс

func Mutex() {
	fmt.Println("init Mutex")
}

func Work(number int, ch chan bool) {
	counter = 0
	for k := 1; k <= 5; k++ {
		counter++
		fmt.Println("Goroutine", number, "-", counter)
	}
	ch <- true
}

func WorkMutex(number int, ch chan bool, mutex *sync.Mutex) {
	mutex.Lock() // блокируем доступ к переменной counter
	counter = 0
	for k := 1; k <= 5; k++ {
		counter++
		fmt.Println("Goroutine", number, "-", counter)
	}
	mutex.Unlock() // деблокируем доступ
	ch <- true
}
