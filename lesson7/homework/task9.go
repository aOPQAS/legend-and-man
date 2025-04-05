package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	ch1 := make(chan int)
	ch2 := make(chan int)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, num := range numbers {
			ch1 <- num
		}
		close(ch1)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch1 {
			ch2 <- num * 2
		}
		close(ch2)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for result := range ch2 {
			fmt.Println(result)
		}
	}()

	wg.Wait()
}
