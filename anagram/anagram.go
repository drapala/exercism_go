package anagram

import (
	"reflect"
	"strings"
)

func getFreq(s string) map[rune]int {
	freq := make(map[rune]int)
	for _, r := range s {
		freq[r]++
	}
	return freq
}

func Detect(subject string, candidates []string) []string {
	// Solution
	var result []string

	// Copy of candidates in lowercase - to be used in comparisons
	candidatesLower := make([]string, len(candidates))

	// Lowercase the subject and make a candidates lowercase copy
	subject = strings.ToLower(subject)
	for i, _ := range candidates {
		candidatesLower[i] = strings.ToLower(candidates[i])
	}
	// Get subject map
	subjectFreq := getFreq(subject)

	// Get candidate map in slice
	candidateFreqArray := make([]map[rune]int, len(candidatesLower))
	for i, candidate := range candidatesLower {
		candidateFreqArray[i] = getFreq(candidate)
	}

	// Compare maps
	// Note: if candidate is exactly same as subject, skip it
	for i, _ := range candidateFreqArray {
		// If candidate is exactly same as subject lowercase, skip it
		if subject == candidatesLower[i]{
			continue
		} else {
			// If Candidate Frequency Map is equal to Subject Frequency Map, add to result 
			if reflect.DeepEqual(subjectFreq, candidateFreqArray[i]){
				// We append the original candidates to preserve the Capitals
				result = append(result, candidates[i])
			}
		}
	}
	return result
}
