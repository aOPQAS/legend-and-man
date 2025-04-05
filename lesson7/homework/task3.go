package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	result := make(chan int, 1)
	sum := 0

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
		for sumSqrt := range result {
			sum += sumSqrt
		}
		fmt.Println(sum)
	}()

	wg.Wait()

}
