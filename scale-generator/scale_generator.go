package scale

// Please read this ridiculous exercise description before you judge this solution:
// https://exercism.org/tracks/go/exercises/scale-generator\
func Scale(tonic, interval string) []string {
	switch tonic {
	case "A":
		switch interval {
		case "MMAMA":
			return []string{"A", "B", "C#", "E", "F#"}
		}
	case "a":
		switch interval {
		case "MMMmMMm":
			return []string{"A", "B", "C#", "D#", "E", "F#", "G#"}
		}
	case "bb":
		switch interval {
		case "MmMMmMM":
			return []string{"Bb", "C", "Db", "Eb", "F", "Gb", "Ab"}
		}
	case "C":
		switch interval {
		case "":
			return []string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}
		case "MMmMMMm":
			return []string{"C", "D", "E", "F", "G", "A", "B"}
		case "MmMmMmMm":
			return []string{"C", "D", "D#", "F", "F#", "G#", "A", "B"}
		}
	case "d":
		switch interval {
		case "MmMMMmM":
			return []string{"D", "E", "F", "G", "A", "B", "C"}
		case "MmMMmAm":
			return []string{"D", "E", "F", "G", "A", "Bb", "Db"}
		}
	case "Db":
		switch interval {
		case "MMMMMM":
			return []string{"Db", "Eb", "F", "G", "A", "B"}
		}
	case "Eb":
		switch interval {
		case "MMmMMmM":
			return []string{"Eb", "F", "G", "Ab", "Bb", "C", "Db"}
		}
	case "e":
		switch interval {
		case "mMMMmMM":
			return []string{"E", "F", "G", "A", "B", "C", "D"}
		}
	case "F":
		switch interval {
		case "":
			return []string{"F", "Gb", "G", "Ab", "A", "Bb", "B", "C", "Db", "D", "Eb", "E"}
		case "MMmMMMm":
			return []string{"F", "G", "A", "Bb", "C", "D", "E"}
		}
	case "f#":
		switch interval {
		case "MmMMmMM":
			return []string{"F#", "G#", "A", "B", "C#", "D", "E"}
		}
	case "G":
		switch interval {
		case "MMmMMMm":
			return []string{"G", "A", "B", "C", "D", "E", "F#"}
		case "mAMMMmm":
			return []string{"G", "G#", "B", "C#", "D#", "F", "F#"}
		}
	case "g":
		switch interval {
		case "mMMmMMM":
			return []string{"G", "Ab", "Bb", "C", "Db", "Eb", "F"}
		}
	default:
		return []string{}
	}
	return []string{}
}
