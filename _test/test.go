package main

import (
	"fmt"
)

type bucketKey struct {
	bucket1 int
	bucket2 int
}

func pathInStack(solutionPath *[]bucketKey, solutionStack *[][]bucketKey) bool {
	// solutionStack: [[{0 0} {2 0} {0 2} {2 2} {1 3}]]
	// solutionPath: [{0 0} {2 0} {0 2} {2 2} {1 3}]
	var found bool = true // Start assuming true

	for _, path := range *solutionStack {
		if len(path) == len(*solutionPath) { // Lengths are same - continue searching
			for i, key := range path { // Go over each bucketkey in this path from the Stack
				if key != (*solutionPath)[i] { // If this bucketKey is NOT the same as the one in the proposed solution at same index
					found = false
				}
			}
			// If found is still true after iterating across, then we found a path that is the same as the proposed solution
			if found {
				return true
			}
		}
		found = true // Reset found to true for path in stack
	}
	return false
}

func main() {
	solutionStack := make([][]bucketKey, 0)
	solutionPath := make([]bucketKey, 0)

	solutionPath = append(solutionPath, bucketKey{0, 0})
	solutionPath = append(solutionPath, bucketKey{2, 0})
	solutionPath = append(solutionPath, bucketKey{0, 2})
	solutionPath = append(solutionPath, bucketKey{2, 2})
	solutionPath = append(solutionPath, bucketKey{1, 3})

	solutionStack = append(solutionStack, solutionPath)

	solutionPath = append(solutionPath, bucketKey{1, 3})

	fmt.Println(pathInStack(&solutionPath, &solutionStack))
}
