package day_6

import (
	"adventofcode2022/util"
	"fmt"
	"log"
	"strconv"
)

var input string

func findNDifferentCharacters(n int, s string) (int, error) {
	// find the first instance of N different characters in the string
	for i := n - 1; i < len(s); i++ {
		b := map[byte]struct{}{}
		for j := 0; j > -n; j-- {
			b[input[i+j]] = struct{}{}
		}
		if len(b) == n {
			solution := i + 1
			return solution, nil
		}
	}
	return 0, fmt.Errorf("did not find %v different characters in the string", n)
}

func partOne() int {
	solution, err := findNDifferentCharacters(4, input)
	if err != nil {
		log.Fatal(err)
	}
	return solution
}

func partTwo() int {
	solution, err := findNDifferentCharacters(14, input)
	if err != nil {
		log.Fatal(err)
	}
	return solution
}

func Call(part string, inputFile string) string {
	input = util.ParseSingleLineInput(inputFile)
	var r int
	if part == "1" {
		r = partOne()
	} else {
		r = partTwo()
	}
	return strconv.Itoa(r)
}
