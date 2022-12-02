package day1

import (
	"fmt"
	"sort"
	"strconv"
)

func Run(input []string) {
	var currentCal int
	var top3 [3]int
	for _, l := range input {
		cal, _ := strconv.Atoi(l)
		currentCal += cal
		if l == "" {
			sort.Ints(top3[:])
			if currentCal > top3[0] {
				top3[0] = currentCal
			}
			currentCal = 0
		}
	}

	var sum int
	for _, c := range top3 {
		sum += c
	}
	fmt.Println(top3[2])
	fmt.Println(sum)
}
