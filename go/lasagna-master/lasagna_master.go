package lasagnamaster

import "fmt"

func PreparationTime(layers []string, perLayerPreparationTime int) int {
	if perLayerPreparationTime == 0 {
		perLayerPreparationTime = 2
	}

	totalPreparationTime := perLayerPreparationTime * len(layers)

	return totalPreparationTime
}

func Quantities(layers []string) (int, float64) {
	noodlesCount := 0
	sauceCount := 0.0

	for _, layer := range layers {
		switch layer {
		case "noodles":
			noodlesCount += 50
		case "sauce":
			sauceCount += 0.2
		}
	}

	return noodlesCount, sauceCount
}

func AddSecretIngredient(friendRecipeList []string, ownRecipeList []string) {
	if len(friendRecipeList) == 0 || len(ownRecipeList) == 0 {
		fmt.Errorf("slice is empty")
	}

	lastFriendValue := friendRecipeList[len(friendRecipeList)-1]
	ownRecipeList[len(ownRecipeList)-1] = lastFriendValue
}

func ScaleRecipe(quantities []float64, numberOfPortions int) []float64 {
	factor := float64(numberOfPortions) / 2.0

	scaled := make([]float64, len(quantities))

	for i := range len(quantities) {
		scaled[i] = quantities[i] * factor
	}

	return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
