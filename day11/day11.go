package day11

import (
	"aoc2024/utils"
	"strconv"
	"strings"
)

var memo = make(map[[2]int]int)

func SolvePart1(input string) int {
	stones := parseStones(input)
	return totalStonesAfterBlinking(stones, 25)
}

func SolvePart2(input string) int {
	stones := parseStones(input)
	return totalStonesAfterBlinking(stones, 75)
}

func totalStonesAfterBlinking(stones []int, times int) int {
	sum := 0
	for _, stone := range stones {
		sum += blinkTimes(stone, times)
	}
	return sum
}

func blinkTimes(stone int, n int) int {
	if n == 0 {
		return 1
	}

	// Check if the result is already computed
	if val, found := memo[[2]int{stone, n}]; found {
		return val
	}

	sum := 0
	for _, stone := range blink(stone) {
		sum += blinkTimes(stone, n-1)
	}

	// Store the result in the memoization map
	memo[[2]int{stone, n}] = sum
	return sum
}

func blink(stone int) []int {
	if stone == 0 {
		return []int{1}
	}

	left, right := split(stone)
	if len(left) == len(right) {
		leftNum, _ := strconv.Atoi(strings.TrimPrefix(left, "0"))
		rightNum, _ := strconv.Atoi(strings.TrimPrefix(right, "0"))
		return []int{leftNum, rightNum}
	}

	return []int{stone * 2024}
}

func split(stone int) (left string, right string) {
	s := strconv.Itoa(stone)
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func parseStones(input string) []int {
	numbers, _ := utils.ToInts(strings.Fields(input))
	return numbers
}
