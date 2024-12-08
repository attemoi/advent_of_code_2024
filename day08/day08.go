package day08

import (
	"aoc2024/utils"
	"strings"
	"unicode"
)

type Vector = utils.Vector

type areaMap struct {
	antennasByType map[rune][]Vector
	width          int
	height         int
}

func newAreaMap() *areaMap {
	return &areaMap{
		antennasByType: make(map[rune][]Vector),
	}
}

func (am *areaMap) addAntenna(r rune, x int, y int) {
	am.antennasByType[r] = append(am.antennasByType[r], Vector{X: x, Y: y})
}

func (am *areaMap) isInBounds(position Vector) bool {
	return position.X >= 0 &&
		position.X < am.width &&
		position.Y >= 0 &&
		position.Y < am.height
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

func findAntiNodesPart1(areaMap areaMap) map[Vector]struct{} {
	antiNodes := make(map[Vector]struct{})
	for _, antennas := range areaMap.antennasByType {
		// Iterate over all pairs
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				// Calculate vector between two antennas
				v := antennas[j].Substract(antennas[i])

				antiNode1 := antennas[i].Substract(v)
				if areaMap.isInBounds(antiNode1) {
					antiNodes[antiNode1] = struct{}{}
				}

				antiNode2 := antennas[j].Add(v)
				if areaMap.isInBounds(antiNode2) {
					antiNodes[antiNode2] = struct{}{}
				}
			}
		}
	}
	return antiNodes
}

func findAntiNodesPart2(areaMap areaMap) map[Vector]struct{} {
	antiNodes := make(map[Vector]struct{})
	for _, antennas := range areaMap.antennasByType {
		// Iterate over all pairs
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				// Calculate vector between two antennas
				v := antennas[j].Substract(antennas[i])

				currentPos := antennas[i]
				for areaMap.isInBounds(currentPos) {
					antiNodes[currentPos] = struct{}{}
					currentPos = currentPos.Substract(v)
				}

				currentPos = antennas[j]
				for areaMap.isInBounds(currentPos) {
					antiNodes[currentPos] = struct{}{}
					currentPos = currentPos.Add(v)
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
