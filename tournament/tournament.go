package tournament

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

func cleaner(input string) []string {
	// Split by newline for each line
	lines := strings.Split(input, "\n")

	// Make output slice
	output := make([]string, 0)

	// Loop over and clean each line
	for _, line := range lines {
		// Remove tabs from each line
		line = strings.TrimSpace(line)

		// Remove comments and empty lines
		if strings.HasPrefix(line, "#") != true && line != "" {
			output = append(output, line)
			//fmt.Println("line: ", line)
		}
	}
	return output
}

// Struct for storing tally
type teamScore struct {
	name string
	mp   int
	w    int
	d    int
	l    int
	p    int
}

// Function to generate score from cleaned input
func generateScore(cleaned []string) []teamScore {
	var tally = make(map[string]teamScore)
	var team, opponent, outcome string
	var teamStruct, opponentStruct teamScore

	for _, line := range cleaned {
		// Split into data points
		linesplit := strings.Split(line, ";")
		team = linesplit[0]
		opponent = linesplit[1]
		outcome = linesplit[2]

		// Copy of teamScore for the Team and Opponent
		teamStruct = tally[team]
		opponentStruct = tally[opponent]

		// Initialize structs if empty
		if teamStruct.name == "" {
			teamStruct = teamScore{name: team}
		}
		if opponentStruct.name == "" {
			opponentStruct = teamScore{name: opponent}
		}

		// MP: Increment matches played
		teamStruct.mp++
		opponentStruct.mp++

		// W: Matches Won
		if outcome == "win" {
			teamStruct.w++
			opponentStruct.l++
		}

		// D: Matches Drawn (Tied)
		if outcome == "draw" {
			teamStruct.d++
			opponentStruct.d++
		}

		// L: Matches Lost
		if outcome == "loss" {
			teamStruct.l++
			opponentStruct.w++
		}

		// P: Points
		if outcome == "win" {
			teamStruct.p += 3
		} else if outcome == "draw" {
			teamStruct.p += 1
			opponentStruct.p += 1
		} else if outcome == "loss" {
			opponentStruct.p += 3
		}

		// Update teamStruct to tally
		tally[team] = teamStruct
		tally[opponent] = opponentStruct
	}

	// Convert tally to slice
	var output = make([]teamScore, 0)
	for _, team := range tally {
		output = append(output, team)
	}
	return output
}

// Append whitespace to the end of the string until it is 31 characters long
func pad(s string) string {
	for len(s) < 31 {
		s += " "
	}
	return s
}

// Return score in expected format from tally
func returnScore(tally []teamScore) string {
	var output = "Team                           | MP |  W |  D |  L |  P"
	for _, team := range tally {
		output += fmt.Sprintf("\n%s| %2d | %2d | %2d | %2d | %2d",
			pad(team.name), team.mp, team.w, team.d, team.l, team.p)
	}
	return output + string(rune(10))
}

// Sort tally in expected format
func sortTally(tally []teamScore) []teamScore {
	// Sort by points, then by name
	sort.Slice(tally, func(i, j int) bool {
		if tally[i].p == tally[j].p {
			return tally[i].name < tally[j].name
		}
		return tally[i].p > tally[j].p
	})
	return tally
}

// Function to tally up the score in expected output format
func Tally(reader io.Reader, writer io.Writer) error {
	// Split reader into lines
	buf := new(strings.Builder)
	_, err := io.Copy(buf, reader)

	// Clean up string
	cleaned := cleaner(string(buf.String()))

	// Loop over each line
	for _, line := range cleaned {
		// Split each line into teams
		outcome := strings.Split(line, ";")

		if len(outcome) != 3 {
			return errors.New("Each line must contain three semicolon-separated values")
		}

		// Check for errors
		if outcome[2] != "win" && outcome[2] != "loss" && outcome[2] != "draw" {
			return errors.New("Invalid input")
		}
	}

	// Generate array of teamScores
	tally := generateScore(cleaned)

	// Sort tally in expected format
	tally = sortTally(tally)

	// Add returnSCore to buffer
	_, _ = fmt.Fprintf(writer, returnScore(tally))

	return err
}
