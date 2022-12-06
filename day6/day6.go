package day6

import (
	"fmt"
	"strings"
)

func Part1(input []string) {
	buffer := input[0]
bufferLoop:
	for i, _ := range buffer {
		for p := 1; p <= 4; p++ {
			if strings.Contains(buffer[i+p:i+4], string(buffer[i+p-1])) {
				continue bufferLoop
			}
		}
		fmt.Println(i + 4)
		break
	}
}

func Part2(input []string) {
	buffer := input[0]
bufferLoop:
	for i, _ := range buffer {
		for p := 1; p <= 14; p++ {
			if strings.Contains(buffer[i+p:i+14], string(buffer[i+p-1])) {
				continue bufferLoop
			}
		}
		fmt.Println(i + 14)
		break
	}
}
