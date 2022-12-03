package main

import (
	"AOC2022/day3"
	"bufio"
	"fmt"
	"os"
)

const day = 3

var input []string

func main() {

	// Gets input for current day
	file, err := os.Open(fmt.Sprintf("day%d/input.txt", day))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	day3.Part2(input)
}
