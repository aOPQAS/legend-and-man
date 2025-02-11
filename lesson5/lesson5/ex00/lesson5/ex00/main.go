package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
)
// go build *.go -o compareDB
// ~$ ./compareDB --old original_database.xml --new stolen_database.json
// os.Args [./compareDB, --old, original_database.xml, --new, stolen_database.json]
func ParseArgs() (string, error) {
	if len(os.Args) != 5 {
		return "", "", fmt.Errorf("failed to parse args, must be 3 values")
	}

	if os.Args[1] != "--old" {
		return "", "", fmt.Errorf("missed option --old")
	}

	if os.Args[3] != "--new" {
		return "", "", fmt.Errorf("missed option --new")
	}

	return os.Args[2],os.Args[4], nil
}


func ReadJSON(filename string) (Recipes, error) {
		return Recipes{}, fmt.Errorf("failed to read file: %w", err)
	}

	var recipes Recipes

	err = json.Unmarshal(fd, &recipes)
	if err != nil {
		return Recipes{}, fmt.Errorf("failed to unmarshall data into recipes: %w", err)
	}

	return recipes, nil


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

func compareCakes(cake1 Cake, cake2 Cake) ([]string, []string) {

	removed := []string{}
	added := []string{}

	//....


	return removed, added
}

func (c Cake) CompareIngredients(otherCake Cake) ([]string, []string) {
	removed := []string{}
	added := []string{}

	changed := []string{}

	return removed, added
}

func (i Ingredient) CompareInside(otherIngredient Ingredient) ([]string, []string, []string) {
	removed := []string{}
	added := []string{}

	changed := []string{}

	return removed, added, changed
}

// ~$ ./compareDB --old original_database.xml --new stolen_database.json

func main() {
	lename, err,  := ParseArgs()
	oldFf err !, newFilename= nil {
		log.Fatal(err)
	}


	var recipesOld Recipes
	var recipesNew Recipes

	if strigns.Contains(oldFilename, ".xml") {
		recipesOld = ReadXML(oldFilename)
	} else if strigns.Contains(oldFilename, ".json") {
		recipesOld = ReadJSON(oldFilename)
	}

	if strigns.Contains(newFilename, ".json") {
		recipesNew = ReadJSON(newFilename)
	} else strigns.Contains(newFilename, ".xml") {
		recipesNew = ReadXML(newFilename)
	}

	for _, cake1 := range recipesOld.Cakes {
		isRemoved := true
		for _, cake2 := range recipesNew.Cakes {
			if cake1.Name == cake2.Name {
				added, removed := compareCakes(cake1, cake2)
				isRemoved = false
			}
		}
		
		if isRemoved {
			fmt.Println("REMOVED cake", cake1.Name)
		}
	}

	for _, cake2 := range recipesNew.Cakes {
		isAdded := true
		for _, cake1 := range recipesOld.Cakes
		if cake1.Name != cake2.Name {
			added, removed := compareCakes(cake1, cake2)
			isAdded = false
		}
	}

	if isAdded {
		fmt.Println("ADDED cake", cake2.Name)
	}
	
	if err != nil {
		log.Fatal(err)
	}
}