package lasagna

//import "fmt"

func PreparationTime(layers []string, time int) int {
    if time == 0 {
    	return len(layers) * 2
    } else {
    	return len(layers) * time
    }
}

func Quantities(layers []string) (int, float64) {
	noodles := 0
    sauce := 0.0

	for _, layer := range layers {
        if layer == "noodles" {
            noodles += 50
        } else if layer == "sauce" {
            sauce += 0.2
        }
    }

	return noodles, sauce
}

func AddSecretIngredient(friendsList []string, myList []string) ([]string) {
	//fmt.Println(friendsList)
    //fmt.Println(myList)
    //fmt.Println(friendsList[len(friendsList) - 1])
    
	return append(myList, friendsList[len(friendsList) - 1])
}

func ScaleRecipe(quantities []float64, numPortions int) ([]float64) {
	reqd := numPortions - 2 // 4
	// Copy args
    clone := make([]float64, len(quantities))
    _ = copy(clone, quantities)
    
    // Already have quantities for 2 portions
    if reqd == 0 {
        return quantities
    } else {
    	for i, _ := range quantities {
            clone[i] += (quantities[i] * float64(reqd)) / 2.0
        }
    }
    return clone
}