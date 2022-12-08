package day8

import (
	"fmt"
	"strconv"
)

func getHeight(t uint8) int {
	height, _ := strconv.Atoi(string(t))
	return height
}

func Part1(input []string) {
	//Variables
	var visCount int
	talltop := make([]int, len(input[0])-2)

	// Count all the border trees
	visCount += len(input)*2 + len(input[0])*2 - 4

	// Generate tall top based on first row
	for i := 0; i < len(input[0])-2; i++ {
		talltop[i] = getHeight(input[0][i+1])
	}

	// Loop through the input
	for i, line := range input {
		//Check if line is first or last
		if i == 0 || i == len(input)-1 {
			continue
		}

	tree:
		// Check if the inside trees are visible
		for t := 0; t < len(line)-2; t++ {
			height := getHeight(line[t+1])

			// Check if the tree is visible from the top
			if height > talltop[t] {
				visCount++
				talltop[t] = height
				continue
			}

			// Check if the tree is visible from the left
			for l := t + 1; l > 0; l-- {
				if height <= getHeight(line[l-1]) {
					break
				}

				if l == 1 {
					visCount++
					continue tree
				}
			}

			// Check if the tree is visible from the right
			for r := t + 1; r < len(line)-1; r++ {
				if height <= getHeight(line[r+1]) {
					break
				}

				if r == len(line)-2 {
					visCount++
					continue tree
				}
			}

			// Check if the tree is visible from the bottom
			for b := i + 1; b < len(input); b++ {
				if height <= getHeight(input[b][t+1]) {
					break
				}

				if b == len(input)-1 {
					visCount++
					continue tree
				}
			}
		}

	}

	fmt.Println(visCount)
}

func Part2(input []string) {
	//Variables
	var topScore int

	// Loop through the input
	for i, line := range input {

		// Get the score of every tree
		for t := 0; t < len(line); t++ {
			// Check if tree is an edge tree
			if i == 0 || i == len(input)-1 || t == 0 || t == len(line)-1 {
				continue
			}

			height := getHeight(line[t])
			score := 1

			// Get the score from the top
			for top := i - 1; top >= 0; top-- {
				if height <= getHeight(input[top][t]) {
					score *= i - top
					break
				}

				if top == 0 {
					score *= i
				}
			}

			// Check if the tree is visible from the left
			for l := t - 1; l >= 0; l-- {
				if height <= getHeight(line[l]) {
					score *= t - l
					break
				}

				if l == 0 {
					score *= t
				}
			}

			// Check if the tree is visible from the right
			for r := t + 1; r < len(line); r++ {
				if height <= getHeight(line[r]) {
					score *= r - t
					break
				}

				if r == len(line)-1 {
					score *= len(line) - 1 - t
				}
			}

			// Check if the tree is visible from the bottom
			for b := i + 1; b < len(input); b++ {
				if height <= getHeight(input[b][t]) {
					score *= b - i
					break
				}
				if b == len(input)-1 {
					score *= len(input) - 1 - i
				}
			}

			// Check if the score is the highest
			if score > topScore {
				topScore = score
				fmt.Println(score, i, t)
			}
		}

	}

	fmt.Println(topScore)
}
