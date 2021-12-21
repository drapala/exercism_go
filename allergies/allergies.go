package allergies

import "fmt"

var allergyList = []struct {
	food string
	score uint
}{
	{
		food: "eggs",
		score: 1,

	},
	{
		food: "peanuts",
		score: 2,
	},
	{
		food: "shellfish",
		score: 4,
	},
	{
		food: "strawberries",
		score: 8,
	},
	{
		food: "tomatoes",
		score: 16,
	},
	{
		food: "chocolate",
		score: 32,
	},
	{
		food: "pollen",
		score: 64,
	},
	{
		food: "cats",
		score: 128,
	},
}

func contains(food string, foods []string) bool {
	for _, f := range foods {
		if f == food {
			return true
		}
	}
	return false
}

func getFood(x uint, foods *[]string) error {
	// If score is zero, nothing to add
	if x == 0 {
		return nil
	}

	// Error handle the case where foods[0] > x
	if allergyList[0].score > x {
		return fmt.Errorf("this score cannot be met with given allergens")
	}

	var newscore uint

	// Get closest food to x
	for i, _ := range allergyList {
		if i == len(allergyList) - 1 {
			// If the last food in the list is less than score
			if allergyList[i].score <= x {
				// If foods slice doesn't contain the food, add it in
				if !contains(allergyList[i].food, *foods) {
					*foods = append(*foods, allergyList[i].food)
				}
				newscore = x - allergyList[i].score
				break
			}
		} else {
			if allergyList[i+1].score > x {
				// If next one is above score, add current
				if !contains(allergyList[i].food, *foods) {
					*foods = append(*foods, allergyList[i].food)
				}
				newscore = x - allergyList[i].score
				break
			}
		}
	}

	// Call recursively
	err := getFood(newscore, foods)

	return err
}

func Allergies(allergies uint) []string {
	// Get list of allergies
	var myslice []string
	err := getFood(allergies, &myslice)
	if err != nil {
		fmt.Println(err)
	}
	return myslice
}

func AllergicTo(allergies uint, allergen string) bool {
	// Return whether the allergen is in the list of allergies
	return contains(allergen, Allergies(allergies))
}
