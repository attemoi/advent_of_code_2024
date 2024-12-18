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
	boxes  []box
	walls  map[Vector]struct{}
}

type Direction int

type box struct {
	pos   Vector
	width int
}

func newWarehouse() *warehouse {
	return &warehouse{
		walls: make(map[Vector]struct{}),
	}
}

func (wh warehouse) collisionCheck(with box) ([]*box, bool) {
	var collisions []*box
	for i := range wh.boxes {
		target := &wh.boxes[i]
		if target.pos.Y != with.pos.Y {
			continue
		}
		if target.pos.X > with.pos.X+with.width-1 {
			continue
		}
		if target.pos.X+target.width-1 < with.pos.X {
			continue
		}
		collisions = append(collisions, target)
	}

	isWallCollision := false
	for x := with.pos.X; x < with.pos.X+with.width; x++ {
		if _, exists := wh.walls[Vector{X: x, Y: with.pos.Y}]; exists {
			isWallCollision = true
			break
		}
	}
	return collisions, isWallCollision
}

func SolvePart1(input string) int {
	inputSections := strings.Split(input, "\n\n")
	warehouse := newWarehouse()
	parseWarehouse(inputSections[0], warehouse)
	for _, dir := range parseMoves(inputSections[1]) {
		moveRobot(warehouse, dir)
	}
	return sumOfBoxCoordinates(*warehouse)
}

func SolvePart2(input string) int {
	inputSections := strings.Split(input, "\n\n")
	warehouse := newWarehouse()
	parseWarehouseDoubleWidth(inputSections[0], warehouse)
	for _, dir := range parseMoves(inputSections[1]) {
		moveRobot(warehouse, dir)
	}
	return sumOfBoxCoordinates(*warehouse)
}

func sumOfBoxCoordinates(warehouse warehouse) int {
	sumOfCoordinates := 0
	for _, box := range warehouse.boxes {
		sumOfCoordinates += gpsCoordinate(box)
	}
	return sumOfCoordinates
}

func moveRobot(warehouse *warehouse, dir Vector) {
	collisions, isWall := warehouse.collisionCheck(box{pos: warehouse.robot.Add(dir), width: 1})
	if isWall {
		return
	} else if len(collisions) == 0 {
		warehouse.robot = warehouse.robot.Add(dir)
	} else {
		if canBePushed(warehouse, collisions[0], dir) {
			push(warehouse, collisions[0], dir)
			warehouse.robot = warehouse.robot.Add(dir)
			return
		}
	}
}

func canBePushed(warehouse *warehouse, box *box, dir Vector) bool {
	target := *box
	target.pos = target.pos.Add(dir)
	collisions, isWallCollision := warehouse.collisionCheck(target)
	if isWallCollision {
		return false
	}
	for _, collision := range collisions {
		if collision == box {
			continue
		}
		if !canBePushed(warehouse, collision, dir) {
			return false
		}
	}
	return true
}

func push(warehouse *warehouse, box *box, dir Vector) {
	target := *box
	target.pos = target.pos.Add(dir)
	collisions, _ := warehouse.collisionCheck(target)
	for _, collision := range collisions {
		if collision == box {
			continue
		}
		push(warehouse, collision, dir)
	}
	box.pos = box.pos.Add(dir)
}

func gpsCoordinate(box box) int {
	return 100*box.pos.Y + box.pos.X
}

func parseWarehouse(input string, warehouse *warehouse) {
	lines := strings.Split(input, "\n")
	warehouse.height = len(lines)
	warehouse.width = len(lines[0])
	for y, line := range lines {
		for x, rune := range line {
			if rune == 'O' {
				warehouse.boxes = append(warehouse.boxes, box{pos: Vector{X: x, Y: y}, width: 1})
			} else if rune == '#' {
				warehouse.walls[Vector{X: x, Y: y}] = struct{}{}
			} else if rune == '@' {
				warehouse.robot = Vector{X: x, Y: y}
			}

		}
	}
}

func parseWarehouseDoubleWidth(input string, warehouse *warehouse) {
	lines := strings.Split(input, "\n")
	warehouse.height = len(lines)
	warehouse.width = len(lines[0]) * 2

	for y, line := range lines {
		for x, rune := range line {
			if rune == 'O' {
				warehouse.boxes = append(warehouse.boxes, box{pos: Vector{X: x * 2, Y: y}, width: 2})
			} else if rune == '#' {
				warehouse.walls[Vector{X: x * 2, Y: y}] = struct{}{}
				warehouse.walls[Vector{X: x*2 + 1, Y: y}] = struct{}{}
			} else if rune == '@' {
				warehouse.robot = Vector{X: x * 2, Y: y}
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
