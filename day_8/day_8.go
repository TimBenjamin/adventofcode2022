package day_8

import (
	"adventofcode2022/util"
	"strconv"
	"strings"
)

var input []string

func logTree(visibleTrees map[string]struct{}, row int, col int, height int) map[string]struct{} {
	ref := strconv.Itoa(row) + "-" + strconv.Itoa(col) + "=" + strconv.Itoa(height)
	visibleTrees[ref] = struct{}{}
	return visibleTrees
}

func getViews() [][]int {
	// parse the input into a more useful structure
	views := [][]int{}
	for _, line := range input {
		sp := strings.Split(line, "")
		view := []int{}
		for _, s := range sp {
			tree, _ := strconv.Atoi(s)
			view = append(view, tree)
		}
		views = append(views, view)
	}
	return views
}

func partOne() int {
	views := getViews()

	// We should not double count trees, so the count will be in a set
	visibleTrees := map[string]struct{}{}

	for row, view := range views {
		// We have to consider each view from both directions.
		// forwards:
		min := -1
		for col, height := range view {
			if height > min {
				visibleTrees = logTree(visibleTrees, row, col, height)
				min = height
			}
		}
		// and now backwards:
		min = -1
		for col := len(view) - 1; col >= 0; col-- {
			if view[col] > min {
				visibleTrees = logTree(visibleTrees, row, col, view[col])
				min = view[col]
			}
		}
	}
	// and now the columns:
	for col := 0; col < len(views); col++ {
		// forwards:
		min := -1
		for row := 0; row < len(views[col]); row++ {
			if views[row][col] > min {
				visibleTrees = logTree(visibleTrees, row, col, views[row][col])
				min = views[row][col]
			}
		}
		// backwards:
		min = -1
		for row := len(views[col]) - 1; row >= 0; row-- {
			if views[row][col] > min {
				visibleTrees = logTree(visibleTrees, row, col, views[row][col])
				min = views[row][col]
			}
		}
	}
	return len(visibleTrees)
}

func getViewDistance(direction string, views [][]int, row int, col int) int {
	viewDistance := 0
	if direction == "right" {
		for c := col + 1; c < len(views[row]); c++ {
			viewDistance++
			if views[row][c] >= views[row][col] {
				break
			}
		}
	} else if direction == "left" {
		for c := col - 1; c >= 0; c-- {
			viewDistance++
			if views[row][c] >= views[row][col] {
				break
			}
		}
	} else if direction == "down" {
		for r := row + 1; r < len(views); r++ {
			viewDistance++
			if views[r][col] >= views[row][col] {
				break
			}
		}
	} else if direction == "up" {
		for r := row - 1; r >= 0; r-- {
			viewDistance++
			if views[r][col] >= views[row][col] {
				break
			}
		}
	}
	return viewDistance
}

func partTwo() int {
	views := getViews()
	highScore := 0
	for row, view := range views {
		for col := range view {
			// for this position (row,col) how far can I see up, down, left, and right; calculate a score
			score := getViewDistance("up", views, row, col) * getViewDistance("down", views, row, col) * getViewDistance("left", views, row, col) * getViewDistance("right", views, row, col)
			if score > highScore {
				highScore = score
			}
		}
	}
	return highScore
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
