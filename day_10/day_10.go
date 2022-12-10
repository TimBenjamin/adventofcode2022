package day_10

import (
	"adventofcode2022/util"
	"fmt"
	"strconv"
	"strings"
)

var input []string

func getCommands() [][]string {
	// `noop` takes 1 cycle and does nothing
	// `addx V` takes 2 cycles and adds V to x
	// rather than keep track of 2 cycles, I'll add an extra noop before every addx and then let addx take 1 cycle.
	commands := [][]string{}
	for _, line := range input {
		if line == "noop" {
			commands = append(commands, []string{"noop"})
		} else {
			sp := strings.Split(line, " ")
			commands = append(commands, []string{"noop"})
			commands = append(commands, []string{sp[0], sp[1]})
		}
	}
	// add on a noop in case the last command is an addx
	commands = append(commands, []string{"noop"})
	return commands
}

func partOne() int {
	commands := getCommands()
	x := 1
	currentCycle := 0
	signalStrengthTotal := 0
	for _, command := range commands {
		currentCycle++
		if currentCycle != 0 && (currentCycle == 20 || (currentCycle-20)%40 == 0) {
			fmt.Printf("Value of x during cycle %v is: %v for a strength of: %v\n", currentCycle, x, (currentCycle * x))
			signalStrengthTotal += (currentCycle * x)
		}
		// the add takes place at the END of the cycle
		if command[0] == "addx" {
			v, _ := strconv.Atoi(command[1])
			x += v
		}
	}
	return signalStrengthTotal
}

func printCrt(crt [][]string) {
	fmt.Println()
	for i := 0; i < len(crt); i++ {
		for j := 0; j < len(crt[i]); j++ {
			fmt.Printf(crt[i][j])
		}
		fmt.Println()
	}
	fmt.Println()
}

func partTwo() int {
	commands := getCommands()
	crt := [][]string{}
	line := []string{}
	pixelPosition := 0
	x := 1
	for _, command := range commands {
		if pixelPosition > 0 && pixelPosition%40 == 0 {
			crt = append(crt, line)
			line = []string{}
			pixelPosition = 0
		}

		if x >= pixelPosition-1 && x <= pixelPosition+1 {
			line = append(line, "#")
		} else {
			line = append(line, ".")
		}
		pixelPosition++
		// the add takes place at the END of the cycle
		if command[0] == "addx" {
			v, _ := strconv.Atoi(command[1])
			x += v
		}
	}
	printCrt(crt)
	return 0
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
