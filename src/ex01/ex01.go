package ex01

import (
	"JSONtoXML/src/ex00"
	"fmt"
	"sort"
)

// compareCakes compares two cakes from old and new Recipes.
// If a cake has been removed or added, it deletes from recipes because it does not need to be viewed anymore
// It returns list of names of removed and added cakes.
func compareCakes(old *ex00.Recipes, new *ex00.Recipes) ([]string, []string) {
	sort.Slice(
		old.Cakes,
		func(i, j int) bool { return old.Cakes[i].Name < old.Cakes[j].Name },
	)
	sort.Slice(
		new.Cakes,
		func(i, j int) bool { return new.Cakes[i].Name < new.Cakes[j].Name },
	)

	removed, added := make([]string, 0, 10), make([]string, 0, 10)
	for i := 0; i < len(old.Cakes); {
		found := false
		for j := 0; j < len(new.Cakes); j++ {
			if old.Cakes[i].Name == new.Cakes[j].Name {
				found = true
				i++
				break
			}
		}
		if !found {
			removed = append(removed, old.Cakes[i].Name)
			old.Cakes = append(old.Cakes[:i], old.Cakes[i+1:]...)
		}
	}

	for i := 0; i < len(new.Cakes); {
		found := false
		for j := 0; j < len(old.Cakes); j++ {
			if new.Cakes[i].Name == old.Cakes[j].Name {
				found = true
				i++
				break
			}
		}
		if !found {
			added = append(added, new.Cakes[i].Name)
			new.Cakes = append(new.Cakes[:i], new.Cakes[i+1:]...)
		}
	}
	return removed, added
}

// compareCookingTime compares cooking time of old cakes with cooking time of new cakes.
// It returns whether the cooking time equals for old and new cake
func compareCookingTime(old *ex00.Cake, new *ex00.Cake) bool {
	return old.StoveTime == new.StoveTime
}

// compareIngredients compares ingredients list of old cakes with ingredients list of new cakes.
// If an ingredient has been removed or added, it deletes from recipe because it does not need to be viewed anymore.
// It returns list of names of removed and added ingredients of cakes and renewed versions of
// old and new ingredients list
func compareIngredients(
	old []ex00.IngredientItem,
	new []ex00.IngredientItem,
) ([]string, []string, []ex00.IngredientItem, []ex00.IngredientItem) {
	// compareCakes(old, new) func execution guarantees that old and new contains the same cakes
	removed, added := make([]string, 0, 10), make([]string, 0, 10)

	sort.Slice(old, func(i, j int) bool {
		return old[i].ItemName < old[j].ItemName
	})
	sort.Slice(new, func(i, j int) bool {
		return new[i].ItemName < new[j].ItemName
	})

	for i := 0; i < len(old); {
		found := false
		for j := 0; j < len(new); j++ {
			if old[i].ItemName == new[j].ItemName {
				found = true
				i++
				break
			}
		}
		if !found {
			removed = append(removed, old[i].ItemName)
			old = append(old[:i], old[i+1:]...)
		}
	}
	for i := 0; i < len(new); {
		found := false
		for j := 0; j < len(old); j++ {
			if new[i].ItemName == old[j].ItemName {
				found = true
				i++
				break
			}
		}
		if !found {
			added = append(added, new[i].ItemName)
			new = append(new[:i], new[i+1:]...)
		}
	}
	return removed, added, old, new
}

func compareIngredientCount(
	old *ex00.IngredientItem,
	new *ex00.IngredientItem,
) bool {
	return old.ItemCount == new.ItemCount
}

// It returns false if units are not equal and none of them equals an empty string. Otherwise true
func compareIngredientUnit(
	old *ex00.IngredientItem,
	new *ex00.IngredientItem,
) bool {
	return !(old.ItemUnit != new.ItemUnit && old.ItemUnit != "" && new.ItemUnit != "")
}

func isIngredientUnitMissed(
	old *ex00.IngredientItem,
	new *ex00.IngredientItem,
) bool {
	return new.ItemUnit == "" && old.ItemUnit != ""
}

