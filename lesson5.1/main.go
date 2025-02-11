package main

import (
	"log"
)

// func main() {
// 	oldFilename, newFilename, err := ParseArgs()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var recipesOld, recipesNew Recipes

// 	if strings.Contains(oldFilename, ".xml") {
// 		recipesOld, err = ReadXML(oldFilename)
// 	} else if strings.Contains(oldFilename, ".json") {
// 		recipesOld, err = ReadJSON(oldFilename)
// 	}
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if strings.Contains(newFilename, ".json") {
// 		recipesNew, err = ReadJSON(newFilename)
// 	} else if strings.Contains(newFilename, ".xml") {
// 		recipesNew, err = ReadXML(newFilename)
// 	}
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	compareCakes(recipesOld, recipesNew)
// }

func main() {
	// Пример вызова
	recipesOld, err := ReadXML("original_database.xml")
	if err != nil {
		log.Fatal(err)
	}

	recipesNew, err := ReadJSON("stolen_database.json")
	if err != nil {
		log.Fatal(err)
	}

	compareCakes(recipesOld, recipesNew)
}
