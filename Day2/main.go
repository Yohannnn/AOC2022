package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("Day2/input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var Score int
	var moves [2]string
	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		switch input[0] {
		case "A":
			switch input[1] {
			case "X":
				moves = [2]string{"A", "Z"}
			case "Y":
				moves = [2]string{"A", "X"}
			case "Z":
				moves = [2]string{"A", "Y"}
			}
		case "B":
			switch input[1] {
			case "X":
				moves = [2]string{"B", "X"}
			case "Y":
				moves = [2]string{"B", "Y"}
			case "Z":
				moves = [2]string{"B", "Z"}
			}
		case "C":
			switch input[1] {
			case "X":
				moves = [2]string{"C", "Y"}
			case "Y":
				moves = [2]string{"C", "Z"}
			case "Z":
				moves = [2]string{"C", "X"}
			}

		}

		Score += calcScore(moves)
	}
	fmt.Println(Score)
}

func calcScore(moves [2]string) (scoreIncrease int) {
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
