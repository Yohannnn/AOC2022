package main

import (
	"AOC2022/day6"
	"bufio"
	"fmt"
	"os"
)

const day = 6

var input []string

func main() {

	// Gets input for current day
	file, err := os.Open(fmt.Sprintf("day%d/input.txt", day))
	if err != nil {
		fmt.Println(err)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
	}

	day6.Part1(input)
	day6.Part2(input)
}
