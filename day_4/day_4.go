package day_4

import (
	"adventofcode2022/util"
	"strconv"
	"strings"
)

var input []string

func partOne() int {
	containments := 0
	for _, line := range input {
		min1, _ := strconv.Atoi(strings.Split(strings.Split(line, ",")[0], "-")[0])
		max1, _ := strconv.Atoi(strings.Split(strings.Split(line, ",")[0], "-")[1])
		min2, _ := strconv.Atoi(strings.Split(strings.Split(line, ",")[1], "-")[0])
		max2, _ := strconv.Atoi(strings.Split(strings.Split(line, ",")[1], "-")[1])
		if (min1 <= min2 && max1 >= max2) || (min2 <= min1 && max2 >= max1) {
			containments++
		}
	}
	return containments
}

func partTwo() int {
	overlaps := 0
	for _, line := range input {
		min1, _ := strconv.Atoi(strings.Split(strings.Split(line, ",")[0], "-")[0])
		max1, _ := strconv.Atoi(strings.Split(strings.Split(line, ",")[0], "-")[1])
		min2, _ := strconv.Atoi(strings.Split(strings.Split(line, ",")[1], "-")[0])
		max2, _ := strconv.Atoi(strings.Split(strings.Split(line, ",")[1], "-")[1])
		if (max1 >= min2 && max1 <= max2) || (min1 >= min2 && min1 <= max2) || (max2 >= min1 && max2 <= max1) || (min2 >= min1 && min2 <= max1) {
			overlaps++
		}
	}
	return overlaps
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
