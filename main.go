package main

import (
	"AOC2022/day2"
	"bufio"
	"fmt"
	"os"
)

const day = 2

var input []string

func main() {
	file, err := os.Open(fmt.Sprintf("day%d/input.txt", day))
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	day2.Part1(input)
	day2.Part2(input)
}
