package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	wg.Done()
	fmt.Println("worker", id, "started")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
		time.Sleep(1 * time.Second)
	}

	wg.Wait()
	fmt.Println("all worker finished")
}

// Завершение через close

package main

import (
	"fmt"
	"time"
)

func worker(jobs <-chan int) {
	for job := range jobs {
		fmt.Println("Processing job:", job)
		time.Sleep(time.Second)
	}
}

func main() {
	jobs := make(chan int)

	go worker(jobs)

	for i := 1; i <= 5; i++ {
		jobs <- i
	}

	close(jobs)

	time.Sleep(time.Second)
	fmt.Println("finish")
}