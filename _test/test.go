package main

import (
	"fmt"
)

type verse struct {
	current_animal   string
	comment          string
	previous_animals []string
}

var verse_1 = verse{
	current_animal:   "fly",
	comment:          "",
	previous_animals: []string{},
}
var verse_2 = verse{
	current_animal:   "spider",
	comment:          "It wriggled and jiggled and tickled inside her.",
	previous_animals: []string{"spider","fly"},
}
var verse_3 = verse{
	current_animal:   "bird",
	comment:          "How absurd to swallow a bird!",
	previous_animals: []string{"bird","spider", "fly"},
}
var verse_4 = verse{
	current_animal:   "cat",
	comment:          "Imagine that, to swallow a cat!",
	previous_animals: []string{"cat", "bird", "spider", "fly"},
}
var verse_5 = verse{
	current_animal:   "dog",
	comment:          "What a hog, to swallow a dog!",
	previous_animals: []string{"dog", "cat", "bird", "spider", "fly"},
}
var verse_6 = verse{
	current_animal:   "goat",
	comment:          "Just opened her throat and swallowed a goat!",
	previous_animals: []string{"goat", "dog", "cat", "bird", "spider", "fly"},
}
var verse_7 = verse{
	current_animal:   "cow",
	comment:          "I don't know how she swallowed a cow!",
	previous_animals: []string{"cow", "goat", "dog", "cat", "bird", "spider", "fly"},
}
var verse_8 = verse{
	current_animal:   "horse",
	comment:          "She's dead, of course!",
	previous_animals: []string{},
}

var verse_map map[int]verse = map[int]verse{
	1: verse_1,
	2: verse_2,
	3: verse_3,
	4: verse_4,
	5: verse_5,
	6: verse_6,
	7: verse_7,
	8: verse_8,
}

func NarratePreviousAnimals(verse_n verse) string {
	if len(verse_n.previous_animals) == 0 { // no looping logic
		return ""
	}
	var result string

	for i := 1; i < len(verse_n.previous_animals); i++ {
		if verse_n.previous_animals[i] == "spider" {
			result += fmt.Sprintf("\nShe swallowed the %s to catch the %s that wriggled and jiggled and tickled inside her.", verse_n.previous_animals[i-1], verse_n.previous_animals[i])
		} else {
			result += fmt.Sprintf("\nShe swallowed the %s to catch the %s.", verse_n.previous_animals[i-1], verse_n.previous_animals[i])
		}
	}
	return result + "\n"
}

func Verse(v int) string {
	var result string

	verse_n := verse_map[v]

	// Add first line and comment to all
	result += fmt.Sprintf("I know an old lady who swallowed a %s.", verse_n.current_animal)
	result += fmt.Sprintf("\n%s", verse_n.comment)

	// Add previous animal narrations
	result += NarratePreviousAnimals(verse_n)

	// Add static comment if not last
	if v!=8 {
		result += "I don't know why she swallowed the fly. Perhaps she'll die."
	}

	return result
}

func main() {
	fmt.Println(Verse(1))
	fmt.Println("\n")
	fmt.Println(Verse(2))
	fmt.Println("\n")
	fmt.Println(Verse(3))
	fmt.Println("\n")
	fmt.Println(Verse(4))
	fmt.Println("\n")
	fmt.Println(Verse(5))
	fmt.Println("\n")
	fmt.Println(Verse(6))
	fmt.Println("\n")
	fmt.Println(Verse(7))
	fmt.Println("\n")
	fmt.Println(Verse(8))
	fmt.Println("\n")
}
