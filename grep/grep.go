package grep

import (
	"fmt"
	"regexp"
	"strings"
)

// flags:
// -n Print the line numbers of each matching line.
// -l Print only the names of files that contain at least one matching line.
// -i Match line using a case-insensitive comparison.
// -v Invert the program -- collect all lines that fail to match the pattern.
// -x Only match entire lines, instead of lines that contain a match.
func Search(pattern string, flags, files []string) []string {
	fmt.Println("Searching for pattern:", pattern)
	fmt.Println("Flags:", flags)
	fmt.Println("Files:", files)

	fileMap := Files(fileContentData)

	// The files we're interested in
	for _, file := range files {
		fmt.Println(file, fileMap[file])
	}

	return []string{}
}

// Returns a key of slices containing the file contents, where:
// 1. The leading and trailing whitespaces removed
// 2. Beautifying characters removed
func Files(fileContentData []string) map[string][]string {
	result := make(map[string][]string)

	var key string
	for _, line := range fileContentData {
		// Trim leading and trailing whitespaces
		line = strings.TrimSpace(line)
		if regexp.MustCompile(`.txt`).FindAllStringIndex(line, -1) != nil { // Key
			key = line // Set key for subsequent lines
		} else if regexp.MustCompile(`------`).FindAllStringIndex(line, -1) != nil { //divider
			// Do nothing
		} else { // Content
			// Replace "|" with " "
			line = strings.Replace(line, "|", "", -1)
			// Trim leading and trailing whitespaces after | cleanup
			line = strings.TrimSpace(line)
			// Get content in result for key
			result[key] = append(result[key], line)
		}
	}
	return result
}
