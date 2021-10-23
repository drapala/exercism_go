package cards

import "fmt"

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// a the given index existed in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
    var card int = 0
    var ok bool = true
    
    if ((index >= len(slice)) || (index < 0)){
        // If outside the range of the slice
        card = 0
        ok = false
    } else {
    	// Within range
    	card = slice[index]
        ok = true    
    }
	return card, ok
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range it is be appended.
func SetItem(slice []int, index, value int) []int {
    if _, ok := GetItem(slice, index); ok == false {
		// Item outside range, append to end of slice        
        slice = append(slice, value)
    } else {
    	// Replace existing value
        slice[index] = value
    }
    return slice
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
    if length <= 0 {
        // Return empty slice
        return []int{}
    } else {
    	// Return prefilled slice
    	slice := make([]int, length)        
        for i := range slice {
            slice[i] = value
        }
    	return slice
    }
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if ((index >= len(slice)) || (index < 0)){
        // If outside the range of the slice, return slice
		fmt.Printf("Invalid item: %v\n", index)
        return slice
    } else if index == 0{
    	// To remove first item
        fmt.Printf("First item: %v\n", index)
    	return slice[1:]
    } else if index == (len(slice)-1){
    	// To remove last item
        fmt.Printf("Last item: %v\n", index)
		return slice[0:len(slice)-1]
    } else {
    	// All other valid cases
        fmt.Printf("Valid item: %v\n", index)
        return append(slice[0:(index)], slice[(index+1):]...)
    }	
}
