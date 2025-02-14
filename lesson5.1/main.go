package main

import (
	"log"
	"strings"
)

// start: go run main.go --old original_database.xml --new stolen_database.json
func main() {
	oldFilename, newFilename, err := ParseArgs()
	if err != nil {
		log.Fatal(err)
	}

	var recipesOld, recipesNew Recipes

	// old recipes
	if strings.Contains(oldFilename, ".xml") {
		recipesOld, err = ReadXML(oldFilename)
	} else if strings.Contains(oldFilename, ".json") {
		recipesOld, err = ReadJSON(oldFilename)
	}
	if err != nil {
		log.Fatal(err)
	}

	// new recipes
	if strings.Contains(newFilename, ".json") {
		recipesNew, err = ReadJSON(newFilename)
	} else if strings.Contains(newFilename, ".xml") {
		recipesNew, err = ReadXML(newFilename)
	}
	if err != nil {
		log.Fatal(err)
	}

	compareCakes(recipesOld, recipesNew)
}
