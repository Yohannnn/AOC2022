package main

import (
	"AOC2022/day5"
	"bufio"
	"fmt"
	"os"
)

const day = 5

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

	day5.Part2(input)
}
