package proverb

import (
	"fmt"
)

func ProverbMaker(first, second string) string {
	if second == "" {
		return fmt.Sprintf("And all for the want of a %s.", first)
	} else {
		return fmt.Sprintf("For want of a %s the %s was lost.", first, second)
	}
}

func Proverb(rhyme []string) []string {
	// Return array
	proverb := []string{}

	switch len(rhyme){
	case 0:
		// For empty string return empty
		return proverb
	case 1: 
		// For single item, return just the last line
		proverb = append(proverb, ProverbMaker(rhyme[0], ""))
		return proverb
	default:
		// Default scenario
		for i := 0; i < len(rhyme) - 1; i++ {
			proverb = append(proverb, ProverbMaker(rhyme[i], rhyme[i+1]))
		}
		proverb = append(proverb, ProverbMaker(rhyme[0], ""))		
		return proverb
	}
}
