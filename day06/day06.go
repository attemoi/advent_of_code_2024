package day06

import (
	"aoc2024/utils"
	"slices"
	"strings"
)

type Vector = utils.Vector

// GUARD

type guard struct {
	position  Vector
	direction Vector
}

func (g guard) move() guard {
	newGuard := g
	newGuard.position.X += g.direction.X
	newGuard.position.Y += g.direction.Y
	return newGuard
}

func (g guard) turnRight() guard {
	newGuard := g
	newGuard.direction = g.direction.Rotate90DegreesClockwise()
	return newGuard
}

// MAP

type areaMap struct {
	obstructions map[Vector]struct{}
	startPos     Vector
	width        int
	height       int
}

func newAreaMap() *areaMap {
	return &areaMap{
		obstructions: make(map[Vector]struct{}),
	}
}

func (am *areaMap) addObstruction(x int, y int) {
	am.obstructions[Vector{X: x, Y: y}] = struct{}{}
}

func (am *areaMap) deleteObstruction(x int, y int) {
	delete(am.obstructions, Vector{X: x, Y: y})
}

func (am *areaMap) isInBounds(position Vector) bool {
	return position.X >= 0 &&
		position.X < am.width &&
		position.Y >= 0 &&
		position.Y < am.height
}

func (am *areaMap) isObstructionAt(position Vector) bool {
	_, exists := am.obstructions[position]
	return exists
}

// SOLUTION

func SolvePart1(input string) int {
	areaMap := parseInput(input)
	up := Vector{X: 0, Y: -1}
	guard := guard{position: areaMap.startPos, direction: up}
	log, _ := simulateGuard(guard, areaMap)
	return countDistinctGuardPositions(log)
}

func SolvePart2(input string) int {
	areaMap := parseInput(input)
	up := Vector{X: 0, Y: -1}
	guard := guard{position: areaMap.startPos, direction: up}
	log, _ := simulateGuard(guard, areaMap)
	return countTimeParadoxOptions(log, areaMap)
}

func countDistinctGuardPositions(log []guard) int {
	positions := make(map[Vector]struct{})
	for _, guard := range log {
		positions[guard.position] = struct{}{}
	}
	return len(positions)
}

func countTimeParadoxOptions(log []guard, obstructions areaMap) int {
	// Brute force ¯\_(ツ)_/¯
	possiblePositions := make(map[Vector]struct{})
	for i := 0; i < len(log)-1; i++ {
		nextPos := log[i+1].position
		_, alreadyTested := possiblePositions[nextPos]
		if nextPos == log[0].position || alreadyTested {
			continue
		}
		obstructions.addObstruction(nextPos.X, nextPos.Y)
		_, isLoop := simulateGuard(log[0], obstructions)
		obstructions.deleteObstruction(nextPos.X, nextPos.Y)
		if isLoop {
			possiblePositions[Vector{X: nextPos.X, Y: nextPos.Y}] = struct{}{}
		}
	}
	return len(possiblePositions)
}

func simulateGuard(guard guard, guardArea areaMap) (log []guard, isLoop bool) {
	for guardArea.isInBounds(guard.position) {
		log = append(log, guard)
		for isGuardFacingObstruction(guard, guardArea) {
			guard = guard.turnRight()
		}
		guard = guard.move()
		if slices.Contains(log, guard) {
			return log, true
		}
	}
	return log, false
}

func isGuardFacingObstruction(guard guard, areaMap areaMap) bool {
	return areaMap.isObstructionAt(guard.move().position)
}

func parseInput(input string) (areaMap areaMap) {
	areaMap = *newAreaMap()
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, rune := range line {
			if rune == '^' {
				areaMap.startPos = Vector{X: x, Y: y}
			} else if rune == '#' {
				areaMap.addObstruction(x, y)
			}
		}
	}
	areaMap.width = len(lines[0])
	areaMap.height = len(lines)
	return areaMap
}
