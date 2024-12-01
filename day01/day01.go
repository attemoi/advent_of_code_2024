package day01

import (
	"sort"
	"strconv"
	"strings"
)

func SolvePart1(input string) int {
	left, right := parseLists(input)
	sort.Ints(left)
	sort.Ints(right)

	return sumOfDiffs(left, right)
}

func sumOfDiffs(left []int, right []int) int {
	sum := 0
	for i, leftNum := range left {
		sum += Abs(leftNum - right[i])
	}
	return sum
}

func SolvePart2(input string) int {
	left, right := parseLists(input)
	sort.Ints(left)
	sort.Ints(right)

	return similarityScore(left, right)
}

func similarityScore(left []int, right []int) int {
	score := 0
	for _, num := range left {
		score += num * countOccurrences(right, num)
	}
	return score
}

func countOccurrences(slice []int, target int) int {
	count := 0
	for _, value := range slice {
		if value == target {
			count++
		}
	}
	return count
}

func parseLists(input string) (left []int, right []int) {
	inputLines := strings.Split(input, "\n")

	for _, line := range inputLines {
		if IsBlank(line) {
			continue
		}
		fields := strings.Fields(line)
		// Trust the input, ignore errors
		num1, _ := strconv.Atoi(fields[0])
		num2, _ := strconv.Atoi(fields[1])
		left = append(left, num1)
		right = append(right, num2)
	}

	return left, right
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func IsBlank(s string) bool {
	return strings.TrimSpace(s) == ""
}
