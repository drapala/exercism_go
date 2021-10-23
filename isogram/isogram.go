package isogram

import (
    "fmt"
    "strings"
)

func IsIsogram (word string) bool {
    word = strings.ToLower(word)// Convert to lowercase
    fmt.Println(word)
    //fmt.Printf("a: %d \n", int('a')) // 97
    //fmt.Printf("z: %d \n", int('z')) // 122
    
    var a[int('z')-int('a')+1]bool // 26 element array for flags

    // Iterate over string
    for _, char := range word {
    	//fmt.Printf("character %c starts at byte position %d\n", char, pos)
		
        // Shifted to fit array index
        i := int(char) - int ('a')

        // If character between a to z:
        if i >= 0 && i <= (len(a) - 1) {
            // If character was already found, return
            if a[i] == true {
                return false
            } else {
            	a[i] = true
            }
        }
        // Else, doesn't matter
	}

    // Else an isogram
    return true
}