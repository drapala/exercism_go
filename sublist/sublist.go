package sublist

type Relation = string

func firstBgMatchIndex(sm, bg []int) int {
	for i := 0; i < len(bg); i++ {
		if bg[i] == sm[0] {
			return i
		}
	}
	return -1
}

func compareFromIndex(sm, bg []int, start int) bool {
	if start+len(sm) > len(bg) {
		return false // bg will go out of bound
	}
	for i := start; i < len(sm); i++ {
		if sm[i-start] != bg[i] {
			return false // Sequentially did not match
		}
	}
	return true // Sequentially matched
}

func Sublist(l1, l2 []int) string {
	var sm, bg []int
	// Handle empty cases
	if len(l1) == 0 && len(l2) == 0 {
		return "equal"
	} else if len(l1) == 0 && len(l2) > 0 {
		return "sublist"
	} else if len(l1) > 0 && len(l2) == 0 {
		return "superlist"
	}
	// Figure out smaller and larger list
	if len(l1) >= len(l2) {
		sm = l2
		bg = l1
	} else {
		sm = l1
		bg = l2
	}
	// Find first match
	i := firstBgMatchIndex(sm, bg)
	if i == -1 { // No match found
		return "unequal"
	}
	// Find out if match is sublist
	if !compareFromIndex(sm, bg, i) { // Not a sublist
		return "unequal"
	} else {
		if len(sm) == len(bg) { // Is equal
			return "equal"
		}
		if len(l1) >= len(l2) { // l1 is superlist
			return "superlist"
		} else {
			return "sublist" // l1 is sublist
		}
	}
}
