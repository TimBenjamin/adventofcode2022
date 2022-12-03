package day_3

import (
	"adventofcode2022/util"
	"fmt"
	"strconv"
)

func getPriority(r rune) int {
	i := int(r)
	if i > 96 && i < 123 {
		return i - 96
	}
	if i > 64 && i < 91 {
		return i - 38
	}
	fmt.Printf("unhandled letter %v - val %v\n", r, i)
	return 0
}

var input []string

func partOne() int {
	priorityTotal := 0
	for _, line := range input {
		// find the letters that appear in both bags
		// there is only one, so as soon as it's found, we finish
		// inefficiently...
	outer:
		for _, b1 := range line[0 : len(line)/2] {
			for _, b2 := range line[len(line)/2:] {
				if b1 == b2 {
					priorityTotal += getPriority(b1)
					break outer
				}
			}
		}
	}
	return priorityTotal
}

// for part 2, I want to avoid repeatedly iterating, so will use a set instead:
func toSet(s string) map[rune]struct{} {
	out := map[rune]struct{}{}
	for _, c := range s {
		out[c] = struct{}{}
	}
	return out
}

func partTwo() int {
	priorityTotal := 0
	// surely the input size is divisible by 3 ;-)
	for i := 0; i < len(input); i += 3 {
		line1 := toSet(input[i])
		line2 := toSet(input[i+1])
		line3 := toSet(input[i+2])
		for b := range line1 {
			if _, ok := line2[b]; ok {
				if _, ok := line3[b]; ok {
					priorityTotal += getPriority(b)
					break
				}
			}
		}
	}
	return priorityTotal
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
