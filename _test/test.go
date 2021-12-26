package main

import (
	"errors"
	"fmt"
)

func Verses(start, stop int) (string, error) {
	var result string

	for n := start; n >= stop; n-- {
		verse, err := Verse(n)
		if err != nil {
			return "", err
		}
		result += verse + "\n"
	}

	return result, nil
}

func nBottles(n int) string {
	return fmt.Sprintf("%d bottles of beer on the wall, %d bottles of beer.\n", n, n)
}

func Verse(n int) (string, error) {
	if n == 0 {
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	} else if n > 1 && n <= 99 {
		return (nBottles(n) + fmt.Sprintf("Take one down and pass it around, %d bottles of beer on the wall.\n",n-1)), nil
	} else if n == 1 {
		return fmt.Sprintf("%d bottle of beer on the wall, %d bottle of beer.\nTake it down and pass it around, no more bottles of beer on the wall.\n", n, n), nil
	}
	return "", errors.New("n out of bounds")
}

func main() {
	fmt.Println(Verse(8))
	fmt.Println(Verse(3))
	fmt.Println(Verse(2))
	fmt.Println(Verse(1))
	fmt.Println(Verse(0))

	fmt.Println(Verses(8, 6))
}
