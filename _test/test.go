// Go program to illustrate how to copy
// a slice into another slice using the
// copy function
package main

import "fmt"

func main() {
	slice_1 := []string{"Geeks", "for", "Geeks", "GFG"}
	slice_2 := make([]string, 3)

	// Before Copying
	fmt.Println("Slice_1: ", slice_1)
	fmt.Println("Slice_2: ", slice_2)

	// Copying slice_1 into slice_2
	// using copy function
	Copy_1 := copy(slice_2, slice_1)
	fmt.Println("\nSlice_1: ", slice_1)
	fmt.Println("Slice_2: ", slice_2)
	fmt.Println("Number of elements copied: ", Copy_1)

	// Copying the slice
	// using copy function
	// see the code clearly
	Copy_2 := copy(slice_1, []string{"123geeks", "gfg"})
	fmt.Println("\nSlice_1 : ", slice_1)
	fmt.Println("Number of elements copied:", Copy_2)

}
