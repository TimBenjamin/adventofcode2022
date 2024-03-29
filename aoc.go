package main

import (
	"adventofcode2022/day_1"
	"adventofcode2022/day_10"
	"adventofcode2022/day_11"
	"adventofcode2022/day_12"
	"adventofcode2022/day_2"
	"adventofcode2022/day_3"
	"adventofcode2022/day_4"
	"adventofcode2022/day_5"
	"adventofcode2022/day_6"
	"adventofcode2022/day_7"
	"adventofcode2022/day_8"
	"adventofcode2022/day_9"
	"fmt"
	"os"
)

// aoc.go <day> <part>
func main() {

	// these exported functions must all have the same return type!
	// therefore any solutions that are ints will be converted to strings
	days := map[string]func(part string, input string) (result string){
		"day_1":  day_1.Call,
		"day_2":  day_2.Call,
		"day_3":  day_3.Call,
		"day_4":  day_4.Call,
		"day_5":  day_5.Call,
		"day_6":  day_6.Call,
		"day_7":  day_7.Call,
		"day_8":  day_8.Call,
		"day_9":  day_9.Call,
		"day_10": day_10.Call,
		"day_11": day_11.Call,
		"day_12": day_12.Call,
	}

	var day string
	var part string
	var input string
	if len(os.Args) != 4 {
		fmt.Println("Incorrect number of args: aoc.go <day> <part> <path/to/input.txt>")
	} else {
		day = os.Args[1]
		part = os.Args[2]
		input = os.Args[3]
		f := "day_" + day
		result := days[f](part, input)
		fmt.Println("result:", result)
	}
}
