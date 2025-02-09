package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

func ParseArgs() (string, error) {
	if len(os.Args) != 3 {
		return "", fmt.Errorf("failed to parse args, must be 3 values")
	}

	if os.Args[1] != "-f" {
		return "", fmt.Errorf("missed option -f")
	}

	return os.Args[2], nil
}

type DBReader interface {
	Read(string) error
}

func ReadJSON(filename string) error {
	fd, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var recipes Recipes

	err = json.Unmarshal(fd, &recipes)
	if err != nil {
		return fmt.Errorf("failed to unmarshall data into recipes: %w", err)
	}

	data, _ := xml.MarshalIndent(recipes, "", "    ")
	fmt.Println(string(data))
	return nil
}

func ReadXML(filename string) error {
	fd, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	var recipes Recipes

	err = xml.Unmarshal(fd, &recipes)
	if err != nil {
		return fmt.Errorf("failed to unmarshall data into recipes: %w", err)
	}

	data, _ := json.MarshalIndent(recipes, "", "    ")
	fmt.Println(string(data))
	return nil
}

func main() {
	filename, err := ParseArgs()
	if err != nil {
		log.Fatal(err)
	}

	err = ReadJSON(filename)
	if err != nil {
		log.Fatal(err)
	}
}
