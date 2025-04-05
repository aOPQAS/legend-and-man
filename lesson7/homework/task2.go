package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	result := make(chan int, 1)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, num := range numbers {
			result <- num * num
		}
		close(result)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for res := range result {
			fmt.Println(res)
		}
	}()

	wg.Wait()
}
