package day_7

import (
	"adventofcode2022/util"
	"fmt"
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

func (dir *Directory) getTotalSize() int {
	size := dir.size
	for _, dir := range dir.subdirectories {
		size += dir.getTotalSize()
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
				//fmt.Printf(" > cd %v\n", to)
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
				//fmt.Printf(" > new current directory is: %v\n", current.name)
			} else {
				// This is a listing, we can populate the current directory with the next lines up until the next $
				// According to the rubric, there is no `ls <dir>`, only `ls` for the current directory.
				//fmt.Printf(" > ls\n")
				i++
				for ; i < len(input); i++ {
					line = input[i]
					//fmt.Printf(" >> line: %v\n", line)
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
	fmt.Printf("Checking: %v\n", dir.name)
	sum := 0
	ts := dir.getTotalSize()
	if ts > 100000 {
		fmt.Printf(" > too big, explore subdirectories\n")
		for _, sub := range dir.subdirectories {
			//fmt.Printf("%v contains %v\n", dir.name, sub.name)
			sum += sub.sumSmallDirectories()
		}
	} else {
		fmt.Printf("add size %v from dir %v\n", ts, dir.name)
		sum += ts
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
