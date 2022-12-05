package day5

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input []string) {
	//Variables
	var crates [][]string
	isMoving := false

	//Loop through input
	for _, l := range input {
		//Skip over separator line
		if l == "" {
			continue
		}

		//Check if all crates have been parsed
		if l[1] == '1' {
			isMoving = true
			continue
		}

		// Parse the crates
		if !isMoving {
			row := strings.Split(l, " ")
			for i := 0; i < len(row); i++ {
				// Parse empty space
				if row[i] == "" && row[i+1] == "" {
					row = append(row[:i], row[i+3:]...)
					continue
				}
				// Parse crates
				if row[i][0] == '[' {
					row[i] = string(row[i][1])
				}
			}
			crates = append(crates, row)
		}

		// Move the crates
		if isMoving {
			// Parse the instruction
			instruction := strings.Split(l, " ")
			count, _ := strconv.Atoi(instruction[1])
			from, _ := strconv.Atoi(instruction[3])
			to, _ := strconv.Atoi(instruction[5])

			// Execute the instruction
			for move := 0; move < count; move++ {
			instuct:
				for fromRow := 0; fromRow < len(crates); fromRow++ {
					crate := crates[fromRow][from-1]
					if crate != "" {
						// Check if a new row needs to be created
						if crates[0][to-1] != "" {
							// Generates new blank row
							newRow := make([][]string, 1)
							newRow[0] = make([]string, len(crates[0]))

							crates = append(newRow, crates...)
							fromRow++
						}

						for toRow := 0; toRow < len(crates); toRow++ {
							// Check if at bottom row
							if toRow == len(crates)-1 {
								crates[toRow][to-1] = crate
								crates[fromRow][from-1] = ""
								break instuct
							}
							if crates[toRow+1][to-1] != "" {
								crates[toRow][to-1] = crate
								crates[fromRow][from-1] = ""
								break instuct
							}
						}
					}
				}
			}
		}
	}

	// Get the top crate of each column
	topCrates := make([]string, len(crates[0]))

	for c := 0; c < len(crates[0]); c++ {
		for r := 0; r < len(crates); r++ {
			if crates[r][c] != "" && topCrates[c] == "" {
				topCrates[c] = crates[r][c]
				break
			}
		}
	}

	// Print the top crates
	for _, crate := range topCrates {
		fmt.Print(crate)
	}
}

func Part2(input []string) {
	//Variables
	var crates [][]string
	isMoving := false

	//Loop through input
	for _, l := range input {
		//Skip over separator line
		if l == "" {
			continue
		}

		//Check if all crates have been parsed
		if l[1] == '1' {
			isMoving = true
			continue
		}

		// Parse the crates
		if !isMoving {
			row := strings.Split(l, " ")
			for i := 0; i < len(row); i++ {
				// Parse empty space
				if row[i] == "" && row[i+1] == "" {
					row = append(row[:i], row[i+3:]...)
					continue
				}
				// Parse crates
				if row[i][0] == '[' {
					row[i] = string(row[i][1])
				}
			}
			crates = append(crates, row)
		}

		// Move the crates
		if isMoving {
			// Parse the instruction
			instruction := strings.Split(l, " ")
			count, _ := strconv.Atoi(instruction[1])
			from, _ := strconv.Atoi(instruction[3])
			to, _ := strconv.Atoi(instruction[5])

			var moveCrates []string
			//Get the crates to move and remove them from the crates slice
			for fromRow := 0; fromRow < len(crates); fromRow++ {
				if crates[fromRow][from-1] != "" {
					moveCrates = make([]string, count)
					for i := 0; i < count; i++ {
						moveCrates[i] = crates[fromRow+i][from-1]
						crates[fromRow+i][from-1] = ""
					}
					break
				}
			}

			//Move the crates
			for bottom := 0; bottom < len(crates); bottom++ {
				if crates[bottom][to-1] != "" || bottom == len(crates)-1 {

					//Increment bottom if at bottom row and its empty
					if bottom == len(crates)-1 && crates[bottom][to-1] == "" {
						bottom++
					}

					//Check if there is enough space to move the crates
					if len(crates)-bottom+count > len(crates) {
						//Generate new blank rows
						neededRows := count + len(crates) - bottom
						newRows := make([][]string, neededRows)
						for i := 0; i < neededRows; i++ {
							newRows[i] = make([]string, len(crates[0]))
						}
						crates = append(newRows, crates...)
						bottom += neededRows
					}

					//Move the crates
					for i := 0; i < count; i++ {
						crates[bottom-i-1][to-1] = moveCrates[count-i-1]
					}
					break
				}
			}
		}
		fmt.Println(crates)
	}

	// Get the top crate of each column
	topCrates := make([]string, len(crates[0]))

	for c := 0; c < len(crates[0]); c++ {
		for r := 0; r < len(crates); r++ {
			if crates[r][c] != "" && topCrates[c] == "" {
				topCrates[c] = crates[r][c]
				break
			}
		}
	}

	// Print the top crates
	for _, crate := range topCrates {
		fmt.Print(crate)
	}
}
