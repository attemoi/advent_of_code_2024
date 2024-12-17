package day15

import (
	"aoc2024/utils"
	"strings"
)

type Vector = utils.Vector

type warehouse struct {
	width  int
	height int
	robot  Vector
	boxes  map[Vector]struct{}
	walls  map[Vector]struct{}
}

func newWarehouse() *warehouse {
	return &warehouse{
		boxes: make(map[Vector]struct{}),
		walls: make(map[Vector]struct{}),
	}
}

func SolvePart1(input string) int {
	inputSections := strings.Split(input, "\n\n")
	warehouse := newWarehouse()
	parseWarehouse(inputSections[0], warehouse)
	for _, dir := range parseMoves(inputSections[1]) {
		moveRobot(warehouse, dir)
	}

	sumOfCoordinates := 0
	for box := range warehouse.boxes {
		sumOfCoordinates += gpsCoordinate(box)
	}
	return sumOfCoordinates
}

func SolvePart2(input string) int {
	// Implement part 2 logic here
	return -1
}

func moveRobot(warehouse *warehouse, dir Vector) {
	moved := push(warehouse, warehouse.robot, dir)
	if moved {
		warehouse.robot = warehouse.robot.Add(dir)
	}
}

func push(warehouse *warehouse, from Vector, dir Vector) bool {
	target := from.Add(dir)
	_, isPushingIntoWall := warehouse.walls[target]
	_, isPushingIntoBox := warehouse.boxes[target]
	if isPushingIntoWall {
		return false
	} else if isPushingIntoBox {
		moved := push(warehouse, target, dir)
		if moved {
			if _, exists := warehouse.boxes[from]; exists {
				delete(warehouse.boxes, from)
				warehouse.boxes[target] = struct{}{}
			}
		}
		return moved
	}
	if _, exists := warehouse.boxes[from]; exists {
		delete(warehouse.boxes, from)
		warehouse.boxes[target] = struct{}{}
	}
	return true

}

func gpsCoordinate(boxPos Vector) int {
	return 100*boxPos.Y + boxPos.X
}

func parseWarehouse(input string, warehouse *warehouse) {
	lines := strings.Split(input, "\n")
	warehouse.height = len(lines)
	warehouse.width = len(lines[0])
	for y, line := range lines {
		for x, rune := range line {
			if rune == 'O' {
				warehouse.boxes[Vector{X: x, Y: y}] = struct{}{}
			} else if rune == '#' {
				warehouse.walls[Vector{X: x, Y: y}] = struct{}{}
			} else if rune == '@' {
				warehouse.robot = Vector{X: x, Y: y}
			}
		}
	}
}

func parseMoves(input string) []Vector {
	var moves []Vector
	inputAsSingleLine := strings.Replace(input, "\n", "", -1)
	for _, rune := range inputAsSingleLine {
		moves = append(moves, parseMove(rune))
	}
	return moves
}

func parseMove(r rune) Vector {
	switch r {
	case '>':
		return Vector{X: 1, Y: 0}
	case 'v':
		return Vector{X: 0, Y: 1}
	case '<':
		return Vector{X: -1, Y: 0}
	case '^':
		return Vector{X: 0, Y: -1}
	default:
		panic("Direction not recognized")
	}
}
