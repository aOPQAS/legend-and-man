// Лёгкие задачи (Easy)

// 1️⃣ Ошибка с синхронизацией с sync.WaitGroup
// Ошибка: Программа не дожидается завершения горутины, и главный поток завершается до того, как горутина успевает выполнить свою задачу.

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		fmt.Println("Goroutine is working!")
// 	}()

// 	wg.Wait()

// }

// 2️⃣ Горутина пытается использовать переменную, которая изменяется в главном потоке
// Ошибка: Переменная изменяется одновременно в главной функции и горутине, что может привести к гонкам данных.

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	var counter int
// 	var mu sync.Mutex

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for i := 0; i < 5; i++ {
// 			mu.Lock()
// 			counter++
// 			mu.Unlock()
// 			time.Sleep(time.Millisecond * 100)
// 		}
// 	}()

// 	wg.Wait()

// 	mu.Lock()
// 	fmt.Println("Counter:", counter)
// 	mu.Unlock()
// }

// 3️⃣ Горутина не завершена из-за блокировки канала
// Ошибка: Программа не завершает горутину, так как канал заблокирован и горутина ждет на канале.

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan string)

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		fmt.Println(<-ch)
// 	}()

// 	ch <- "hello"

// 	close(ch)
// 	wg.Wait()
// }

// Средние задачи (Medium)

// 4️⃣ Ошибка с горутинами и циклом
// Ошибка: Горутины не правильно обрабатывают индекс из цикла, из-за чего все горутины работают с одинаковым значением индекса.

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 1; i <= 3; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			fmt.Println("Goroutine:", i)
// 		}()
// 	}

// 	wg.Wait()

// }

// 5️⃣ Горутины не синхронизированы с помощью sync.WaitGroup
// Ошибка: Горутины не дожидаются завершения, что вызывает проблемы с синхронизацией.

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func worker(id int) {
// 	fmt.Printf("Worker %d started\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 0; i < 3; i++ {
// 		wg.Add(1)
// 		go func(id int) {
// 			defer wg.Done()
// 			worker(id)
// 		}(i)
// 	}

// 	wg.Wait()

// }

// 6️⃣ Горутины с тайм-аутом
// Ошибка: Тайм-аут используется неправильно, что приводит к преждевременному завершению горутины.

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch := make(chan bool)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- true
// 	}()

// 	select {
// 	case <-ch:
// 		fmt.Println("Task completed")
// 	case <-time.After(3 * time.Second):
// 		fmt.Println("Timed out")
// 	}
// }

// Сложные задачи (Hard)

// 7️⃣ Горутины с обработкой ошибок
// Ошибка: Ошибка в одной горутине не учитывается и не обрабатывается должным образом, что вызывает проблемы в дальнейшем.

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func worker(id int, wg *sync.WaitGroup, ch chan error) {
// 	defer wg.Done()

// 	if id%2 == 0 {
// 		ch <- fmt.Errorf("worker %d encountered an error", id)
// 	} else {
// 		fmt.Printf("worker %d completed successfully\n", id)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan error, 5)

// 	for i := 0; i < 5; i++ {
// 		wg.Add(1)
// 		go worker(i, &wg, ch)
// 	}

// 	wg.Wait()
// 	close(ch)

// 	for err := range ch {
// 		if err != nil {
// 			fmt.Println("Error:", err)
// 		}
// 	}
// 	// Ошибка: нужно правильно обработать ошибку, а не просто игнорировать
// }

// 8️⃣ Горутины с доступом к общим данным без синхронизации
// Ошибка: Несколько горутин изменяют общую переменную без использования синхронизации, что может привести к гонкам данных.

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var counter int

// func worker(wg *sync.WaitGroup, mu *sync.Mutex) {
// 	defer wg.Done()
// 	mu.Lock()
// 	counter++
// 	mu.Unlock()
// }

// func main() {
// 	var wg sync.WaitGroup
// 	var mu sync.Mutex

// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go worker(&wg, &mu)
// 	}

// 	wg.Wait()

// 	mu.Lock()
// 	fmt.Println("Counter:", counter)
// 	mu.Unlock()
// }

// 9️⃣ Неверное использование каналов для координации горутин
// Ошибка: Использование каналов не для их прямого назначения (для передачи данных), что вызывает неправильное поведение программы.

// надо деделать!!

package main

import "fmt"

func worker(ch chan bool) {
	// Ошибка: горутина неправильно использует канал для передачи сигналов
	<-ch
	fmt.Println("Worker completed")
}

func main() {
	ch := make(chan bool)

	for i := 0; i < 5; i++ {
		go worker(ch)
	}

	// Ошибка: мы не отправляем сигнал в канал для запуска горутин
	close(ch)
}
