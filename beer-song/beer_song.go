package beer

import (
	"errors"
	"fmt"
)

func Song() string {
	verses, _ := Verses(99, 0)
	return verses
}

func Verses(start, stop int) (string, error) {
	if start <= stop {
		return "", errors.New("start less than stop")
	}

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
	var bottle string
	if n == 1 {
		bottle = "bottle"
	} else {
		bottle = "bottles"
	}
	return fmt.Sprintf("%d %s of beer on the wall, %d %s of beer.\n", n, bottle, n, bottle)
}

func nTakedown(n int) string {
	var bottle string
	if n == 2 {
		bottle = "bottle"
	} else {
		bottle = "bottles"
	}
	var oneorit string
	// Take ____ down and pass it around
	if n == 1 {
		oneorit = "it"
	} else {
		oneorit = "one"
	}

	var numleft string
	// ...pass it around, ____ bottles of beer....
	if n == 1 {
		numleft = "no more"
	} else {
		numleft = fmt.Sprintf("%d", n-1)
	}
	return fmt.Sprintf("Take %s down and pass it around, %s %s of beer on the wall.\n", oneorit, numleft, bottle)
}

func Verse(n int) (string, error) {
	if n == 0 {
		return "No more bottles of beer on the wall, no more bottles of beer.\nGo to the store and buy some more, 99 bottles of beer on the wall.\n", nil
	} else if n >= 1 && n <= 99 {
		return (nBottles(n) + nTakedown(n)), nil
	}
	return "", errors.New("n out of bounds")
}