// Лёгкие задачи (Easy)
// 1️⃣ Ошибка с чтением из закрытого канала

// Ошибка: Канал закрывается, но код пытается читать из него, не проверив, что канал открыт. Нужно исправить, чтобы избежать паники.

package main

import "fmt"

func main() {
    ch := make(chan int)

    close(ch)
    fmt.Println(<-ch) // Ошибка: невозможно читать из закрытого канала
}

// 2️⃣ Ошибка с каналом и sync.WaitGroup

// Ошибка: WaitGroup не используется корректно, что приводит к проблемам с синхронизацией.

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- "Hello"
	}()

	fmt.Println(<-ch)
	wg.Wait() // Ошибка: WaitGroup не синхронизирован с завершением горутины
}

// 3️⃣ Ошибка с блокировкой канала

// Ошибка: Канал переполняется, и программа блокируется. Нужно реализовать механизм отправки и получения данных через канал без блокировки.

package main

import "fmt"

func main() {
	ch := make(chan int) // Канал без буфера

	ch <- 1
	ch <- 2 // Ошибка: канал переполнен, блокировка
	fmt.Println(<-ch)
}

// Средние задачи (Medium)
// 4️⃣ Неправильная работа с каналом в горутинах

// Ошибка: Канал не закрывается после отправки данных. Из-за этого горутины продолжают работать, даже если данные уже переданы.

package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Message 1"
		ch <- "Message 2"
		ch <- "Message 3"
		// Канал не закрывается, из-за чего главная горутина может ждать навсегда.
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}

// 5️⃣ Ошибка с каналами и select

// Ошибка: Использование select без корректной обработки случая, когда все каналы заняты.

package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	select {
	case msg := <-ch1:
		fmt.Println("Received from ch1:", msg)
	case msg := <-ch2:
		fmt.Println("Received from ch2:", msg)
	// Ошибка: нет default case и оба канала не содержат данных, программа зависнет
	}
}

// 6️⃣ Проблема с чтением данных из канала после его закрытия

// Ошибка: Программа пытается читать из закрытого канала, но не учитывает, 
// что канал может быть уже закрыт.

package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch)

	// Ошибка: чтение из закрытого канала без проверки на состояние канала
	fmt.Println(<-ch) // Программа зависнет
}

// Сложные задачи (Hard)
// 7️⃣ Горутины с блокировкой из-за попытки записи в канал после его закрытия

// Ошибка: Канал закрыт, но ещё происходит запись в него, что вызывает блокировку. 
// Hужно исправить код, чтобы избежать записи в закрытый канал.

package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch) // Канал закрыт

	// Ошибка: попытка записать в закрытый канал
	ch <- 10
	ch <- 20
}

// 8️⃣ Ошибка с несколькими горутинами, работающими с каналом

// Ошибка: Несколько горутин пытаются работать с одним каналом, но не используют правильную 
// синхронизацию, что может привести к панике или блокировке.

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
		go worker(ch) // Ошибка: несколько горутин не синхронизированы с каналом
	}

	for i := 0; i < 5; i++ {
		ch <- i
	}

	close(ch) // Ошибка: все горутины могут не успеть завершиться до закрытия канала
}

// 9️⃣ Программа зависает из-за ошибок синхронизации канала с sync.WaitGroup

// Ошибка: sync.WaitGroup не используется правильно, из-за чего горутины не завершатся.

package main

import (
	"fmt"
	"sync"
)

func worker(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(<-ch)
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(1)
	go worker(ch, &wg)

	ch <- "Task 1"
	ch <- "Task 2"
	wg.Wait() // Ошибка: после завершения горутины канал не закрыт, могут быть проблемы с чтением
}