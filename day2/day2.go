package day2

import (
	"fmt"
	"strings"
)

func Part1(input []string) {
	var score int

	// Loop through input
	for _, l := range input {
		moves := strings.Split(l, " ")
		score += calcScore(moves)
	}
	fmt.Println(score)
}

func Part2(input []string) {
	// Variables
	var score int
	var moves []string

	// Loop through input
	for _, l := range input {
		command := strings.Split(l, " ")
		switch command[0] {
		case "A":
			switch command[1] {
			case "X":
				moves = []string{"A", "Z"}
			case "Y":
				moves = []string{"A", "X"}
			case "Z":
				moves = []string{"A", "Y"}
			}
		case "B":
			switch command[1] {
			case "X":
				moves = []string{"B", "X"}
			case "Y":
				moves = []string{"B", "Y"}
			case "Z":
				moves = []string{"B", "Z"}
			}
		case "C":
			switch command[1] {
			case "X":
				moves = []string{"C", "Y"}
			case "Y":
				moves = []string{"C", "Z"}
			case "Z":
				moves = []string{"C", "X"}
			}

		}
		score += calcScore(moves)
	}
	fmt.Println(score)
}

func calcScore(moves []string) (scoreIncrease int) {
	switch moves[0] {
	case "A":
		switch moves[1] {
		case "X":
			scoreIncrease += 1
			scoreIncrease += 3
		case "Y":
			scoreIncrease += 2
			scoreIncrease += 6
		case "Z":
			scoreIncrease += 3
		}
	case "B":
		switch moves[1] {
		case "X":
			scoreIncrease += 1
		case "Y":
			scoreIncrease += 2
			scoreIncrease += 3
		case "Z":
			scoreIncrease += 3
			scoreIncrease += 6
		}
	case "C":
		switch moves[1] {
		case "X":
			scoreIncrease += 1
			scoreIncrease += 6
		case "Y":
			scoreIncrease += 2
		case "Z":
			scoreIncrease += 3
			scoreIncrease += 3
		}
	}
	return scoreIncrease
}
