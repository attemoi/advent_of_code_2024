package day12

import (
	"strconv"
	"strings"
)

func SolvePart1(input string) int {
	inputLines := strings.Split(input, "\n")
	return sum(inputLines)
}

func sum(values []string) int {
	sum := 0
	for _, value := range values {
		if number, err := strconv.Atoi(value); err == nil {
			sum += number
		}
	}
	return sum
}

func SolvePart2(input string) int {
	// Implement part 2 logic here
	return 0
}
