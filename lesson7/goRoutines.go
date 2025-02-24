package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	t := time.Now()
	rand.Seed(t.UnixNano())

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		ParseURL("http:/example.com/")
	}()

	ParseURL("http:/example.com/")

	wg.Wait()

	fmt.Printf("Parsing completed. Time Elapsed: %.2f seconds\n", time.Since(t).Seconds())
}

func ParseURL(url string) {
	for i := 0; i < 5; i++ {
		latency := rand.Intn(500) + 500

		time.Sleep(time.Duration(latency) * time.Millisecond)

		fmt.Printf("parsing <%s> - Step %d - Latency %d ms\n", url, i+1, latency)
	}
}