func isIngredientUnitAdded(
	old *ex00.IngredientItem,
	new *ex00.IngredientItem,
) bool {
	return old.ItemUnit == "" && new.ItemUnit != ""
}

// CompareDB examines the new database for changes compared to the old version.
// CompareDB accepts both .json and .xml file paths as an old and a new path
func CompareDB(old *string, new *string) {
	_, oldRecipes := ex00.ParseDB(old)
	_, newRecipes := ex00.ParseDB(new)

	removed, added := compareCakes(oldRecipes, newRecipes)
	for i := 0; i < len(added); i++ {
		fmt.Printf("ADDED cake \"%s\"\n", added[i])
	}
	for i := 0; i < len(removed); i++ {
		fmt.Printf("REMOVED cake \"%s\"\n", removed[i])
	}

	for i := 0; i < len(oldRecipes.Cakes); i++ {
		if !compareCookingTime(&oldRecipes.Cakes[i], &newRecipes.Cakes[i]) {
			fmt.Printf(
				"CHANGED cooking time for cake \"%s\" - \"%s\" instead of \"%s\"\n",
				oldRecipes.Cakes[i].Name,
				newRecipes.Cakes[i].StoveTime,
				oldRecipes.Cakes[i].StoveTime,
			)
		}
		removed, added, oldRecipes.Cakes[i].Ingredients, newRecipes.Cakes[i].Ingredients = compareIngredients(
			oldRecipes.Cakes[i].Ingredients,
			newRecipes.Cakes[i].Ingredients,
		)
		for j := 0; j < len(added); j++ {
			fmt.Printf(
				"ADDED ingredient \"%s\" for cake \"%s\"\n",
				added[j],
				oldRecipes.Cakes[i].Name,
			)
		}
		for j := 0; j < len(removed); j++ {
			fmt.Printf(
				"REMOVED ingredient \"%s\" for cake \"%s\"\n",
				removed[j],
				oldRecipes.Cakes[i].Name,
			)
		}
		for j := 0; j < len(oldRecipes.Cakes[i].Ingredients); j++ {
			if !compareIngredientUnit(
				&oldRecipes.Cakes[i].Ingredients[j],
				&newRecipes.Cakes[i].Ingredients[j],
			) {
				fmt.Printf(
					"CHANGED unit for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n",
					oldRecipes.Cakes[i].Ingredients[j].ItemName,
					oldRecipes.Cakes[i].Name,
					newRecipes.Cakes[i].Ingredients[j].ItemUnit,
					oldRecipes.Cakes[i].Ingredients[j].ItemUnit,
				)
			}
			if !compareIngredientCount(
				&oldRecipes.Cakes[i].Ingredients[j],
				&newRecipes.Cakes[i].Ingredients[j],
			) {
				fmt.Printf(
					"CHANGED unit count for ingredient \"%s\" for cake \"%s\" - \"%s\" instead of \"%s\"\n",
					oldRecipes.Cakes[i].Ingredients[j].ItemName,
					oldRecipes.Cakes[i].Name,
					newRecipes.Cakes[i].Ingredients[j].ItemCount,
					oldRecipes.Cakes[i].Ingredients[j].ItemCount,
				)
			}
			if isIngredientUnitAdded(
				&oldRecipes.Cakes[i].Ingredients[j],
				&newRecipes.Cakes[i].Ingredients[j],
			) {
				fmt.Printf(
					"ADDED unit \"%s\" for ingredient \"%s\" for cake \"%s\"\n",
					newRecipes.Cakes[i].Ingredients[j].ItemUnit,
					oldRecipes.Cakes[i].Ingredients[j].ItemName,
					oldRecipes.Cakes[i].Name,
				)
			}
			if isIngredientUnitMissed(
				&oldRecipes.Cakes[i].Ingredients[j],
				&newRecipes.Cakes[i].Ingredients[j],
			) {
				fmt.Printf(
					"REMOVED unit \"%s\" from ingredient \"%s\" for cake \"%s\"\n",
					oldRecipes.Cakes[i].Ingredients[j].ItemUnit,
					oldRecipes.Cakes[i].Ingredients[j].ItemName,
					oldRecipes.Cakes[i].Name,
				)
			}
		}
	}
}
