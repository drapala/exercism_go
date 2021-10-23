package luhn

import (
	"math"
	"strconv"
	"strings"
)

func Valid(number string) bool{
    // Trim whitespaces
	clean := strings.ReplaceAll(number, " ", "")
    
    // Print number
    //fmt.Printf("Number: %s | length: %d \n", clean, len(clean))
    
	// Check for non-digit characters
    if _, err := strconv.Atoi(clean); err == nil {
        //  Number is valid
    } else {
    	// Not a valid number, some weird characters must exist
        //fmt.Printf("%s \n", err)
        //fmt.Printf("================= \n")
        return false
    }

    // If length is 1 or less
    if len(clean) <= 1 {
        //fmt.Printf("Length is <= 1; exit \n")
        //fmt.Printf("================= \n")
        return false
    }

    // Now we have a number with no whitespace, weird characters, and is 1 or more in length
	// Our return sum
    sum := 0

    
	// Loop from end
    for i:= len(clean) - 1; i >= 0; i-- {
        r := string(clean[i]) // Number at position i in string
        r_num, _ := strconv.Atoi(r) // Numerical value of r
        pos := len(clean) - i // Position for Luhn
    	//fmt.Printf("position: %d | value: %s \n", pos, r)

		if math.Mod(float64(pos), 2) == 0 {
			// For every second digit - double it
			var newdigit int = r_num * 2
            // If digit is greater than 9, subtract 9
            if newdigit >= 9 {
                newdigit -= 9
            }
    		
            //fmt.Printf("newdigit: %d \n", newdigit)    	
        	
            // Then add it to sum
        	sum += newdigit
        } else {
        	// Else just add it to sum
            sum += r_num
        }
    }

    //fmt.Printf("sum: %d \n", sum)
    //fmt.Printf("================= \n")
    
    if math.Mod(float64(sum), 10) == 0 {
        return true
    } else {
    	return false
    }
    
    // For all other cases
    return false
}