package main

import (
	"log"
	"os"
)

// the file should read like this: ./compareFS --old snapshot1.txt --new snapshot2.txt

func main() {
	oldFile := os.Args[2]
	newFile := os.Args[4]

	err := compareFiles(oldFile, newFile)
	if err != nil {
		log.Fatal("Error:", err)
	}

}
