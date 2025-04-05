// Лёгкие задачи (Easy)

// 1️⃣ Ошибка с чтением из закрытого канала
// Ошибка: Канал закрывается, но код пытается читать из него, не проверив, что канал открыт. Нужно исправить, чтобы избежать паники.

// package main

// import "fmt"

// func main() {
// 	ch := make(chan int)
// 	close(ch)

// 	val, ok := <- ch 
// 	if !ok {
// 		fmt.Println("Канал зыкрыт")
// 	} else {
// 		fmt.Println("Канал открыт")
// 	}
// }

// 2️⃣ Ошибка с каналом и sync.WaitGroup
// Ошибка: WaitGroup не используется корректно, что приводит к проблемам с синхронизацией.

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan string, 1)

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		fmt.Println(<-ch)
// 	}()

// 	ch <- "Hello"
// 	close(ch)
// 	wg.Wait() 
// }

// 3️⃣ Проблема с блокировкой канала
// Ошибка: Канал переполняется, и программа блокируется. Нужно реализовать механизм отправки и получения данных через канал без блокировки.

// package main

// import "fmt"

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan int, 3) 

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		for val := range ch {
// 			fmt.Println(val)
// 		}
// 	}()

// 	ch <- 5
// 	ch <- 12
// 	ch <- 69

// 	close(ch)
// 	wg.Wait()
// }

// Средние задачи (Medium)

// 4️⃣ Неправильная работа с каналом в горутинах
// Ошибка: Канал не закрывается после отправки данных. Из-за этого горутины продолжают работать, даже если данные уже переданы.

// package main

// import "fmt"

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan string)

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		ch <- "Message 1"
// 		ch <- "Message 2"
// 		ch <- "Message 3"
// 		close(ch)
// 	}()
	
// 	wg.Wait()

// 	for msg := range ch {
// 		fmt.Println(msg)
// 	}

// }

// 5️⃣ Ошибка с каналами и select
// Ошибка: Использование select без корректной обработки случая, когда все каналы заняты.

// package main

// import "fmt"

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	select {
// 	case msg := <-ch1:
// 		fmt.Println("Received from ch1:", msg)
// 	case msg := <-ch2:
// 		fmt.Println("Received from ch2:", msg)
// 	default:
// 		fmt.Println("Нет доступных данных")
// 	}
// }

// 6️⃣ Проблема с чтением данных из канала после его закрытия
// Ошибка: Программа пытается читать из закрытого канала, но не учитывает, что канал может быть уже закрыт.

// package main

// import "fmt"

// func main() {
// 	ch := make(chan int)
// 	close(ch)

// 	val, ok := <- ch
// 		if !ok {
// 			fmt.Println("канал закрыт")
// 		} else {
// 			fmt.Println("Канал открыт")
// 		}
// }

// Сложные задачи (Hard)

// 7️⃣ Горутины с блокировкой из-за попытки записи в канал после его закрытия
// Ошибка: Канал закрыт, но ещё происходит запись в него, что вызывает блокировку. Нужно исправить код, чтобы избежать записи в закрытый канал.

package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	close(ch)

	ch <- 10
	ch <- 20
}

// 8️⃣ Ошибка с несколькими горутинами, работающими с каналом
// Ошибка: Несколько горутин пытаются работать с одним каналом, но не используют правильную синхронизацию, что может привести к панике или блокировке.

package main

import "fmt"

func worker(ch chan int) {
	for val := range ch {
		fmt.Println("Processing:", val)
	}
}

func main() {
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		go worker(ch)
	}

	for i := 0; i < 5; i++ {
		ch <- i
	}

	close(ch) 
}

// 9️⃣ Программа зависает из-за ошибок синхронизации канала с sync.WaitGroup
// Ошибка: sync.WaitGroup не используется правильно, из-за чего горутины не завершатся.

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func worker(ch chan string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Println(<-ch)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan string, 2)

// 	for i := 1; i <= 2; i++ {
// 		wg.Add(1)
// 		go worker(ch, &wg)
// 	}

// 	ch <- "Task 1"
// 	ch <- "Task 2"

// 	close(ch)
// 	wg.Wait()
// }