package piglatin

import (
	"regexp"
	"strings"
)

func VowelSounds(sentence string) bool {
	var startsWithVowel bool
	// Catches all vowelsounds
	results := regexp.MustCompile("a|e|i|o|u|xr|yt").FindAllStringIndex(sentence, -1)
	
	if len(results) > 0 { // If non-empty
		if results[0][0] == 0 { // If first letter is vowel
			startsWithVowel = true
		}
	}
	return startsWithVowel
}

func ConsonantClusterIndex(sentence string) int {
	results := regexp.MustCompile("a|e|i|o|u|xr|yt").FindAllStringIndex(sentence, -1)
	if len(results) > 0 { // If non-empty
		return results[0][0] // Index of first vowel 
	}
	return -1
}

func QUIndex(sentence string) int {
	// Since we process this after vowels, no need to deal with it
	results := regexp.MustCompile("qu").FindAllStringIndex(sentence, -1)
	if len(results) > 0 { // If non-empty
		return results[0][1] // Index of qu ending
	}
	return -1
}

// Operates per word
func PigConverter(word string) string {
	// 1. Figure out how to grab leading "vowelsounds" at beginning
	startsWithVowel := VowelSounds(word) 
	if startsWithVowel { // Starts with vowelsound
		return word + "ay"
	} 
	// From here, all starts with consonants
	// ------------------------------------------------------------	
	// 2. Figure out how to grab leading "consonant clusters" at beginning terminating in:
	//   - "qu"
	qu_index := QUIndex(word)

	if qu_index != -1 { // If qu found
		return word[qu_index:] + word[0:qu_index] + "ay"
	}
	//   - "y" second letter of 2 letter word
	if len(word) == 2 && word[1] == 'y' {
		return word[1:] + word[0:1] + "ay"
	}
	//   - an actual vowel
	cluster_index := ConsonantClusterIndex(word)
	if cluster_index != -1 { // If non-empty
		return word[cluster_index:] + word[:cluster_index] + "ay"
	}
	return "" // stub
}

// Splits a sentence on whitespace and returns array
func SplitStringOnWhitespace(sentence string) []string {
	return regexp.MustCompile(" ").Split(sentence, -1)
}

func Sentence(sentence string) string {
	var result string
	// Loop over words and convert
	for _, word := range SplitStringOnWhitespace(sentence) {
		result += PigConverter(word) + " "
	}
	return strings.TrimSpace(result) // Trim any extra whitespace added above
}
