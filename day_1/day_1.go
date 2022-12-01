package day_1

import (
	"adventofcode2022/util"
	"sort"
	"strconv"
)

var input []string

func partOne() int {
	mostCalories := 0
	currentCalories := 0
	for _, line := range input {
		if line == "" {
			if currentCalories > mostCalories {
				mostCalories = currentCalories
			}
			currentCalories = 0
		} else {
			c, _ := strconv.Atoi(line)
			currentCalories += c
		}
	}
	if currentCalories > mostCalories {
		mostCalories = currentCalories
	}
	return mostCalories
}

func partTwo() int {
	caloryCounts := []int{}
	currentCalories := 0
	for _, line := range input {
		if line == "" {
			caloryCounts = append(caloryCounts, currentCalories)
			currentCalories = 0
		} else {
			c, _ := strconv.Atoi(line)
			currentCalories += c
		}
	}
	caloryCounts = append(caloryCounts, currentCalories)
	sort.Ints(caloryCounts)
	return caloryCounts[len(caloryCounts)-3] + caloryCounts[len(caloryCounts)-2] + caloryCounts[len(caloryCounts)-1]
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
