package day_11

import (
	"adventofcode2022/util"
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

var input []string

type Monkey struct {
	number         int
	items          []int
	operation      string
	divisibleTest  int
	ifTrueTo       int
	ifFalseTo      int
	numInspections int
}

/*
Monkey 0:

	  Starting items: 79, 98
	  Operation: new = old * 19
	  Test: divisible by 23
		If true: throw to monkey 2
		If false: throw to monkey 3
*/
func getMonkeys() []Monkey {
	monkeys := []Monkey{}
	for i := 0; i < len(input); i++ {
		if len(input[i]) > 6 && input[i][0:6] == "Monkey" {
			number, _ := strconv.Atoi(input[i][7:8])
			itemsString := strings.Split(input[i+1], ": ")[1]
			itemsStrings := strings.Split(itemsString, ", ")
			items := []int{}
			for _, itemString := range itemsStrings {
				item, _ := strconv.Atoi(itemString)
				items = append(items, item)
			}
			operation := input[i+2][19:]
			divisibleTest, _ := strconv.Atoi(input[i+3][21:])
			ifTrueTo, _ := strconv.Atoi(input[i+4][29:])
			ifFalseTo, _ := strconv.Atoi(input[i+5][30:])
			monkeys = append(monkeys, Monkey{
				number:         number,
				items:          items,
				operation:      operation,
				divisibleTest:  divisibleTest,
				ifTrueTo:       ifTrueTo,
				ifFalseTo:      ifFalseTo,
				numInspections: 0,
			})
			//fmt.Printf("Monkey: %v\n", monkeys[len(monkeys)-1])
			i += 6
		}
	}
	return monkeys
}

func applyOperation(a int, operation string) int {
	ops := strings.Split(operation, " ")
	// first operand is always "old"
	var b int
	if ops[2] != "old" {
		b, _ = strconv.Atoi(ops[2])
	} else {
		b = a
	}
	// operator is only ever * or +
	if ops[1] == "+" {
		return a + b
	} else if ops[1] == "*" {
		return a * b
	} else {
		log.Fatalf("Invalid operator: %v\n", ops[1])
		return 0
	}
}

func partOne() int {
	monkeys := getMonkeys()
	numRounds := 20
	for i := 0; i < numRounds; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				worryLevel := applyOperation(item, monkey.operation)
				worryLevel /= 3
				if worryLevel%monkey.divisibleTest == 0 {
					monkeys[monkey.ifTrueTo].items = append(monkeys[monkey.ifTrueTo].items, worryLevel)
				} else {
					monkeys[monkey.ifFalseTo].items = append(monkeys[monkey.ifFalseTo].items, worryLevel)
				}
				monkeys[monkey.number].numInspections++
			}
			monkeys[monkey.number].items = []int{}
		}
		fmt.Printf("After round %v monkey items are:\n", i+1)
		for _, monkey := range monkeys {
			fmt.Printf("Monkey %v: %v\n", monkey.number, monkey.items)
		}
		fmt.Println()
	}
	//fmt.Println("Monkey activity:")
	activity := []int{}
	for _, monkey := range monkeys {
		//fmt.Printf("Monkey %v: %v\n", monkey.number, monkey.numInspections)
		activity = append(activity, monkey.numInspections)
	}
	sort.Ints(activity)
	return activity[len(activity)-1] * activity[len(activity)-2]
}

func applyOperationPartTwo(a int, operation string) int {
	ops := strings.Split(operation, " ")
	// first operand is always "old"
	var b int
	if ops[2] != "old" {
		b, _ = strconv.Atoi(ops[2])
	} else {
		b = a
	}
	// operator is only ever * or +
	if ops[1] == "+" {
		return a + b
	} else {
		fmt.Printf("Op: %v returns %v * %v\n", operation, a, b)
		return a * b
	}
}

func partTwo() int {
	monkeys := getMonkeys()
	numRounds := 20
	for i := 0; i < numRounds; i++ {
		for _, monkey := range monkeys {
			for _, item := range monkey.items {
				worryLevel := applyOperationPartTwo(item, monkey.operation)
				fmt.Printf("  div test: %v\n", monkey.divisibleTest)
				if worryLevel%monkey.divisibleTest == 0 { // = "if t is a factor of W"
					// if t is a factor of W, then: W = t * x
					// could we just store t?
					monkeys[monkey.ifTrueTo].items = append(monkeys[monkey.ifTrueTo].items, worryLevel)
				} else {
					// divtest + remainder
					// monkey.divisibleTest+(worryLevel%monkey.divisibleTest)
					monkeys[monkey.ifFalseTo].items = append(monkeys[monkey.ifFalseTo].items, worryLevel)
				}
				monkeys[monkey.number].numInspections++
			}
			monkeys[monkey.number].items = []int{}
		}
		fmt.Printf("After round %v monkey items are:\n", i+1)
		for _, monkey := range monkeys {
			fmt.Printf("Monkey %v: %v\n", monkey.number, monkey.items)
		}
		fmt.Println()
	}
	fmt.Println("Monkey activity:")
	activity := []int{}
	for _, monkey := range monkeys {
		fmt.Printf("Monkey %v: %v\n", monkey.number, monkey.numInspections)
		activity = append(activity, monkey.numInspections)
	}
	sort.Ints(activity)
	return activity[len(activity)-1] * activity[len(activity)-2]
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
