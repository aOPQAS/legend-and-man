package main

import (
	"fmt"
	"math"
	"sort"
)

func mainn() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("error while scanning number n:", err)
		return
	}

	if n <= 0 {
		fmt.Println("n can only be positive")
		return
	}

	array := make([]float64, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&array[i])
		if err != nil {
			fmt.Println("expected integer, try again...")
			i--
		}
	}
	median := Median(array)
	mean := Mean(array)
	sd := SD(array, mean)

	fmt.Println("Median:", median)
	fmt.Println("Mean:", mean)
	fmt.Println("SD:", sd)
}

// медиана
func Median(array []float64) float64 {
	sort.Float64s(array)
	n := len(array)

	if n%2 == 0 {
		return (array[n/2-1] + array[n/2]) / 2
	}
	return array[n/2]
}

// среднее значение
func Mean(array []float64) float64 {
	sum := 0.0
	for _, num := range array {
		sum += num
	}
	return sum / float64(len(array))
}

// отклонение
func SD(array []float64, mean float64) float64 {
	sum := 0.0
	for _, i := range array {
		sum += (i - mean) * (i - mean)
	}
	return math.Sqrt(sum / float64(len(array)))
}