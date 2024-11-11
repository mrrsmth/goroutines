package wait

import (
	"fmt"
	"sync"
	"time"
)

func Wait() {
	// создаем WaitGroup
	wg := sync.WaitGroup{}
	// цикл с 5-ю итерациями
	for i := 0; i < 5; i++ {
		// добавляем в список ожидания одну горутину
		wg.Add(1)
		go func(i int) {
			// говорим, чтобы в конце анонимной функции одна горутина из списка ожидания исчезла
			defer wg.Done()
			// засыпаем, имитируя какую-то работу
			time.Sleep(time.Duration(i) * time.Second)
			fmt.Println("Горутина", i, "завершила свое выполнение!!!")
		}(i)
	}
	// ожидаем незавершившиеся горутины
	wg.Wait()
	fmt.Println("Все горутины завершили свое выполнение!!!!!!")
}
