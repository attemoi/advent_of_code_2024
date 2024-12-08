package day08

import (
	"strings"
	"unicode"
)

type vector struct {
	x, y int
}

type areaMap struct {
	antennasByType map[rune][]vector
	width          int
	height         int
}

func newAreaMap() *areaMap {
	return &areaMap{
		antennasByType: make(map[rune][]vector),
	}
}

func (am *areaMap) addAntenna(r rune, x int, y int) {
	am.antennasByType[r] = append(am.antennasByType[r], vector{x: x, y: y})
}

func (am *areaMap) isInBounds(position vector) bool {
	return position.x >= 0 &&
		position.x < am.width &&
		position.y >= 0 &&
		position.y < am.height
}

func SolvePart1(input string) int {
	areaMap := parseMap(input)
	antiNodes := findAntiNodesPart1(areaMap)
	return len(antiNodes)
}

func SolvePart2(input string) int {
	areaMap := parseMap(input)
	antiNodes := findAntiNodesPart2(areaMap)
	return len(antiNodes)
}

func findAntiNodesPart1(areaMap areaMap) map[vector]struct{} {
	antiNodes := make(map[vector]struct{})
	for _, antennas := range areaMap.antennasByType {
		// Iterate over all pairs
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				// Calculate vector between two antennas
				v := vector{x: antennas[j].x - antennas[i].x, y: antennas[j].y - antennas[i].y}

				antiNode1 := vector{antennas[i].x - v.x, antennas[i].y - v.y}
				if areaMap.isInBounds(antiNode1) {
					antiNodes[antiNode1] = struct{}{}
				}

				antiNode2 := vector{antennas[j].x + v.x, antennas[i].y + v.y}
				if areaMap.isInBounds(antiNode2) {
					antiNodes[antiNode2] = struct{}{}
				}
			}
		}
	}
	return antiNodes
}

func findAntiNodesPart2(areaMap areaMap) map[vector]struct{} {
	antiNodes := make(map[vector]struct{})
	for _, antennas := range areaMap.antennasByType {
		// Iterate over all pairs
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				// Calculate vector between two antennas
				v := vector{x: antennas[j].x - antennas[i].x, y: antennas[j].y - antennas[i].y}

				currentPos := antennas[i]
				for areaMap.isInBounds(currentPos) {
					antiNodes[currentPos] = struct{}{}
					currentPos.x -= v.x
					currentPos.y -= v.y
				}

				currentPos = antennas[j]
				for areaMap.isInBounds(currentPos) {
					antiNodes[currentPos] = struct{}{}
					currentPos.x += v.x
					currentPos.y += v.y
				}
			}
		}
	}
	return antiNodes
}

func parseMap(input string) areaMap {
	am := *newAreaMap()
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, rune := range line {
			if unicode.IsLetter(rune) || unicode.IsNumber(rune) {
				am.addAntenna(rune, x, y)
			}
		}
	}
	am.width = len(lines[0])
	am.height = len(lines)
	return am
}
