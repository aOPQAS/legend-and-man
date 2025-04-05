package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.Mutex
	data := make(map[int]string)
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			data[i] = fmt.Sprintf("Value %d", i)
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	mu.Lock()
	for key, value := range data {
		fmt.Println(key, value)
	}
	mu.Unlock()
}
