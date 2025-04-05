package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

func worker(id int, ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		fmt.Printf("Worker %d: %s\n", id, data)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <number of workers>")
		return
	}

	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Println("Please provide a valid number of workers.")
		return
	}

	dataChannel := make(chan string)

	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChannel, &wg)
	}

	go func() {
		for i := 1; ; i++ {
			dataChannel <- fmt.Sprintf("Data %d", i)
		}
	}()

	wg.Wait()
}
