package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	n := 10
	c := make(chan int)

	var mu sync.Mutex

	go func() {
		for i := 0; i < n; i++ {
			mu.Lock()
			c <- i
			mu.Unlock()
			time.Sleep(1 * time.Second)
		}
		close(c)
	}()

	timeout := time.After(10 * time.Second)

	for {
		select {
		case val, ok := <-c:
			if !ok {
				fmt.Println("Канал закрыт")
				return
			}
			fmt.Println("Получено:", val+1)
		case <-timeout:
			fmt.Println("Время закончилось")
			return
		}
	}
}
