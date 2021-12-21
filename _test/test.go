package main

import "fmt"

func main() {
	var myList = make([]int, 2)
	myList[0] = 1
	myList[1] = 2
	fmt.Println(myList)
	myList[2] = 3
}
