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
	// Input
	fmt.Println("Size of bucket one:", sizeBucketOne) // the size of bucket one
	fmt.Println("Size of bucket two:", sizeBucketTwo) // the size of bucket two
	fmt.Println("Goal amount:", goalAmount)           // the desired number of liters to reach
	fmt.Println("Start bucket:", startBucket)         // which bucket to fill first - "one" or "two"

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
	solutionPath := make([]bucketKey, 0)    // solution path
	solutionStack := make([][]bucketKey, 0) // stack of Solutions
	var possible bool

	// Add in the initial state so we do not revisit
	solutionPath = append(solutionPath, bucketKey{0, 0})
	moves++

	fmt.Println("================")
	fmt.Println("B1, B2")
	if startBucket == "one" {
		possible = bucketSolver(startBucket, sizeBucketOne, sizeBucketOne, 0, sizeBucketTwo, goalAmount, &solutionPath, &solutionStack)
	} else {
		possible = bucketSolver(startBucket, 0, sizeBucketOne, sizeBucketTwo, sizeBucketTwo, goalAmount, &solutionPath, &solutionStack)
	}
	fmt.Println("================")
	fmt.Println("Solution Stack:", solutionStack)
	fmt.Println("Solution Path:", solutionPath)
	fmt.Println("================")

	// ======================================================================
	sol1 := solutionPath[len(solutionPath)-1].bucket1
	sol2 := solutionPath[len(solutionPath)-1].bucket2
	moves = len(solutionPath) - 1

	fmt.Println("Solution: ", sol1, ",", sol2)

	if sol1 == goalAmount {
		goalBucket = "one"
		otherBucket = sol2
	} else {
		goalBucket = "two"
		otherBucket = sol1
	}

	if !possible {
		e = fmt.Errorf("no solution found")
	}
	// ======================================================================
	// Validation
	fmt.Println("Goal bucket:", goalBucket)
	fmt.Println("# of Moves:", moves)
	fmt.Println("Other bucket:", otherBucket)
	fmt.Println("Error:", e)

	return goalBucket, moves, otherBucket, e
}

func visited(key bucketKey, solutionPath []bucketKey) bool {
	for _, v := range solutionPath {
		if v == key {
			return true
		}
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
func bucketSolver(startBucket string, amt1, size1, amt2, size2, goal int, solutionPath *[]bucketKey, solutionStack *[][]bucketKey) bool {
	// Check if the goal is already reached
	if (amt1 == goal) || (amt2 == goal) {
		// Add step to solution Path
		*solutionPath = append(*solutionPath, bucketKey{amt1, amt2})

		// Add solutionPath to solutionStack
		*solutionStack = append(*solutionStack, *solutionPath)

		fmt.Println(amt1, ",", amt2, " | Done!")
		return true
	}
	// Check if starting bucket is empty and other is full - invalid state
	if (startBucket == "one") && (amt1 == 0) && (amt2 == size2) {
		return false
	} else if (startBucket == "two") && (amt2 == 0) && (amt1 == size1) {
		return false
	}

	// Checks if we have already visited this state
	// If not - proceed
	if !visited(bucketKey{amt1, amt2}, *solutionPath) {
		fmt.Println(amt1, ",", amt2)

		// Add to solution Path
		*solutionPath = append(*solutionPath, bucketKey{amt1, amt2})

		// For inter-bucket transfers - e.g. for Bucket 1 as giver
		// amt1: amount of water in giver
		// size2 - amt2: amount of water taker can take
		// Minimum of these 2 numbers is the amount of water that can be transferred
		from1to2 := MinInt(amt1, size2-amt2)
		from2to1 := MinInt(amt2, size1-amt1)

		// Checks all 6 actions recursively and see if we can reach the goal down any path via ||
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
