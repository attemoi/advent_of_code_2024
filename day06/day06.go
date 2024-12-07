package day06

import (
	"slices"
	"strings"
)

// VECTOR

type vector struct {
	x, y int
}

func (v vector) Rotate90DegreesClockwise() vector {
	return vector{
		x: -v.y,
		y: v.x,
	}
}

// GUARD

type guard struct {
	position  vector
	direction vector
}

func (g guard) move() guard {
	newGuard := g
	newGuard.position.x += g.direction.x
	newGuard.position.y += g.direction.y
	return newGuard
}

func (g guard) turnRight() guard {
	newGuard := g
	newGuard.direction = g.direction.Rotate90DegreesClockwise()
	return newGuard
}

// MAP

type areaMap struct {
	obstructions map[vector]struct{}
	startPos     vector
	width        int
	height       int
}

func newAreaMap() *areaMap {
	return &areaMap{
		obstructions: make(map[vector]struct{}),
	}
}

func (am *areaMap) addObstruction(x int, y int) {
	am.obstructions[vector{x: x, y: y}] = struct{}{}
}

func (am *areaMap) deleteObstruction(x int, y int) {
	delete(am.obstructions, vector{x: x, y: y})
}

func (am *areaMap) isInBounds(position vector) bool {
	return position.x >= 0 &&
		position.x < am.width &&
		position.y >= 0 &&
		position.y < am.height
}

func (am *areaMap) isObstructionAt(position vector) bool {
	_, exists := am.obstructions[position]
	return exists
}

// SOLUTION

func SolvePart1(input string) int {
	areaMap := parseInput(input)
	up := vector{x: 0, y: -1}
	guard := guard{position: areaMap.startPos, direction: up}
	log, _ := simulateGuard(guard, areaMap)
	return countDistinctGuardPositions(log)
}

func SolvePart2(input string) int {
	areaMap := parseInput(input)
	up := vector{x: 0, y: -1}
	guard := guard{position: areaMap.startPos, direction: up}
	log, _ := simulateGuard(guard, areaMap)
	return countTimeParadoxOptions(log, areaMap)
}

func countDistinctGuardPositions(log []guard) int {
	positions := make(map[vector]struct{})
	for _, guard := range log {
		positions[guard.position] = struct{}{}
	}
	return len(positions)
}

func countTimeParadoxOptions(log []guard, obstructions areaMap) int {
	// Brute force ¯\_(ツ)_/¯
	possiblePositions := make(map[vector]struct{})
	for i := 0; i < len(log)-1; i++ {
		nextPos := log[i+1].position
		_, alreadyTested := possiblePositions[nextPos]
		if nextPos == log[0].position || alreadyTested {
			continue
		}
		obstructions.addObstruction(nextPos.x, nextPos.y)
		_, isLoop := simulateGuard(log[0], obstructions)
		obstructions.deleteObstruction(nextPos.x, nextPos.y)
		if isLoop {
			possiblePositions[vector{x: nextPos.x, y: nextPos.y}] = struct{}{}
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
				areaMap.startPos = vector{x: x, y: y}
			} else if rune == '#' {
				areaMap.addObstruction(x, y)
			}
		}
	}
	areaMap.width = len(lines[0])
	areaMap.height = len(lines)
	return areaMap
}
