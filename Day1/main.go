package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	file, err := os.Open("Day1/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var currentCal int
	var top3 [3]int
	for scanner.Scan() {
		var cal int
		fmt.Sscanf(scanner.Text(), "%d", &cal)
		currentCal += cal
		if scanner.Text() == "" {
			sort.Ints(top3[:])
			if currentCal > top3[0] {
				top3[0] = currentCal
			}
			currentCal = 0
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	var sum int
	for _, c := range top3 {
		sum += c
	}
	fmt.Println(top3[2])
	fmt.Println(sum)
}
