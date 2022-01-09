package twobucket

import "fmt"

// ######################################################################
//                               THEORY
// ######################################################################
// Memoization: https://www.youtube.com/watch?v=a7EjmdQzPqY
// ######################################################################
// Basically caches function call in a map
// ######################################################################
// 							RULES  & ACTIONS
// ======================================================================
// RULE 1: Pouring one bucket into the other bucket until either:
// 		a) the first bucket is empty
// 		b) the second bucket is full
// ======================================================================
// ACTION 1: First bucket to second bucket
// ACTION 2: Second bucket to first bucket
// ======================================================================
// RULE 2: Emptying a bucket and doing nothing to the other.
// ======================================================================
// ACTION 3: Empty the first bucket completely
// ACTION 4: Empty the second bucket completely
// ======================================================================
// RULE 3: Filling a bucket and doing nothing to the other.
// ======================================================================
// ACTION 5: Fill the first bucket
// ACTION 6: Fill the second bucket
// ######################################################################

type bucketKey struct {
	bucket1 int
	bucket2 int
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Solve(sizeBucketOne, sizeBucketTwo, goalAmount int, startBucket string) (string, int, int, error) {
	// Error handling
	if sizeBucketOne <= 0 || sizeBucketTwo <= 0 || goalAmount <= 0 {
		return "", 0, 0, fmt.Errorf("invalid bucket sizes")
	}
	if (startBucket != "one") && (startBucket != "two") {
		return "", 0, 0, fmt.Errorf("invalid start bucket name")
	}

	// Output
	var goalBucket string // which bucket should end up with the desired number of liters
	var moves int         // the number of moves required to reach the desired number of liters
	var otherBucket int   // the number of liters remaining in the other bucket
	var e error           // error

	// ======================================================================
	// Implementation
	// ======================================================================
	// Recursion - visit all 6 possible actions one by one until one of them
	// returns True.
	// To avoid infinite recursion, we use a map to store the results of
	// previous calls.
	// ======================================================================
	var solutionPath []bucketKey
	solutionStack := make([][]bucketKey, 0) // stack of Solutions
	var possible bool = true

	for possible { // While possible is true, keep iterating
		solutionPath = make([]bucketKey, 0)                  // Initialize solutionPath
		solutionPath = append(solutionPath, bucketKey{0, 0}) // Add in the initial state so we do not revisit

		if startBucket == "one" { // Force filling the bucket
			possible = bucketSolver(startBucket, sizeBucketOne, sizeBucketOne, 0, sizeBucketTwo, goalAmount, solutionPath, &solutionStack)
		} else {
			possible = bucketSolver(startBucket, 0, sizeBucketOne, sizeBucketTwo, sizeBucketTwo, goalAmount, solutionPath, &solutionStack)
		}
	}

	// If no solution in stack, return error
	if len(solutionStack) == 0 {
		e = fmt.Errorf("no solution found")
		return "", 0, 0, e
	}

	// Find the smallest solution in the stack
	solutionPath = smallestPathInStack(&solutionStack)
	sol1 := solutionPath[len(solutionPath)-1].bucket1
	sol2 := solutionPath[len(solutionPath)-1].bucket2
	moves = len(solutionPath) - 1

	if sol1 == goalAmount {
		goalBucket = "one"
		otherBucket = sol2
	} else {
		goalBucket = "two"
		otherBucket = sol1
	}

	return goalBucket, moves, otherBucket, e
}

// Find the smallest path in the stack - most optimal solution
func smallestPathInStack(solutionStack *[][]bucketKey) []bucketKey {
	var smallestPath []bucketKey = (*solutionStack)[0]
	var smallestPathLength int = len(smallestPath)
	for _, path := range *solutionStack {
		if len(path) < smallestPathLength {
			smallestPath = path // Update smallest path
			smallestPathLength = len(path)
		}
	}
	return smallestPath
}

// Check if node is already visited
func visited(key bucketKey, solutionPath []bucketKey) bool {
	for _, v := range solutionPath {
		if v == key {
			return true
		}
	}
	return false
}

// Check if proposed solution has already been traversed in a solution contained in the stack
func pathInStack(solutionPath []bucketKey, solutionStack *[][]bucketKey) bool {
	var found bool = true // Start assuming true

	for _, path := range *solutionStack {
		for i, key := range path { // Go over each bucketkey in this path from the Stack
			// Check we are not out of bounds for solutionPath
			if !(i >= len(solutionPath)) {
				if key != (solutionPath)[i] { // If this bucketKey is NOT the same as the one in the proposed solution at same index
					found = false
				}
			}
		}
		// If found is still true after iterating across the path in the Stack, then we found a path that is the same as the proposed solution up until the last bucketKey
		if found {
			return true
		}
		found = true // Reset found to true for path in stack
	}
	return false
}

// Recursive function which prints the intermediate steps to reach the final
// solution and return boolean value. Only returns false if all possible permutations are exhausted

// INPUTS:
// amt1 and amt2 are the amount of water present
// size1 and size2 are the capacity of the buckets
// goal is the desired amount of water in the buckets
// visited is a map which stores the result of any intermediate steps to avoid infinite recursion

// OUTPUTS:
// True if solution is possible, otherwise False.
func bucketSolver(startBucket string, amt1, size1, amt2, size2, goal int, solutionPath []bucketKey, solutionStack *[][]bucketKey) bool {
	// Check if starting bucket is empty and other is full - invalid state - even if it's the goal
	if (startBucket == "one") && (amt1 == 0) && (amt2 == size2) {
		return false
	} else if (startBucket == "two") && (amt2 == 0) && (amt1 == size1) {
		return false
	}

	// Check if:
	// 1. the goal is already reached and,
	// 2. we haven't visited this key state yet
	if ((amt1 == goal) || (amt2 == goal)) && (!visited(bucketKey{amt1, amt2}, solutionPath)) {
		// Add final step to solution Path
		solutionPath = append(solutionPath, bucketKey{amt1, amt2})
		// Add solutionPath to solutionStack only if it is not already in the stack
		if !pathInStack(solutionPath, solutionStack) {
			// Add solutionPath to solutionStack
			*solutionStack = append(*solutionStack, solutionPath)
			return true
		} else {
			// If solutionPath is already in the stack, then we have reached a duplicate solution
			return false
		}
	}

	// Checks if we have already visited this state
	// If not - proceed
	if !visited(bucketKey{amt1, amt2}, solutionPath) {
		// Add to solution Path
		solutionPath = append(solutionPath, bucketKey{amt1, amt2})

		// For inter-bucket transfers - e.g. for Bucket 1 as giver
		// amt1: amount of water in giver
		// size2 - amt2: amount of water taker can take
		// Minimum of these 2 numbers is the amount of water that can be transferred
		from1to2 := MinInt(amt1, size2-amt2)
		from2to1 := MinInt(amt2, size1-amt1)

		// Checks all 6 actions recursively and see if we can reach the goal down any path via ||
		// OR should mean that if there is a single path that can reach the goal, then we return true
		return (
		// ACTION 1: First bucket to second bucket
		bucketSolver(startBucket, amt1-from1to2, size1, amt2+from1to2, size2, goal, solutionPath, solutionStack) ||
			// ACTION 2: Second bucket to first bucket
			bucketSolver(startBucket, amt1+from2to1, size1, amt2-from2to1, size2, goal, solutionPath, solutionStack) ||
			// ACTION 3: Empty the first bucket completely
			bucketSolver(startBucket, 0, size1, amt2, size2, goal, solutionPath, solutionStack) ||
			// ACTION 4: Empty the second bucket completely
			bucketSolver(startBucket, amt1, size1, 0, size2, goal, solutionPath, solutionStack) ||
			// ACTION 5: Fill the first bucket
			bucketSolver(startBucket, size1, size1, amt2, size2, goal, solutionPath, solutionStack) ||
			// ACTION 6: Fill the second bucket
			bucketSolver(startBucket, amt1, size1, size2, size2, goal, solutionPath, solutionStack))
	} else {
		// If we have visited this state - return False
		return false
	}
}
