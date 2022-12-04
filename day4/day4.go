package day4

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(input []string) {
	// Variables
	var invalidRangeCount int

	// Loop through input
	for _, l := range input {
		ranges := strings.Split(l, ",")
		r1 := strings.Split(ranges[0], "-")
		r2 := strings.Split(ranges[1], "-")

		var r1int, r2int [2]int
		for i := 0; i < len(r1); i++ {
			r1int[i], _ = strconv.Atoi(r1[i])
			r2int[i], _ = strconv.Atoi(r2[i])
		}

		// Check if first range is contained in second range
		if r1int[0] >= r2int[0] && r1int[1] <= r2int[1] {
			invalidRangeCount++
			continue
		}

		// Check if second range is contained in first range
		if r2int[0] >= r1int[0] && r2int[1] <= r1int[1] {
			invalidRangeCount++
			continue
		}
	}
	fmt.Println(invalidRangeCount)
}

func Part2(input []string) {
	// Variables
	var overlappingCount int

	// Loop through input
	for _, l := range input {
		ranges := strings.Split(l, ",")
		r1 := strings.Split(ranges[0], "-")
		r2 := strings.Split(ranges[1], "-")

		var r1int, r2int [2]int
		for i := 0; i < len(r1); i++ {
			r1int[i], _ = strconv.Atoi(r1[i])
			r2int[i], _ = strconv.Atoi(r2[i])
		}

		// Generate first range as a slice
		r1slice := make([]int, r1int[1]-r1int[0]+1)
		for i := range r1slice {
			r1slice[i] = i + r1int[0]
		}

		// Generate second range as a slice
		r2slice := make([]int, r2int[1]-r2int[0]+1)
		for i := range r2slice {
			r2slice[i] = i + r2int[0]
		}

		// Check if the two ranges overlap
	mainLoop:
		for _, i := range r1slice {
			for _, j := range r2slice {
				if i == j {
					overlappingCount++
					break mainLoop
				}
			}
		}
	}
	fmt.Println(overlappingCount)
}
