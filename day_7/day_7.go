package day_7

import (
	"adventofcode2022/util"
	"log"
	"strconv"
	"strings"
)

var input []string

type Directory struct {
	name           string
	size           int
	parent         *Directory
	files          map[string]int        // filename => size
	subdirectories map[string]*Directory // dirname => Directory
}

var root *Directory

func (dir *Directory) addSize(size int) {
	dir.size += size
	if dir.parent != nil {
		dir.parent.addSize(size)
	}
}

func (dir *Directory) getBranchSize() int {
	size := dir.size
	for _, dir := range dir.subdirectories {
		size += dir.getBranchSize()
	}
	return size
}

func makeFilesystem() {
	root = &Directory{
		name:           "/",
		size:           0,
		parent:         nil,
		files:          map[string]int{},
		subdirectories: map[string]*Directory{},
	}
	current := root
	for i := 0; i < len(input); i++ {
		line := input[i]
		if line[0:1] == "$" {
			// either cd or ls
			if line[2:4] == "cd" {
				to := line[5:]
				if to == "/" {
					current = root
				} else if to == ".." {
					if current.parent == nil {
						log.Fatal("Cannot go up from root directory")
					}
					current = current.parent
				} else {
					for dirName, dir := range current.subdirectories {
						if dirName == to {
							current = dir
						}
					}
				}
			} else {
				// This is a listing, we can populate the current directory with the next lines up until the next $
				// According to the rubric, there is no `ls <dir>`, only `ls` for the current directory.
				i++
				for ; i < len(input); i++ {
					line = input[i]
					if line[0:1] == "$" {
						i--
						break
					}
					contents := strings.Split(line, " ")
					if contents[0] == "dir" {
						dirName := contents[1]
						if _, ok := current.subdirectories[contents[0]]; !ok {
							// directory not yet logged, add it to tree
							newDir := &Directory{
								name:           dirName,
								size:           0,
								parent:         current,
								files:          map[string]int{},
								subdirectories: map[string]*Directory{},
							}
							current.subdirectories[newDir.name] = newDir
						}
					} else {
						// it's a file, add it to the current directory's file map
						fileSize, _ := strconv.Atoi(contents[0])
						current.size += fileSize
						if current.parent != nil {
							current.parent.addSize(fileSize)
						}
						fileName := contents[1]
						if _, ok := current.files[fileName]; !ok {
							current.files[fileName] = fileSize
						}
					}
				}
			}
		}
	}
}

func (dir *Directory) sumSmallDirectories() int {
	sum := 0
	if dir.size > 100000 {
		for _, sub := range dir.subdirectories {
			sum += sub.sumSmallDirectories()
		}
	} else {
		sum += dir.getBranchSize()
	}
	return sum
}

func partOne() int {
	makeFilesystem()
	// find all of the directories with a total size of AT MOST 100000
	// then calculate the sum of their total sizes
	// files can be counted more than once
	return root.sumSmallDirectories()
}

func (dir *Directory) findClosest(neededSpace int) int {
	if dir.size >= neededSpace {
		if dir.size < currentClosest {
			currentClosest = dir.size
		}
		for _, subdir := range dir.subdirectories {
			subdir.findClosest(neededSpace)
		}
	}
	return currentClosest
}

var currentClosest int

func partTwo() int {
	makeFilesystem()
	diskSpace := 70000000
	updateSize := 30000000
	unusedSpace := diskSpace - root.size
	neededSpace := updateSize - unusedSpace
	// Find the smallest directory that, if deleted, would free up enough space on the filesystem to run the update.
	// What is the total size of that directory?
	currentClosest = root.size
	return root.findClosest(neededSpace)
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
