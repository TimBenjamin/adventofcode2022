package day_2

import (
	"adventofcode2022/util"
	"strconv"
	"strings"
)

var input []string

func doWin(score int) int {
	return score + 6
}
func doDraw(score int) int {
	return score + 3
}
func partOne() int {
	shapes := map[string]int{
		"A": 1, // Rock
		"B": 2, // Paper
		"C": 3, // Scissors
		"X": 1, // Rock
		"Y": 2, // Paper
		"Z": 3, // Scissors
	}
	score := 0
	for _, line := range input {
		pair := strings.Split(line, " ")
		// first add the score for the shape I selected
		score += shapes[pair[1]]
		// next add the score for the outcome
		if shapes[pair[0]] == shapes[pair[1]] {
			score = doDraw(score)
		} else if pair[0] == "A" && pair[1] == "Y" {
			score = doWin(score)
		} else if pair[0] == "B" && pair[1] == "Z" {
			score = doWin(score)
		} else if pair[0] == "C" && pair[1] == "X" {
			score = doWin(score)
		}
		// (add nothing if it's a lose)
	}
	return score
}

func partTwo() int {
	shapes := map[string]int{
		"A": 1, // Rock
		"B": 2, // Paper
		"C": 3, // Scissors
	}
	score := 0
	for _, line := range input {
		pair := strings.Split(line, " ")
		// the shape needed is in a modulo order
		if pair[1] == "X" {
			// lose
			// ((shapes[pair[0]] - 1 + 2) % 3) + 1
			score += ((shapes[pair[0]] + 1) % 3) + 1
		} else if pair[1] == "Y" {
			// draw
			score += shapes[pair[0]]
			score = doDraw(score)
		} else if pair[1] == "Z" {
			// win
			// ((shapes[pair[0]] - 1 + 1) % 3) + 1
			score += (shapes[pair[0]] % 3) + 1
			score = doWin(score)
		}
	}
	return score
}

func Call(part string, inputFile string) string {
	input = util.ParseInputIntoLines(inputFile)
	var r int
	if part == "1" {
		r = partOne()
	} else {
		r = partTwo()
	}
	return strconv.Itoa(r)
}
