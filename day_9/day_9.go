package day_9

import (
	"adventofcode2022/util"
	"fmt"
	"strconv"
	"strings"
)

var input []string

type Instruction struct {
	direction string
	amount    int
}

func getInstructions() []Instruction {
	instructions := []Instruction{}
	for _, line := range input {
		sp := strings.Split(line, " ")
		amount, _ := strconv.Atoi(sp[1])
		instructions = append(instructions, Instruction{direction: sp[0], amount: amount})
	}
	return instructions
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func partOne() int {
	instructions := getInstructions()
	headPositionLog := [][]int{}
	headPositionLog = append(headPositionLog, []int{0, 0})
	tailPositionSet := map[string]struct{}{} // I'll log the positions of the tail in here as a set so I don't double count
	currentHeadPosition := []int{0, 0}
	currentTailPosition := []int{0, 0}
	position := strconv.Itoa(currentTailPosition[0]) + "-" + strconv.Itoa(currentTailPosition[1])
	tailPositionSet[position] = struct{}{}
	for _, Instruction := range instructions {
		for i := 0; i < Instruction.amount; i++ {

			if Instruction.direction == "U" {
				currentHeadPosition[1]++
			} else if Instruction.direction == "D" {
				currentHeadPosition[1]--
			} else if Instruction.direction == "R" {
				currentHeadPosition[0]++
			} else if Instruction.direction == "L" {
				currentHeadPosition[0]--
			}
			dx := currentHeadPosition[0] - currentTailPosition[0]
			dy := currentHeadPosition[1] - currentTailPosition[1]

			if dy == 0 {
				if dx > 1 {
					currentTailPosition[0]++
				} else if dx < -1 {
					currentTailPosition[0]--
				}
			} else if dx == 0 {
				if dy > 1 {
					currentTailPosition[1]++
				} else if dy < -1 {
					currentTailPosition[1]--
				}
			} else if abs(dx) == 2 || abs(dy) == 2 {
				// tail gets pulled to where head last was
				currentTailPosition = headPositionLog[len(headPositionLog)-1]
			}
			hp := []int{currentHeadPosition[0], currentHeadPosition[1]}
			headPositionLog = append(headPositionLog, hp)
			position := strconv.Itoa(currentTailPosition[0]) + "-" + strconv.Itoa(currentTailPosition[1])
			tailPositionSet[position] = struct{}{}
		}
	}

	/*
		Tail positions should be :
		[0 0]
		[0 0]
		[1 0]
		[2 0]
		[3 0]
		[3 0]
		[4 1]
		[4 2]
		[4 3]
		[4 3]
		[3 4]
		[2 4]
		[2 4]
		[2 4]
		[2 4]
		[3 3]
		[4 3]
		[4 3]
		[4 3]
		[4 3]
		[3 2]
		[2 2]
		[1 2]
		[1 2]
		[1 2]
	*/
	return len(tailPositionSet)
}

func partTwo() int {
	// similar to part 1, except rather than just headPositionLog
	// I need a log that is 10 long, where the head is in [0]
	instructions := getInstructions()
	// position is []int{x,y}
	// position log is [][]int
	// the whole snake is [][][]int
	snake := [][][]int{}
	for i := 0; i < 10; i++ {
		positionLog := [][]int{}
		positionLog = append(positionLog, []int{0, 0})
		snake = append(snake, positionLog)
	}
	tailPositionSet := map[string]struct{}{} // I'll log the positions of the tail in here as a set so I don't double count
	position := strconv.Itoa(snake[9][0][0]) + "," + strconv.Itoa(snake[9][0][1])
	tailPositionSet[position] = struct{}{}
	for _, Instruction := range instructions {
		for i := 0; i < Instruction.amount; i++ {
			// do the instruction to the head ofthe snake (snake[0])
			// then apply the rules for drag-along to elements 1-9
			// first add a new item to the log for each element of the snake, which can be the same as the previous step
			for s := 0; s < 10; s++ {
				position := []int{snake[s][len(snake[s])-1][0], snake[s][len(snake[s])-1][1]}
				snake[s] = append(snake[s], position)
			}
			if Instruction.direction == "U" {
				snake[0][len(snake[0])-1][1]++
			} else if Instruction.direction == "D" {
				snake[0][len(snake[0])-1][1]--
			} else if Instruction.direction == "R" {
				snake[0][len(snake[0])-1][0]++
			} else if Instruction.direction == "L" {
				snake[0][len(snake[0])-1][0]--
			}
			// now apply the dragging rules to each element that follows the head, in turn
			for s := 1; s < 10; s++ {
				dx := snake[s-1][len(snake[s-1])-1][0] - snake[s][len(snake[s])-1][0]
				dy := snake[s-1][len(snake[s-1])-1][1] - snake[s][len(snake[s])-1][1]
				if dy == 0 {
					if dx > 1 {
						snake[s][len(snake[s])-1][0]++
					} else if dx < -1 {
						snake[s][len(snake[s])-1][0]--
					}
				} else if dx == 0 {
					if dy > 1 {
						snake[s][len(snake[s])-1][1]++
					} else if dy < -1 {
						snake[s][len(snake[s])-1][1]--
					}
				} else if abs(dx) == 2 || abs(dy) == 2 {
					// this element gets pulled to where the previous segment last was
					// here is where the bug is - this needs to cascade somehow
					snake[s][len(snake[s])-1] = snake[s-1][len(snake[s-1])-2]
				}
			}
			for _, s := range snake {
				fmt.Println(s)
			}
			fmt.Println("---")
			// update the set of tail positions
			position := strconv.Itoa(snake[9][len(snake[9])-1][0]) + "," + strconv.Itoa(snake[9][len(snake[9])-1][1])
			tailPositionSet[position] = struct{}{}
		}
	}
	for k := range tailPositionSet {
		fmt.Println(k)
	}
	return len(tailPositionSet)
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
