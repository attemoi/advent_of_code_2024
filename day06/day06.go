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
	Position  vector
	Direction vector
}

func (g guard) move() guard {
	newGuard := g
	newGuard.Position.x += g.Direction.x
	newGuard.Position.y += g.Direction.y
	return newGuard
}

func (g guard) TurnRight() guard {
	newGuard := g
	newGuard.Direction = g.Direction.Rotate90DegreesClockwise()
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

func (am *areaMap) IsObstructionAt(position vector) bool {
	_, exists := am.obstructions[position]
	return exists
}

// SOLUTION

func SolvePart1(input string) int {
	areaMap := parseInput(input)
	up := vector{x: 0, y: -1}
	guard := guard{Position: areaMap.startPos, Direction: up}
	log, _ := simulateGuard(guard, areaMap)
	return countDistinctGuardPositions(log)
}

func SolvePart2(input string) int {
	areaMap := parseInput(input)
	up := vector{x: 0, y: -1}
	guard := guard{Position: areaMap.startPos, Direction: up}
	log, _ := simulateGuard(guard, areaMap)
	return countTimeParadoxOptions(log, areaMap)
}

func countDistinctGuardPositions(log []guard) int {
	positions := make(map[vector]struct{})
	for _, guard := range log {
		positions[guard.Position] = struct{}{}
	}
	return len(positions)
}

func countTimeParadoxOptions(log []guard, obstructions areaMap) int {
	// Brute force ¯\_(ツ)_/¯
	possiblePositions := make(map[vector]struct{})
	for i := 0; i < len(log)-1; i++ {
		nextPos := log[i+1].Position
		_, alreadyTested := possiblePositions[nextPos]
		if nextPos == log[0].Position || alreadyTested {
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
	for guardArea.isInBounds(guard.Position) {
		log = append(log, guard)
		for isGuardFacingObstruction(guard, guardArea) {
			guard = guard.TurnRight()
		}
		guard = guard.move()
		if slices.Contains(log, guard) {
			return log, true
		}
	}
	return log, false
}

func isGuardFacingObstruction(guard guard, areaMap areaMap) bool {
	return areaMap.IsObstructionAt(guard.move().Position)
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
