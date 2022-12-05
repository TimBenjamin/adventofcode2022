package day_5

import (
	"adventofcode2022/util"
	"fmt"
	"strconv"
	"strings"
)

var input []string

// NB this generates stacks with the top item at index 0
func getStacks() [][]rune {
	stacks := [][]rune{}
	for _, line := range input {
		if len(line) == 0 {
			break
		}
		stackIndex := 0
		for i, r := range line {
			if i-1 == 0 || (i-1)%4 == 0 {
				if string(r) == "1" {
					break
				}
				if string(r) == " " {
					// this position is empty, but create an empty stack in this position if it doesn't exist
					if len(stacks) <= stackIndex {
						stack := []rune{}
						stacks = append(stacks, stack)
					}
					stackIndex++
					continue
				}
				if len(stacks) <= stackIndex {
					stack := []rune{r}
					stacks = append(stacks, stack)
				} else {
					stacks[stackIndex] = append(stacks[stackIndex], r)
				}
				stackIndex++
			}
		}
	}
	return stacks
}

type Instruction struct {
	move int
	from int
	to   int
}

func getInstructions() []Instruction {
	instructions := []Instruction{}
	for _, line := range input {
		if strings.Contains(line, "move") {
			sp := strings.Split(line, " ")
			// move 10 from 2 to 7
			move, _ := strconv.Atoi(sp[1])
			from, _ := strconv.Atoi(sp[3])
			to, _ := strconv.Atoi(sp[5])
			i := Instruction{
				move: move,
				from: from - 1,
				to:   to - 1,
			}
			instructions = append(instructions, i)
		}
	}
	return instructions
}

// func printStacks(stacks [][]rune) {
// 	for _, stack := range stacks {
// 		fmt.Printf("[")
// 		for _, s := range stack {
// 			fmt.Printf("%v", string(s))
// 		}
// 		fmt.Printf("]\n")
// 	}
// }

func partOne() int {
	stacks := getStacks()
	instructions := getInstructions()
	for _, instruction := range instructions {
		for i := instruction.move; i > 0; i-- {
			crate := stacks[instruction.from][0]
			stacks[instruction.from] = stacks[instruction.from][1:]
			stacks[instruction.to] = append([]rune{crate}, stacks[instruction.to]...)
		}
	}
	solution := ""
	for _, stack := range stacks {
		solution += string(stack[0])
	}
	fmt.Println(solution)
	return 0
}

func partTwo() int {
	stacks := getStacks()
	instructions := getInstructions()
	for _, instruction := range instructions {
		crates := stacks[instruction.from][0:instruction.move]
		remains := stacks[instruction.from][instruction.move:]
		stacks[instruction.from] = []rune{}
		stacks[instruction.from] = append(stacks[instruction.from], remains...)
		stacks[instruction.to] = append(crates, stacks[instruction.to]...)
	}
	solution := ""
	for _, stack := range stacks {
		solution += string(stack[0])
	}
	fmt.Println(solution)
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
