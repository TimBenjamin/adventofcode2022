package day_12

import (
	"adventofcode2022/util"
	"fmt"
	"strconv"
	"strings"
)

var input []string

type Coord struct {
	x           int
	y           int
	energyLevel int
	key         string
}

func makeKey(x int, y int) string {
	return fmt.Sprint(x, "-", y)
}

func getGrid() (map[string]Coord, Coord, Coord) {
	grid := map[string]Coord{}
	start := Coord{}
	end := Coord{}
	for y, line := range input {
		sp := strings.Split(line, "")
		for x, c := range sp {
			coord := Coord{x: x, y: y, energyLevel: int(c[0]), key: makeKey(x, y)}
			if c == "S" {
				coord.energyLevel = int('a')
				start = coord
			}
			if c == "E" {
				coord.energyLevel = int('z')
				end = coord
			}
			grid[coord.key] = coord
		}
	}
	return grid, start, end
}

func getNavigableCoords(coord Coord, grid map[string]Coord, visited map[string]struct{}) []Coord {
	navigable := []Coord{}
	// down:
	testKey := makeKey(coord.x, coord.y+1)
	if _, coordExists := grid[testKey]; coordExists && grid[testKey].energyLevel <= coord.energyLevel+1 {
		if _, visitedCoord := visited[testKey]; !visitedCoord {
			navigable = append(navigable, grid[testKey])
		}
	}
	// right:
	testKey = makeKey(coord.x+1, coord.y)
	if _, coordExists := grid[testKey]; coordExists && grid[testKey].energyLevel <= coord.energyLevel+1 {
		if _, visitedCoord := visited[testKey]; !visitedCoord {
			navigable = append(navigable, grid[testKey])
		}
	}
	// up:
	testKey = makeKey(coord.x, coord.y-1)
	if _, coordExists := grid[testKey]; coordExists && grid[testKey].energyLevel <= coord.energyLevel+1 {
		if _, visitedCoord := visited[testKey]; !visitedCoord {
			navigable = append(navigable, grid[testKey])
		}
	}
	// left:
	testKey = makeKey(coord.x-1, coord.y)
	if _, coordExists := grid[testKey]; coordExists && grid[testKey].energyLevel <= coord.energyLevel+1 {
		if _, visitedCoord := visited[testKey]; !visitedCoord {
			navigable = append(navigable, grid[testKey])
		}
	}
	return navigable
}

func partOne() int {
	// turn the input into a slice of Coords, with numbers instead of letters for energy level
	grid, start, end := getGrid()
	fmt.Printf("Start: %v / End: %v\n", start, end)
	// for _, coord := range grid {
	// 	fmt.Println(coord)
	// }
	// curX := start.x
	// curY := start.y
	curEnergy := start.energyLevel
	fmt.Printf("current Energy: %v\n", curEnergy)

	// TODO: a BFS might work

	// need a place to store nodes that we have visited, will use Coord.key as the key
	visited := map[string]struct{}{}
	fmt.Printf("Visited how many: %v\n", len(visited))

	// need a helper function to get all the navigable Coords (NESW) from my current Coord
	navigableCoords := getNavigableCoords(start, grid, visited)
	for _, c := range navigableCoords {
		fmt.Printf("navigable coord: %v\n", c)
	}

	return 0
}

func partTwo() int {
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
