package day04

import (
	"strings"
)

type Vector struct {
	X, Y int
}

func SolvePart1(input string) int {
	runeGrid := parseGrid(input)
	return countXmasInstances(runeGrid)
}

func SolvePart2(input string) int {
	runeGrid := parseGrid(input)
	return countMasCrossInstances(runeGrid)
}

func countXmasInstances(runeGrid [][]rune) int {
	allDirections := []Vector{
		{X: 0, Y: -1},  // NORTH
		{X: 1, Y: -1},  // NORTH EAST
		{X: 1, Y: 0},   // EAST
		{X: 1, Y: 1},   // SOUTH EAST
		{X: 0, Y: 1},   // SOUTH
		{X: -1, Y: 1},  // SOUTH WEST
		{X: -1, Y: 0},  // WEST
		{X: -1, Y: -1}, // NORTH WEST
	}

	sum := 0
	for y := range runeGrid {
		for x := range runeGrid[y] {
			for _, dir := range allDirections {
				if isWordFound(x, y, dir, "XMAS", runeGrid) {
					sum++
				}
			}
		}
	}
	return sum
}

func countMasCrossInstances(runeGrid [][]rune) int {
	sum := 0
	for y := range runeGrid {
		for x := range runeGrid[y] {
			southEast := Vector{X: 1, Y: 1}
			southWest := Vector{X: -1, Y: 1}

			if (isWordFound(x, y, southEast, "SAM", runeGrid) || isWordFound(x, y, southEast, "MAS", runeGrid)) &&
				(isWordFound(x+2, y, southWest, "SAM", runeGrid) || isWordFound(x+2, y, southWest, "MAS", runeGrid)) {
				sum++
			}

		}
	}
	return sum
}

func isWordFound(col int, row int, direction Vector, word string, runeGrid [][]rune) bool {
	if isOutOfBounds(col, row, direction, word, runeGrid) {
		return false
	}

	for i, rune := range word {
		currentCol := col + i*direction.X
		currentRow := row + i*direction.Y
		if rune != runeGrid[currentRow][currentCol] {
			return false
		}
	}

	return true
}

func isOutOfBounds(col int, row int, direction Vector, word string, runeGrid [][]rune) bool {
	numRows := len(runeGrid)
	numColumns := len(runeGrid[0])
	endCol := col + direction.X*(len(word)-1)
	endRow := row + direction.Y*(len(word)-1)
	return col < 0 ||
		col > (numColumns-1) ||
		row < 0 ||
		row > (numRows-1) ||
		endCol < 0 ||
		endCol > (numColumns-1) ||
		endRow < 0 ||
		endRow > (numRows-1)
}

func parseGrid(input string) [][]rune {
	var grid [][]rune
	for _, line := range strings.Split(input, "\n") {
		grid = append(grid, []rune(line))
	}
	return grid
}
