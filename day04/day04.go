package day04

import (
	"strings"
)

type vector struct {
	x, y int
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
	allDirections := []vector{
		{x: 0, y: -1},  // NORTH
		{x: 1, y: -1},  // NORTH EAST
		{x: 1, y: 0},   // EAST
		{x: 1, y: 1},   // SOUTH EAST
		{x: 0, y: 1},   // SOUTH
		{x: -1, y: 1},  // SOUTH WEST
		{x: -1, y: 0},  // WEST
		{x: -1, y: -1}, // NORTH WEST
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
			southEast := vector{x: 1, y: 1}
			southWest := vector{x: -1, y: 1}

			if (isWordFound(x, y, southEast, "SAM", runeGrid) || isWordFound(x, y, southEast, "MAS", runeGrid)) &&
				(isWordFound(x+2, y, southWest, "SAM", runeGrid) || isWordFound(x+2, y, southWest, "MAS", runeGrid)) {
				sum++
			}

		}
	}
	return sum
}

func isWordFound(col int, row int, direction vector, word string, runeGrid [][]rune) bool {
	if isOutOfBounds(col, row, direction, word, runeGrid) {
		return false
	}

	for i, rune := range word {
		currentCol := col + i*direction.x
		currentRow := row + i*direction.y
		if rune != runeGrid[currentRow][currentCol] {
			return false
		}
	}

	return true
}

func isOutOfBounds(col int, row int, direction vector, word string, runeGrid [][]rune) bool {
	numRows := len(runeGrid)
	numColumns := len(runeGrid[0])
	endCol := col + direction.x*(len(word)-1)
	endRow := row + direction.y*(len(word)-1)
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
