package main

import (
	"fmt"
)

func main() {
	ingredients := []float64{}
	for {
		var n float64
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}
		ingredients = append(ingredients, n)
	}

	fmt.Println(maxCalderone(ingredients))
}


func maxCalderone(ingredients []float64) float64 {
	if len(ingredients) == 1 {
		return ingredients[0];
	}

	max := 0.;

	for i := 0; i < len(ingredients); i++ {
		for j := i + 1; j < len(ingredients); j++ {
            copyInts := make([]float64, len(ingredients))
			copy(copyInts, ingredients)
			copyInts = append(copyInts[:i], append(copyInts[i+1:j], copyInts[j+1:]...)...)
			copyInts = append(copyInts, float64(ingredients[i] + ingredients[j]) / 2)
			result := maxCalderone(copyInts)
			if (result > max) {
				max = result
			}
		}
	}
	return max
}