package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for i := 0; i < len(birdsPerDay); i++ {
		total += birdsPerDay[i]
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	// Get our start and end index for the specific week
	startindex := (week - 1) * 7
	endindex := startindex + 7
	weekList := birdsPerDay[startindex:endindex]
	return TotalBirdCount(weekList)
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	// TODO:
	// Add +1 to every other day, starting from index 0	
	for i:=0; i<len(birdsPerDay); i+=2 {
		birdsPerDay[i] += 1
	}
	return birdsPerDay
}
