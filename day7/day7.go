package day7

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type dir struct {
	name     string
	parent   *dir
	children map[string]*dir
	files    []int
}

func checkSize(dir *dir, sum *int) int {
	size := 0
	for _, f := range dir.files {
		size += f
	}
	for _, d := range dir.children {
		size += checkSize(d, sum)
	}
	if size <= 100000 {
		*sum += size
	}

	return size
}

func getSize(dir *dir, sizes *[]int) int {
	size := 0
	for _, f := range dir.files {
		size += f
	}
	for _, d := range dir.children {
		size += getSize(d, sizes)
	}

	*sizes = append(*sizes, size)
	return size
}

func ParseCommands(input []string) (root dir) {
	//Variables
	var cd *dir
	root = dir{name: "root", children: make(map[string]*dir)}

	//Loop through input
	for _, l := range input {
		//Split line into parts
		parts := strings.Split(l, " ")

		// Switch on what type of line it is
		switch parts[0] {

		case "$":
			switch parts[1] {

			case "cd":
				switch parts[2] {
				case "/":
					cd = &root
				case "..":
					cd = cd.parent
				default:
					cd = cd.children[parts[2]]
				}
			case "ls":
				continue
			}

		case "dir":
			// Adds the directory as a child of the current directory
			cd.children[parts[1]] = &dir{name: parts[1], parent: cd, children: make(map[string]*dir)}

		default:
			size, _ := strconv.Atoi(parts[0])
			cd.files = append(cd.files, size)
		}
	}
	return root
}

func Part1(input []string) {
	//Variables
	var sum int

	//Parse commands
	root := ParseCommands(input)

	//Calculate size of each directory
	checkSize(&root, &sum)
	fmt.Println(sum)
}

func Part2(input []string) {
	//Parse commands
	root := ParseCommands(input)

	//Calculate size of each directory and then sort it
	var sizes []int
	rootSize := getSize(&root, &sizes)
	sort.Ints(sizes)

	//Find best size to delete
	neededSpace := 30000000 - (70000000 - rootSize)

	for _, s := range sizes {
		if neededSpace-s <= 0 {
			fmt.Println(s)
			break
		}
	}
}
