package blackjack
	
// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	// Create a map
	deck := make(map[string]int)

    // Populate values
    deck["ace"] = 11
    deck["eight"] = 8
    deck["two"] = 2
    deck["nine"] = 9
    deck["three"] = 3
    deck["ten"] = 10
    deck["four"] = 4
    deck["jack"] = 10
    deck["five"] = 5
    deck["queen"] = 10
    deck["six"] = 6
    deck["king"] = 10
    deck["seven"] = 7
    deck["other"] = 0

	return deck[card]
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	var flag bool;
    switch ParseCard(card1) + ParseCard(card2) {
    case 21:
		flag = true
    default:
    	flag = false
    }
    return flag
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	switch isBlackjack{
        case true:
			switch dealerScore{
                case 11, 10:
            		return "S"
				// Dealer doesn't have ace, figure (Jack, Queen, King) or a ten --> win
                default:
            		return "W"
            }
        case false:
    		switch dealerScore{
				case 11:
            		return "P"
                default:
            		return "S"
            }		
    }
	return "S"
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	switch {
        case handScore >= 17:
    		return "S"
        case handScore <= 11:
    		return "H"
        case (handScore <= 16 && handScore >= 12):
    		switch {
                case dealerScore >= 7:
            		return "H"
                default: 
            		return "S"
            }
    }

    return "S"
}
