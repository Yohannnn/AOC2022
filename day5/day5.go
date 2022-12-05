package day5

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input []string) {
	//Variables
	var crates [][]string
	var cratesString []string
	isMoving := false

	//Loop through input
	for _, l := range input {
		// Parse the crates
		if !isMoving {
			if l == "" {
				columRow := strings.ReplaceAll(cratesString[len(cratesString)-1], " ", "")
				columnCount, _ := strconv.Atoi(string(columRow[len(columRow)-1]))
				cratesString = cratesString[:len(cratesString)-1]
				crates = make([][]string, columnCount)
				// Parse information from the rows
				for _, r := range cratesString {
					row := strings.Split(r, " ")
					for c := 0; c < len(row); c++ {
						// Parse empty space
						if row[c] == "" && row[c+1] == "" {
							row = append(row[:c], row[c+3:]...)
							continue
						}

						// Parse crate
						if row[c][0] == '[' {
							crates[c] = append(crates[c], string(row[c][1]))
						}
					}
				}
				isMoving = true
				continue
			}
			cratesString = append(cratesString, l)
		}

		// Move the crates
		if isMoving {
			// Parse the instruction
			instruction := strings.Split(l, " ")
			count, _ := strconv.Atoi(instruction[1])
			from, _ := strconv.Atoi(instruction[3])
			from--
			to, _ := strconv.Atoi(instruction[5])
			to--

			// Make copy of crates from
			cratesFrom := make([]string, len(crates[from]))
			copy(cratesFrom, crates[from])

			// Get the crates to move and reverse their order
			cratesMoving := cratesFrom[:count]
			for i, j := 0, len(cratesMoving)-1; i < j; i, j = i+1, j-1 {
				cratesMoving[i], cratesMoving[j] = cratesMoving[j], cratesMoving[i]
			}

			// Execute the instruction
			crates[to] = append(cratesMoving, crates[to]...)
			crates[from] = crates[from][count:]
		}
	}

	// Get the top crate of each column
	for c := 0; c < len(crates); c++ {
		fmt.Print(crates[c][0])
	}
	fmt.Println()
}

func Part2(input []string) {
	//Variables
	var crates [][]string
	var cratesString []string
	isMoving := false

	//Loop through input
	for _, l := range input {
		// Parse the crates
		if !isMoving {
			if l == "" {
				columRow := strings.ReplaceAll(cratesString[len(cratesString)-1], " ", "")
				columnCount, _ := strconv.Atoi(string(columRow[len(columRow)-1]))
				cratesString = cratesString[:len(cratesString)-1]
				crates = make([][]string, columnCount)
				// Parse information from the rows
				for _, r := range cratesString {
					row := strings.Split(r, " ")
					for c := 0; c < len(row); c++ {
						// Parse empty space
						if row[c] == "" && row[c+1] == "" {
							row = append(row[:c], row[c+3:]...)
							continue
						}

						// Parse crate
						if row[c][0] == '[' {
							crates[c] = append(crates[c], string(row[c][1]))
						}
					}
				}
				isMoving = true
				continue
			}
			cratesString = append(cratesString, l)
		}

		// Move the crates
		if isMoving {
			// Parse the instruction
			instruction := strings.Split(l, " ")
			count, _ := strconv.Atoi(instruction[1])
			from, _ := strconv.Atoi(instruction[3])
			from--
			to, _ := strconv.Atoi(instruction[5])
			to--

			// Make copy of crates from
			cratesFrom := make([]string, len(crates[from]))
			copy(cratesFrom, crates[from])

			// Execute the instruction
			crates[to] = append(cratesFrom[:count], crates[to]...)
			crates[from] = crates[from][count:]
		}
	}

	// Get the top crate of each column
	for c := 0; c < len(crates); c++ {
		fmt.Print(crates[c][0])
	}
	fmt.Println()
}
