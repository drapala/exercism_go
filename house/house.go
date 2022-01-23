package house

import "fmt"

var nounToActionMap = map[string]string{
	"house that Jack built.":           "",
	"malt":                             "lay in the",
	"rat":                              "ate the",
	"cat":                              "killed the",
	"dog":                              "worried the",
	"cow with the crumpled horn":       "tossed the",
	"maiden all forlorn":               "milked the",
	"man all tattered and torn":        "kissed the",
	"priest all shaven and shorn":      "married the",
	"rooster that crowed in the morn":  "woke the",
	"farmer sowing his corn":           "kept the",
	"horse and the hound and the horn": "belonged to the",
}
var sequence = []string{"house that Jack built.", "malt", "rat", "cat", "dog", "cow with the crumpled horn", "maiden all forlorn", "man all tattered and torn", "priest all shaven and shorn", "rooster that crowed in the morn", "farmer sowing his corn", "horse and the hound and the horn"}

func Verse(v int) string {
	stack := sequence[:v]
	var result string
	for i := v - 1; i >= 0; i-- {
		if i == v-1 {
			result += fmt.Sprintf("This is the %s", stack[i])
		} else {
			result += fmt.Sprintf("\nthat %s %s", nounToActionMap[stack[i+1]], stack[i])
		}
	}
	return result
}

func Song() string {
	var result string
	for i := 1; i <= len(sequence); i++ {
		result += Verse(i)
		if i != len(sequence) {
			result += "\n\n"
		}
	}
	return result
}
