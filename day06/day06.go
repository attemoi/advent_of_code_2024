package day06

import (
	"aoc2024/utils"
	"slices"
	"strings"
)

type Vector = utils.Vector
type GuardArea = utils.Grid[struct{}]

// GUARD

type guard struct {
	position  Vector
	direction Vector
}

// SOLUTION

func SolvePart1(input string) int {
	grid := utils.ParseGrid(input, parseObstruction)
	guard := findGuard(input)
	log, _ := simulateGuard(guard, *grid)
	return countDistinctGuardPositions(log)
}

func SolvePart2(input string) int {
	grid := utils.ParseGrid(input, parseObstruction)
	guard := findGuard(input)
	log, _ := simulateGuard(guard, *grid)
	return countTimeParadoxOptions(log, *grid)
}

func parseObstruction(c rune) (struct{}, bool) {
	if c == '#' {
		return struct{}{}, true
	}
	var zeroValue struct{}
	return zeroValue, false
}

func countDistinctGuardPositions(log []guard) int {
	positions := make(map[Vector]struct{})
	for _, guard := range log {
		positions[guard.position] = struct{}{}
	}
	return len(positions)
}

func countTimeParadoxOptions(log []guard, grid GuardArea) int {
	// Brute force ¯\_(ツ)_/¯
	possiblePositions := make(map[Vector]struct{})
	for i := 0; i < len(log)-1; i++ {
		nextPos := log[i+1].position
		_, alreadyTested := possiblePositions[nextPos]
		if nextPos == log[0].position || alreadyTested {
			continue
		}

		grid.SetZeroVal(nextPos)
		_, isLoop := simulateGuard(log[0], grid)
		grid.Delete(nextPos)
		if isLoop {
			possiblePositions[Vector{X: nextPos.X, Y: nextPos.Y}] = struct{}{}
		}
	}
	return len(possiblePositions)
}

func simulateGuard(guard guard, grid GuardArea) (log []guard, isLoop bool) {
	for grid.IsInBounds(guard.position) {
		log = append(log, guard)
		for isGuardFacingObstruction(guard, grid) {
			guard.direction = guard.direction.Rotate90DegreesClockwise()
		}
		guard.position = guard.position.Add(guard.direction)
		if slices.Contains(log, guard) {
			return log, true
		}
	}
	return log, false
}

func isGuardFacingObstruction(guard guard, grid GuardArea) bool {
	_, isObstructionAtPos := grid.Get(guard.position.Add(guard.direction))
	return isObstructionAtPos
}

func findGuard(input string) guard {
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, rune := range line {
			if rune == '^' {
				return guard{position: Vector{X: x, Y: y}, direction: Vector{X: 0, Y: -1}}
			}
		}
	}
	panic("Start position not found")
}
