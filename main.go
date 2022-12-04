package main

import (
	"AOC2022/day4"
	"bufio"
	"fmt"
	"os"
)

const day = 4

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

	day4.Part2(input)
}
