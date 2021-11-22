package robotname

import (
	"errors"
	"fmt"
)

// Global variable to track the name of the robot
var GlobalNum []int = []int{0, 0, 0}

// Define the Robot type
type Robot struct {
	name string
}

// Method to return a unique Robot name
func (r *Robot) Name() (string, error) {
    // Return new name if we don't have one yet
	if r.name == "" {
        // Calculate name based on GlobalNum
        robotName := calculateName(GlobalNum)
		// Increment GlobalNum
		var err error
		GlobalNum, err = incrementNum(GlobalNum)
		// Check if out of range
		if err == nil {
			r.name = robotName
			return robotName, nil
		} else {
			return "", err
		}
    } else {
        // If name already exists, return existing name
        return r.name, nil
    }
}

// Method to reset name of robot
func (r *Robot) Reset() {
	// Reset name
	r.name = ""
}

// Increment GlobalNum according to our algorithm
func incrementNum(num []int) ([]int, error) {
    num[2]++
    // Increment second char
    if num[2] >= 1000 {
        num[1]++
        num[2] = 0 // Back to 0
    }
    // Increment first char
    if num[1] >= 26 {
        num[0]++
        num[1] = 0 // Back to A
    }
    // Check if we have reached the end
    if num[0] >= 26 && num[2] > 0 {
        return nil, errors.New("index out of range")
    }
    return num, nil
}

// Uses n to deterministically calculate the name of the robot
func calculateName(num []int) string {
    // Represents `AA000`
    arr := [3]int{65,65,0}
    var str string
    // Calculate name[0]
    str = str + string(rune(arr[0] + num[0]))
    // Calculate name[1]
    str = str + string(rune(arr[1] + num[1]))
    // Calculate name[2:4]
    str = str + fmt.Sprintf("%03d", num[2])
    return str
}

