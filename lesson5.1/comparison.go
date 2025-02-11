package main

import (
	"fmt"
	"log"
)

func compareCakes(cake1 Cake, cake2 Cake) ([]string, []string) {
	removed := []string{}
	added := []string{}

	// удаляем торт если название не совпадает
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

	// Проверяем изменение времени приготовления
	// CHANGED cooking time for cake "Red Velvet Strawberry Cake" - "45 min" instead of "40 min"
	if cake1.Stovetime != cake2.Stovetime {
		fmt.Printf("CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n", cake1.Name, cake2.Stovetime, cake1.Stovetime)
	}

	oldIngredients := cake1.Ingredients
	newIngredients := cake2.Ingredients

	// Поиск удаленных ингредиентов
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

	// Поиск добавленных ингредиентов
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

	//
	// REMOVED unit "pieces" for ingredient "Cinnamon" for cake  "Red Velvet Strawberry Cake"
	if err != nil {
		log.Fatal(err)
	}

	return removed, added

}