package grep

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// flags:
// -n Also print the line numbers of each matching line.
// -l Print only the names of files that contain at least one matching line.
// -i Match line using a case-insensitive comparison.
// -v Invert the program -- collect all lines that fail to match the pattern.
// -x Only match entire lines, instead of lines that contain a match.
func Search(pattern string, flags, files []string) []string {
	// Flag related
	var printFileName bool
	// Append fileName?
	if len(files) > 1 {
		printFileName = true
	} else {
		printFileName = false
	}
	// Generate the map of files
	fileMap := Files(fileContentData)
	// Result slice
	result := []string{}
	// -l: Check if matches and return filenames
	if contains(flags, "-l") {
		for _, file := range files {
			if !reflect.DeepEqual(matchPattern(pattern, file, fileMap[file], printFileName, contains(flags, "-n"), contains(flags, "-i"), contains(flags, "-x"), contains(flags, "-v")), []string{}) {
				result = append(result, file)
			}
		}
		return result
	}
	// Loop over files we're interested in and append to result
	for _, file := range files {
		result = append(result, matchPattern(pattern, file, fileMap[file], printFileName, contains(flags, "-n"), contains(flags, "-i"), contains(flags, "-x"), contains(flags, "-v"))...)
	}
	return result
}

// Check if a string is in a slice
func contains(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

// If the pattern is found in the file, append that line to the result
func matchPattern(pattern, fileName string, content []string, printFileName, printLine, insensitive, wholeLine, invert bool) []string {
	result := []string{}
	var addLine bool
	for i, line := range content {
		if insensitive { // i: Convert to lowercase
			line = strings.ToLower(line)
			pattern = strings.ToLower(pattern)
		}
		if wholeLine { // -x: Turn pattern into line
			if pattern == line {
				addLine = true
			}
		} else { // Compare pattern with line
			if regexp.MustCompile(pattern).FindAllStringIndex(line, -1) != nil {
				addLine = true
			}
		}
		if invert { // -v: Invert result
			addLine = !addLine
		}
		// Append line to result based on boolean
		if addLine {
			if printFileName {
				if printLine {
					result = append(result, fmt.Sprintf("%s:%d:%s", fileName, i+1, content[i])) // -n 
				} else {
					result = append(result, fmt.Sprintf("%s:%s", fileName, content[i]))
				}
			} else {
				if printLine {
					result = append(result, fmt.Sprintf("%d:%s", i+1, content[i])) // -n
				} else {
					result = append(result, content[i])
				}
			}
		}
		addLine = false // Reset
	}
	return result
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
		} else if regexp.MustCompile(`------`).FindAllStringIndex(line, -1) != nil { // Divider
			// Do nothing
		} else if regexp.MustCompile(`^\s*$`).FindAllStringIndex(line, -1) != nil { // Empty line
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
