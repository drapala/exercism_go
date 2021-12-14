package strand

func ToRNA(dna string) string {
	rna := ""

	switch len(dna) {
	case 0:
		return rna
	default:
		for _, value := range dna {
			switch value {
			case 'G':
				rna += string('C')
			case 'C':
				rna += string('G')
			case 'T':
				rna += string('A')
			case 'A':
				rna += string('U')
			}
		}
		return rna
	}
}
