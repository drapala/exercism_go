package main

import "fmt"

func main() {
	digit := 1024
	ones := digit % 10
	tens := digit / 10 % 10
	hundreds := digit / 100 % 10
	thousands := digit / 1000 % 10

	fmt.Println(ones, " ", tens, " ", hundreds, " ", thousands)
}
