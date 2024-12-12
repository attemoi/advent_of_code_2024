package day12

import (
	"aoc2024/utils"
	"strings"
)

// CONSTANTS

type Vector = utils.Vector

var (
	up         = Vector{X: 0, Y: -1}
	right      = Vector{X: 1, Y: 0}
	down       = Vector{X: 0, Y: 1}
	left       = Vector{X: -1, Y: 0}
	directions = []Vector{
		up, right, down, left,
	}
)

// GARDEN

type garden struct {
	plants [][]rune
	width  int
	height int
}

func (g garden) plantAt(pos Vector) rune {
	if !g.isInBounds(pos) {
		return -1
	}
	return g.plants[pos.Y][pos.X]
}

func (g garden) findRegions() []region {

	visited := make([][]bool, g.height)
	for i := range visited {
		visited[i] = make([]bool, g.width)
	}

	var dfs func(pos Vector, plant rune, reg region) region
	dfs = func(pos Vector, plant rune, reg region) region {

		if !g.isInBounds(pos) ||
			visited[pos.Y][pos.X] ||
			g.plants[pos.Y][pos.X] != plant {
			return reg
		}
		visited[pos.Y][pos.X] = true

		reg.plots++

		// Count corners and fences
		orthogonalDir1 := up
		orthogonalDir2 := right
		diagonalDir := up.Add(right)

		for i := 0; i < 4; i++ {
			orth1Plant := g.plantAt(pos.Add(orthogonalDir1))
			orth2Plant := g.plantAt(pos.Add(orthogonalDir2))
			diagPlant := g.plantAt(pos.Add(diagonalDir))
			if plant != orth1Plant && plant != orth2Plant {
				reg.corners++
			}
			if plant == orth1Plant && plant == orth2Plant && plant != diagPlant {
				reg.corners++
			}
			if orth1Plant != plant {
				reg.fences++
			}
			orthogonalDir1 = orthogonalDir1.Rotate90DegreesClockwise()
			orthogonalDir2 = orthogonalDir2.Rotate90DegreesClockwise()
			diagonalDir = diagonalDir.Rotate90DegreesClockwise()
		}

		// Continue search in orthogonal directions
		for _, dir := range directions {
			reg = dfs(pos.Add(dir), plant, reg)
		}
		return reg
	}

	// depth-first-search for each plot
	var regions []region
	for y, row := range g.plants {
		for x, plant := range row {
			if visited[y][x] {
				continue
			}
			regions = append(regions, dfs(Vector{X: x, Y: y}, plant, region{}))
		}
	}
	return regions

}

// REGION

type region struct {
	fences  int
	corners int
	plots   int
}

// SOLUTION

func SolvePart1(input string) int {
	garden := parseGarden(input)
	totalPrice := 0
	for _, r := range garden.findRegions() {
		totalPrice += r.plots * r.fences
	}
	return totalPrice
}

func SolvePart2(input string) int {
	garden := parseGarden(input)
	totalPrice := 0
	for _, r := range garden.findRegions() {
		totalPrice += r.plots * r.corners
	}
	return totalPrice
}

func (g *garden) isInBounds(position Vector) bool {
	return position.X >= 0 &&
		position.X < g.width &&
		position.Y >= 0 &&
		position.Y < g.height
}

func parseGarden(input string) (garden garden) {
	lines := strings.Split(input, "\n")
	garden.width = len(lines[0])
	garden.height = len(lines)
	garden.plants = make([][]rune, len(lines))

	for y, line := range lines {
		for _, rune := range line {
			garden.plants[y] = append(garden.plants[y], rune)
		}
	}
	return garden
}
