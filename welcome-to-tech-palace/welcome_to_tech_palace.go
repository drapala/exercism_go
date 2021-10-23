package techpalace

import (
    "fmt"
    "strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
    //panic("Please implement the WelcomeMessage() function")
}

func ReturnMultiple (char string, num int) string {
	s := ""
    for i := 1; i <= num; i++ {
        s += char
    }
	return s
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	// ("*" * numStarsPerLine) + "\n" + welcomeMsg + "\n" + ("*" * numStarsPerLine)
	border := ReturnMultiple("*", numStarsPerLine)
	s :=  border + "\n" + welcomeMsg + "\n" + border
    fmt.Println(s)
    return s
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	fmt.Println("Old message:")
    fmt.Println(oldMsg)
    fmt.Println("New message:")
    s := strings.ReplaceAll(oldMsg, "*", "")
	fmt.Println(s)
	s = strings.TrimSpace(s)
    fmt.Println(s)
    
    return s
    //panic("Please implement the CleanupMessage() function")
}
