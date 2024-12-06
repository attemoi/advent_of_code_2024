package day06

import (
	"slices"
	"strings"
)

type Vector struct {
	X, Y int
}

func (v Vector) Rotate90DegreesClockwise() Vector {
	return Vector{
		X: -v.Y,
		Y: v.X,
	}
}

type Guard struct {
	Position  Vector
	Direction Vector
}

func (g Guard) Move() Guard {
	newGuard := g
	newGuard.Position.X += g.Direction.X
	newGuard.Position.Y += g.Direction.Y
	return newGuard
}

func (g Guard) TurnRight() Guard {
	newGuard := g
	newGuard.Direction = g.Direction.Rotate90DegreesClockwise()
	return newGuard
}

func SolvePart1(input string) int {
	obstructions, guardPos := parseInput(input)
	up := Vector{X: 0, Y: -1}
	guard := Guard{Position: guardPos, Direction: up}
	log, _ := simulateGuard(guard, obstructions)
	return countDistinctGuardPositions(log)
}

func SolvePart2(input string) int {
	obstructions, guardPos := parseInput(input)
	up := Vector{X: 0, Y: -1}
	guard := Guard{Position: guardPos, Direction: up}
	log, _ := simulateGuard(guard, obstructions)
	return countTimeParadoxOptions(log, obstructions)
}

func countDistinctGuardPositions(log []Guard) int {
	uniqueElements := make(map[Vector]struct{})
	for _, guard := range log {
		uniqueElements[guard.Position] = struct{}{}
	}
	return len(uniqueElements)
}

func countTimeParadoxOptions(log []Guard, obstructions [][]rune) int {
	return 0
}

func simulateGuard(guard Guard, obstructions [][]rune) (log []Guard, isLoop bool) {
	for isInBounds(guard.Position, obstructions) {
		// Check if we have already been in this position
		if slices.Contains(log, guard) {
			return log, true
		}
		log = append(log, guard)
		guard = guard.Move()
		if isGuardFacingObstruction(guard, obstructions) {
			guard = guard.TurnRight()
		}
	}
	return log, false
}

func isGuardFacingObstruction(guard Guard, obstructions [][]rune) bool {
	nextPos := guard.Move().Position
	return isInBounds(nextPos, obstructions) &&
		obstructions[nextPos.Y][nextPos.X] == '#'
}

func isInBounds(position Vector, obstructions [][]rune) bool {
	return position.X >= 0 &&
		position.X < len(obstructions[0]) &&
		position.Y >= 0 &&
		position.Y < len(obstructions)
}

func parseInput(input string) (obstructions [][]rune, guardPos Vector) {
	for y, line := range strings.Split(input, "\n") {
		guardIndex := strings.Index(line, "^")
		if guardIndex >= 0 {
			guardPos = Vector{X: guardIndex, Y: y}
		}
		obstructions = append(obstructions, []rune(line))
	}
	return obstructions, guardPos
}
