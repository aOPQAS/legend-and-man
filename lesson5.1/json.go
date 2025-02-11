package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadJSON(filename string) (Recipes, error) {
	fd, err := os.ReadFile(filename)
	if err != nil {
		return Recipes{}, fmt.Errorf("failed to read file: %w", err)
	}

	var recipes Recipes

	err = json.Unmarshal(fd, &recipes)
	if err != nil {
		return Recipes{}, fmt.Errorf("failed to unmarshall data into recipes: %w", err)
	}

	return recipes, nil
}
