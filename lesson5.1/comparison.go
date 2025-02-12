package main

import (
	"fmt"
	"log"
)

func compareCakes(cake1 Cake, cake2 Cake) ([]string, []string) {
	removed := []string{}
	added := []string{}

	// delete cake if name does not match
	// ADDED cake "Moonshine Muffin"
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

	// REMOVED cake "Blueberry Muffin Cake"
	for _, cake2 := range recipesNew.Cakes {
		isAdded := true
		for _, cake1 := range recipesOld.Cakes {
			if cake1.Name != cake2.Name {
				added, removed := compareCakes(cake1, cake2)
				isAdded = false
		}
	}

	if isAdded {
		fmt.Println("ADDED cake", cake2.Name)
	}

	// Checking the change in cooking time 
	// CHANGED cooking time for cake "Red Velvet Strawberry Cake" - "45 min" instead of "40 min"
	if cake1.Stovetime != cake2.Stovetime {
		fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", cake1.Name, cake2.Stovetime, cake1.Stovetime)
	}

	oldIngredients := cake1.Ingredients
	newIngredients := cake2.Ingredients

	// Search for deleted ingredients
	// REMOVED ingredient "Vanilla extract" for cake  "Red Velvet Strawberry Cake"
	for _, oldIng := range oldIngredients {
		isRemoved := true
		for _, newIng := range newIngredients {
			if oldIng.Itemname == newIng.Itemname {
				isRemoved = false
			}
		}
		if isRemoved {
			fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", oldIng.Itemname, cake1.Name)
		}
	}

	// Search for added ingredients
	// ADDED ingredient "Coffee beans" for cake  "Red Velvet Strawberry Cake"
	for _, newIng := range newIngredients {
		isAdded := true
		for _, oldIng := range oldIngredients {
			if newIng.Itemname == oldIng.Itemname {

				isAdded = false
			}
		}
		if isAdded {
			fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", newIng.Itemname, cake2.Name)
		}
	}

	/*
	needs to be completed:
	CHANGED unit count for ingredient "Strawberries" for cake  "Red Velvet Strawberry Cake" - "8" instead of "7"
	REMOVED unit "pieces" for ingredient "Cinnamon" for cake  "Red Velvet Strawberry Cake"
	*/


	// CHANGED unit for ingredient "Flour" for cake  "Red Velvet Strawberry Cake" - "mugs" instead of "cups"

	if cake1.Itemunit != cake2.Itemunit {
		fmt.Printf("CHANGED ingredient unit for cake \"%s\" - \"%s\" instead of \"%s\"\n", cake1.Name, cake2.Itemunit, cake1.Itemunit)
	}
	
	// CHANGED unit count for ingredient "Strawberries" for cake  "Red Velvet Strawberry Cake" - "8" instead of "7"

	if cake1.Itemcount != cake2.Itemcount {
		fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n",cake1.Itemname, cakeName, cake2.Itemcount, cake1.Itemcount)
	}
	
	// Search for deleted ingredients (removed a unit of pieces)
	// REMOVED unit "pieces" for ingredient "Cinnamon" for cake  "Red Velvet Strawberry Cake"



	
	// code...



	if err != nil {
		log.Fatal(err)
	}

	return removed, added

}