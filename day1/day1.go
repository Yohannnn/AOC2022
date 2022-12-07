package day1

import (
	"fmt"
	"sort"
	"strconv"
)

func Part1(input []string) {
	// Variables
	var currentCal int
	var highestCal int

	// Loop through input
	for _, l := range input {
		// Get cal value and add it to the current cal count for this elf
		cal, _ := strconv.Atoi(l)
		currentCal += cal

		// If it's the end of the elf's list, check if the current cal count is higher than the lowest of the top 3
		if l == "" {
			if currentCal > highestCal {
				highestCal = currentCal
			}
			currentCal = 0
		}
	}

	// Print the results
	fmt.Println(highestCal)
}

func Part2(input []string) {
	// Variables
	var currentCal int
	var top3 [3]int
	var sum int

	// Loop through input
	for _, l := range input {
		// Get cal value and add it to the current cal count for this elf
		cal, _ := strconv.Atoi(l)
		currentCal += cal

		// If it's the end of the elf's list, check if the current cal count is higher than the lowest of the top 3
		if l == "" {
			sort.Ints(top3[:])
			if currentCal > top3[0] {
				top3[0] = currentCal
			}
			currentCal = 0
		}
	}

	// Add the top 3 together
	for _, c := range top3 {
		sum += c
	}

	// Print the results
	fmt.Println(sum)
}
