package main

import (
	"fmt"
	"os"
)

func main() {
	file := "gogogo.txt"

	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Println("Ошибка при чтении файла: ", err)
		return
	}

	fmt.Println(string(content))
}
