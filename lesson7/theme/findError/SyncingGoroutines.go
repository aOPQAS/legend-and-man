// Лёгкие задачи (Easy)

// 1️⃣ Ошибка в использовании sync.WaitGroup
// Ошибка: Главный поток завершится до того, как горутины успеют завершить свою работу, 
// потому что не используется правильная синхронизация с sync.WaitGroup.

package main

import (
	"fmt"
	"sync"
)

func worker() {
	fmt.Println("Goroutine is working!")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go worker()

	// Ошибка: главный поток завершится, прежде чем горутина завершит свою работу
	wg.Wait()
}

// 2️⃣ Синхронизация с sync.WaitGroup для нескольких горутин
// Ошибка: Главный поток не ждет завершения всех горутин, что приводит к тому, 
// что программа завершает выполнение раньше, чем горутины.

package main

import (
	"fmt"
	"sync"
)

func worker(id int) {
	fmt.Printf("Worker %d is working\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i)
	}

	// Ошибка: главный поток завершится раньше, чем все горутины
	wg.Wait()
}

// 3️⃣ Ошибка при использовании sync.Mutex для синхронизации доступа к переменной
// Ошибка: Несколько горутин могут одновременно изменять общую переменную, 
// что вызывает гонку данных.

package main

import (
	"fmt"
	"sync"
)

var counter int

func worker() {
	counter++
}

func main() {
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		mu.Lock()
		worker()
		mu.Unlock()
	}

	// Ошибка: доступ к `counter` не синхронизирован, гонка данных
	fmt.Println("Counter:", counter)
}

// Средние задачи (Medium)

// 4️⃣ Ошибки в передаче сигналов между горутинами
// Ошибка: Канал не используется правильно для синхронизации работы горутин, 
// из-за чего одна горутина завершает работу раньше, чем другие.

package main

import (
	"fmt"
	"sync"
)

func worker(id int, ch chan bool) {
	<-ch
	fmt.Printf("Worker %d has completed\n", id)
}

func main() {
	ch := make(chan bool)
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, ch)
		}(i)
	}

	// Ошибка: канал не синхронизирует работу горутин
	close(ch)
	wg.Wait()
}

// 5️⃣ Ошибки в использовании sync.WaitGroup с динамическим числом горутин
// Ошибка: Программа не учитывает изменение количества горутин в процессе выполнения, 
// из-за чего количество ожиданий в WaitGroup не совпадает с количеством горутин.

package main

import (
	"fmt"
	"sync"
)

func worker(id int) {
	fmt.Printf("Worker %d is working\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i)
	}

	// Ошибка: количество горутин не соответствует количеству вызовов `wg.Add`
	wg.Wait()
}

// 6️⃣ Ошибка с использованием sync.Mutex для защиты доступа к нескольким общим переменным
// Ошибка: При использовании нескольких горутин для изменения разных переменных не используется 
// должная синхронизация.

package main

import (
	"fmt"
	"sync"
)

var counter1 int
var counter2 int

func worker() {
	counter1++
	counter2++
}

func main() {
	var mu sync.Mutex
	for i := 0; i < 1000; i++ {
		mu.Lock()
		worker()
		mu.Unlock()
	}

	// Ошибка: одна блокировка защищает несколько переменных, но это не всегда эффективно
	fmt.Println("Counter1:", counter1)
	fmt.Println("Counter2:", counter2)
}

// Сложные задачи (Hard)

// 7️⃣ Сложности с sync.WaitGroup и каналами для синхронизации завершения горутин
// Ошибка: Слишком много горутин пытаются работать одновременно, но синхронизация 
// с sync.WaitGroup не выполнена корректно, что приводит к неожиданному завершению программы.

package main

import (
	"fmt"
	"sync"
)

func worker(id int, ch chan bool) {
	<-ch
	fmt.Printf("Worker %d has completed\n", id)
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan bool)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, ch)
		}(i)
	}

	// Ошибка: не управляем потоком сигналов для горутин
	close(ch)
	wg.Wait()
}

// 8️⃣ Горутины с таймаутом с помощью sync.Mutex
// Ошибка: Использование таймаутов для горутин не синхронизировано с помощью Mutex, 
// что может привести к неправильному завершению работы.

package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, mu *sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Printf("Worker %d started\n", id)
	time.Sleep(2 * time.Second)
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			worker(id, &mu)
		}(i)
	}

	// Ошибка: горутины не синхронизированы с помощью времени ожидания, программа завершится 
	// до завершения горутин
	wg.Wait()
}

// 9️⃣ Использование каналов и sync.Mutex для синхронизации нескольких горутин и данных
// Ошибка: Несколько горутин пытаются одновременно читать и записывать данные в общий канал, 
// но нет правильной синхронизации.

package main

import (
	"fmt"
	"sync"
)

var counter int

func worker(ch chan int, mu *sync.Mutex) {
	mu.Lock()
	counter += <-ch
	mu.Unlock()
}

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(ch, &mu)
		}()
	}

	for i := 0; i < 5; i++ {
		ch <- 1
	}

	// Ошибка: нет синхронизации для доступа к каналу, гонки данных
	wg.Wait()
	fmt.Println("Counter:", counter)
}