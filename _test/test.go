package main

import "fmt"

func main() {
	str1 := "Volkswagen Beetle"
    str2 := "Volkswagen Golf"

    fmt.Printf("%s < %s = %v\n", str1, str2, str1 < str2)
}