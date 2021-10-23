	package scrabble

import (
 "unicode"
)

func makeScrabbleMap() map[string]int {
	scrabbleMap := make(map[string]int)
    
	oneArray := [...]string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}
    twoArray := [...]string{"D", "G"}
    threeArray := [...]string{"B", "C", "M", "P"}
    fourArray := [...]string{"F", "H", "V", "W", "Y"}
    fiveArray := [...]string{"K"}
    eightrray := [...]string{"J", "X"}
    tenArray := [...]string{"Q", "Z"}
    
    for _, rune:= range oneArray {
        //fmt.Println(rune)
        scrabbleMap[rune] = 1
    }
	for _, rune:= range twoArray {
        scrabbleMap[rune] = 2
    }
	for _, rune:= range threeArray {
        scrabbleMap[rune] = 3
    }
	for _, rune:= range fourArray {
        scrabbleMap[rune] = 4
    }
	for _, rune:= range fiveArray {
        scrabbleMap[rune] = 5
    }
	for _, rune:= range eightrray {
        scrabbleMap[rune] = 8
    }
	for _, rune:= range tenArray {
        scrabbleMap[rune] = 10
    }	
    return scrabbleMap
}

func Score(word string) int {
	scrabbleMap := makeScrabbleMap()

    score := 0
	for _, rune := range word {
        score += scrabbleMap[string(unicode.ToUpper(rune))]
    }
    return score
}