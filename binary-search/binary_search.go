package binarysearch

func SearchInts(list []int, key int) int {
	var l, r int = 0, len(list) - 1
	// Loop until l and r meet
	for l <= r {
		var m int = (l + r) / 2
		if list[m] < key {
			l = m + 1 // Go to right half
		} else if list[m] > key {
			r = m - 1 // Go to left half
		} else {
			return m // Found
		}
	}
	return -1 // Not found
}
