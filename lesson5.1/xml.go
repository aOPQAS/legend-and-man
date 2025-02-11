package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

func ReadXML(filename string) (Recipes, error) {
	fd, err := os.ReadFile(filename)
	if err != nil {
		return Recipes{}, fmt.Errorf("failed to read file: %w", err)
	}

	var recipes Recipes

	err = xml.Unmarshal(fd, &recipes)
	if err != nil {
		return Recipes{}, fmt.Errorf("failed to unmarshall data into recipes: %w", err)
	}

	return recipes, nil
}
