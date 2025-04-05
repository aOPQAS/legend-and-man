// package main

// import (
// 	"fmt"
// 	"strings"
// )

// // Find Closest Number to Max
// func findMax(array []int) int {
// 	max := array[0]
// 	for _, num := range array {
// 		if num > max {
// 			max = num
// 		}
// 	}
// 	return max
// }

// // Check if a string is a palindrome
// func isPalindrome(s string) bool {
// 	s = strings.ToLower(s)
// 	n := len(s)
// 	for i := 0; i < n/2; i++ {
// 		if s[i] != s[n-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }

// // Fibonacci sequence
// func fibonacci(n int) int {
// 	if n <= 1 { // 1, 2, 2, 3, 3, 5, 5, 8 tsss selo!!
// 		return n
// 	}
// 	a, b := 0, 1
// 	for i := 2; i <= n; i++ {
// 		a, b = b, a+b
// 	}
// 	return b
// }

// // Reverse a string
// func reverseString(s string) string {
// 	result := ""
// 	for i := len(s) - 1; i >= 0; i-- {
// 		result += string(s[i])
// 	}
// 	return result
// }

// // Count duplicate characters in a string
// func countDuplicates(s string) int {
// 	charCount := make(map[rune]int)
// 	s = strings.ToLower(s)
// 	for _, char := range s {
// 		charCount[char]++
// 	}

// 	duplicateCount := 0
// 	for _, count := range charCount {
// 		if count > 1 {
// 			duplicateCount++
// 		}
// 	}
// 	return duplicateCount
// }

// // Search Insert Position
// func searchInsert(nums []int, target int) int {
// 	for i, num := range nums {
// 		if num >= target {
// 			return i
// 		}
// 	}
// 	return len(nums)
// }

// func main() {
// 	fmt.Println("Max number:", findMax([]int{-4, -2, -11, -9, -8}))
// 	fmt.Println("Is palindrome:", isPalindrome("Racecar"))
// 	fmt.Println("Is palindrome:", isPalindrome("hello"))
// 	fmt.Println("Fibonacci(10):", fibonacci(10))
// 	fmt.Println("Reversed string:", reverseString("agagaga"))
// 	fmt.Println("Duplicate count:", countDuplicates("abcde"))
// 	fmt.Println("Duplicate count:", countDuplicates("asgogoogogoa"))
// 	fmt.Println("Insert position:", searchInsert([]int{5, 61, 754, 8, 23, 4, 5, 676}, 8))
// }