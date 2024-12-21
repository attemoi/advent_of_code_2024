package day08

import (
	"aoc2024/utils"
	"unicode"
)

type Vector = utils.Vector
type Grid = utils.Grid[rune]

func SolvePart1(input string) int {
	grid := utils.ParseGrid(input, parseAntenna)
	antiNodes := findAntiNodesPart1(grid)
	return len(antiNodes)
}

func SolvePart2(input string) int {
	grid := utils.ParseGrid(input, parseAntenna)
	antiNodes := findAntiNodesPart2(grid)
	return len(antiNodes)
}

func parseAntenna(input rune) (rune, bool) {
	if unicode.IsLetter(input) || unicode.IsNumber(input) {
		return input, true
	}
	return 0, false
}

func findAntiNodesPart1(grid *Grid) map[Vector]struct{} {
	antiNodes := make(map[Vector]struct{})
	for _, antennas := range grid.GroupByValue() {
		doForAllPairs(antennas, func(a, b Vector) {
			vectorBetweenAntennas := b.Substract(a)

			antiNode1 := a.Substract(vectorBetweenAntennas)
			if grid.IsInBounds(antiNode1) {
				antiNodes[antiNode1] = struct{}{}
			}

			antiNode2 := b.Add(vectorBetweenAntennas)
			if grid.IsInBounds(antiNode2) {
				antiNodes[antiNode2] = struct{}{}
			}

		})
	}
	return antiNodes
}

func findAntiNodesPart2(grid *Grid) map[Vector]struct{} {
	antiNodes := make(map[Vector]struct{})
	for _, antennas := range grid.GroupByValue() {
		doForAllPairs(antennas, func(a, b Vector) {
			vectorBetweenAntennas := b.Substract(a)

			currentPos := a
			for grid.IsInBounds(currentPos) {
				antiNodes[currentPos] = struct{}{}
				currentPos = currentPos.Substract(vectorBetweenAntennas)
			}

			currentPos = b
			for grid.IsInBounds(currentPos) {
				antiNodes[currentPos] = struct{}{}
				currentPos = currentPos.Add(vectorBetweenAntennas)
			}

		})
	}
	return antiNodes
}

// run function for each possible pair of two elements within a slice
func doForAllPairs[T any](slice []T, process func(a, b T)) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			process(slice[i], slice[j])
		}
	}
}
