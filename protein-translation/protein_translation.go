package protein

import "fmt"

var ErrStop = fmt.Errorf("stop codon found")
var ErrInvalidBase = fmt.Errorf("base not recognized")

func FromCodon(codon string) (string, error) {
	switch codon {
	case "AUG":
		return "Methionine", nil
	case "UUU", "UUC":
		return "Phenylalanine", nil
	case "UUA", "UUG":
		return "Leucine", nil
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", nil
	case "UAU", "UAC":
		return "Tyrosine", nil
	case "UGU", "UGC":
		return "Cysteine", nil
	case "UGG":
		return "Tryptophan", nil
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

func FromRNA(rna string) ([]string, error) {
	var result = make([]string, 0)

	// Error handling
	if len(rna)%3 != 0 {
		return result, ErrInvalidBase
	}
	var protein string
	var err error
	// Go through RNA string in 3-character chunks
	for i := 0; i < len(rna); i += 3{
		// Encode in chunks of 3
		protein, err = FromCodon(rna[i:i+3])

		// Error handling
		if err != nil {
			if err == ErrStop {
				// Stop code
				return result, nil
			} else if err == ErrInvalidBase {
				//Invalid codon
				return nil, err
			}
		}
		// Append protein to result
		result = append(result, protein)
	}
	return result, err
}