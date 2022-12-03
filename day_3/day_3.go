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

func partTwo() int {
	priorityTotal := 0
	// surely the input size is divisible by 3 ;-)
	for i := 0; i < len(input); i += 3 {
		// need to find the one common letter in these 3 strings
		// there is only one, so as soon as it's found, we finish
	outer:
		for _, b1 := range input[i] {
			for _, b2 := range input[i+1] {
				if b2 == b1 {
					for _, b3 := range input[i+2] {
						if b3 == b1 {
							priorityTotal += getPriority((b1))
							break outer
						}
					}
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
