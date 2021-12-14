package dna

import "fmt"

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	// Init
	var h Histogram = make(map[rune]int)
	h['A'] = 0
	h['C'] = 0
	h['G'] = 0
	h['T'] = 0

	for _, value := range(d) {
		// If value is none of these, return error
		if value != 'A' && value != 'C' && value != 'G' && value != 'T' {
			return h, fmt.Errorf("not a valid DNA") 
		}
		h[value] += 1	
	}
	return h, nil
}
