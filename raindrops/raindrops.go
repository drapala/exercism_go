package raindrops

import (
    "strconv"
    "fmt"
)

func Convert(x int) string {    
    fmt.Printf("x: %d \n", x)
	fmt.Printf("x/3: %d \n", x%3)

    s := ""

    if x%3 == 0{
        s += "Pling"
    }
	if x%5 == 0{
        s += "Plang"
    }
	if x%7 == 0{
        s += "Plong"
    }

	if s == "" {
        s += strconv.Itoa(x)
    }  
    return s
}