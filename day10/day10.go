package day10

import (
	"aoc2024/utils"
	"strconv"
	"strings"
)

type Vector = utils.Vector

type heightMap struct {
	elevations [][]int
	width      int
	height     int
}

func (hm *heightMap) isInBounds(position Vector) bool {
	return position.X >= 0 &&
		position.X < hm.width &&
		position.Y >= 0 &&
		position.Y < hm.height
}

func (hm heightMap) elevationAt(position Vector) int {
	if !hm.isInBounds(position) {
		return -1
	}
	return hm.elevations[position.Y][position.X]
}

type trail struct {
	head  Vector
	paths []path
}

type path []Vector

func SolvePart1(input string) int {
	heightMap, trailheadPositions := parseHeightMap(input)
	trails := findTrails(trailheadPositions, heightMap)
	sumOfScores := 0
	for _, trail := range trails {
		sumOfScores += score(trail)
	}
	return sumOfScores
}

func SolvePart2(input string) int {
	heightMap, trailheadPositions := parseHeightMap(input)
	trails := findTrails(trailheadPositions, heightMap)
	sumOfRatings := 0
	for _, trail := range trails {
		sumOfRatings += rating(trail)
	}
	return sumOfRatings
}

func findTrails(trailheadPositions []Vector, heightMap heightMap) []trail {
	var trails []trail
	for _, position := range trailheadPositions {
		trails = append(trails, trail{head: position, paths: traverse(position, heightMap)})
	}
	return trails
}

func score(trail trail) int {
	uniqueEndPositions := make(map[Vector]struct{})
	for _, path := range trail.paths {
		uniqueEndPositions[path[len(path)-1]] = struct{}{}
	}
	return len(uniqueEndPositions)
}

func rating(trail trail) int {
	return len(trail.paths)
}

func traverse(position Vector, heightMap heightMap) (paths []path) {
	path := []Vector{position}

	if heightMap.elevationAt(position) == 9 {
		paths = append(paths, path)
		return
	}

	directions := []Vector{
		{X: 0, Y: -1}, // up
		{X: 1, Y: 0},  // right
		{X: 0, Y: 1},  // down
		{X: -1, Y: 0}, // left
	}

	for _, dir := range directions {
		nextPos := position.Add(dir)
		if heightMap.elevationAt(nextPos)-heightMap.elevationAt(position) == 1 {
			for _, pathFromNextPos := range traverse(nextPos, heightMap) {
				newPath := append(path, pathFromNextPos...)
				paths = append(paths, newPath)
			}
		}
	}
	return paths
}

func parseHeightMap(input string) (heightMap heightMap, trailheadPositions []Vector) {
	lines := strings.Split(input, "\n")
	heightMap.width = len(lines[0])
	heightMap.height = len(lines)
	heightMap.elevations = make([][]int, heightMap.height)

	for y, line := range lines {
		for x, rune := range line {
			elevation, _ := strconv.Atoi(string(rune))
			heightMap.elevations[y] = append(heightMap.elevations[y], elevation)
			if elevation == 0 {
				trailheadPositions = append(trailheadPositions, Vector{X: x, Y: y})
			}
		}
	}

	return heightMap, trailheadPositions
}
