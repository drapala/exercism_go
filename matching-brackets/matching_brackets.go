package brackets

func contains(s []rune, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func BracketSequenceCorrect(input string) bool {
	var stack []string
	open := []rune{'[', '(', '{'}
	close := []rune{']', ')', '}'}

	for _, c := range input {
		if contains(open, c) { // Open Bracket
			stack = append(stack, string(c)) // Push to stack
		} else if contains(close, c) { // Close Bracket
			if len(stack) == 0 { // Hit closing Bracket with empty stack
				return false
			}
			peek := stack[len(stack)-1] // Peek 
			switch c{
			case ']':
				if peek != "[" {
					return false
				} else {
					stack = stack[:len(stack)-1] // Pop
				}
			case ')':
				if peek != "(" {
					return false
				} else {
					stack = stack[:len(stack)-1] // Pop
				}
			case '}':
				if peek != "{" {
					return false
				} else {
					stack = stack[:len(stack)-1] // Pop
				}
			}
		}
	}

	return len(stack) == 0
}

func Bracket(input string) bool {
	var cleaned string
	match := []rune{'[', ']', '(', ')', '{', '}'}
	for _, c := range input {
		if contains(match, c) {
			cleaned += string(c)
		}
	}
	return BracketSequenceCorrect(cleaned)
}
