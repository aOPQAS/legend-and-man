package main

import (
	"fmt"
)

func compareCakes(oldRecipes, newRecipes Recipes) {
	// ADDED cake
	// ADDED cake "Moonshine Muffin"
	for _, cake2 := range newRecipes.Cakes {
		isAdded := true
		for _, cake1 := range oldRecipes.Cakes {
			if cake1.Name == cake2.Name {
				isAdded = false
				break
			}
		}
		if isAdded {
			fmt.Printf("ADDED cake \"%s\"\n", cake2.Name)
		}
	}

	// REMOVED cake
	// REMOVED cake "Blueberry Muffin Cake"
	for _, cake1 := range oldRecipes.Cakes {
		isRemoved := true
		for _, cake2 := range newRecipes.Cakes {
			if cake1.Name == cake2.Name {
				isRemoved = false
				break
			}
		}
		if isRemoved {
			fmt.Printf("REMOVED cake \"%s\"\n", cake1.Name)
		}
	}

	// CHANGED cooking time
	// CHANGED cooking time for cake "Red Velvet Strawberry Cake" - "45 min" instead of "40 min"
	for _, cake1 := range oldRecipes.Cakes {
		for _, cake2 := range newRecipes.Cakes {
			if cake1.Name == cake2.Name && cake1.Stovetime != cake2.Stovetime {
				fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", cake1.Name, cake2.Stovetime, cake1.Stovetime)
			}
		}
	}

	// Check the ingredients, if the names are the same
	for _, cake1 := range oldRecipes.Cakes {
		for _, cake2 := range newRecipes.Cakes {
			if cake1.Name == cake2.Name {
				oldIngredients := cake1.Ingredients
				newIngredients := cake2.Ingredients
				// ADDED ingredient
				// ADDED ingredient "Coffee beans" for cake  "Red Velvet Strawberry Cake"
				for _, newIng := range newIngredients {
					isAdded := true
					for _, oldIng := range oldIngredients {
						if newIng.Itemname == oldIng.Itemname {
							isAdded = false
							break
						}
					}
					if isAdded {
						fmt.Printf("ADDED ingredient \"%s\" for cake \"%s\"\n", newIng.Itemname, cake2.Name)
					}
				}

				// REMOVED ingredient
				// REMOVED ingredient "Vanilla extract" for cake  "Red Velvet Strawberry Cake"
				for _, oldIng := range oldIngredients {
					isRemoved := true
					for _, newIng := range newIngredients {
						if oldIng.Itemname == newIng.Itemname {
							isRemoved = false
							break
						}
					}
					if isRemoved {
						fmt.Printf("REMOVED ingredient \"%s\" for cake \"%s\"\n", oldIng.Itemname, cake1.Name)
					}
				}

				// CHANGED unit and count
				for _, oldIng := range oldIngredients {
					for _, newIng := range newIngredients {
						if oldIng.Itemname == newIng.Itemname {
							// CHANGED unit for ingredient "Flour" for cake  "Red Velvet Strawberry Cake" - "mugs" instead of "cups"
							if oldIng.Itemcount != newIng.Itemcount {
								fmt.Printf("CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldIng.Itemname, cake1.Name, newIng.Itemcount, oldIng.Itemcount)
							}
							// CHANGED unit count for ingredient "Strawberries" for cake  "Red Velvet Strawberry Cake" - "8" instead of "7"
							if oldIng.Itemunit != newIng.Itemunit {
								if newIng.Itemunit != "" {
									fmt.Printf("CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n", oldIng.Itemname, cake1.Name, newIng.Itemunit, oldIng.Itemunit)
								}
							}
						}
					}
				}

				// REMOVED unit
				// REMOVED unit "pieces" for ingredient "Cinnamon" for cake  "Red Velvet Strawberry Cake"
				for _, oldIng := range oldIngredients {
					for _, newIng := range newIngredients {
						if oldIng.Itemname == newIng.Itemname && oldIng.Itemunit != "" && newIng.Itemunit == "" {
							fmt.Printf("REMOVED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n", oldIng.Itemunit, oldIng.Itemname, cake1.Name)
						}
					}
				}
			}
		}
	}
}
