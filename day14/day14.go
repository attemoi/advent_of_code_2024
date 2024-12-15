package day14

import (
	"aoc2024/utils"
	"fmt"
	"strings"
)

type Vector = utils.Vector

type robot struct {
	position Vector
	velocity Vector
}

func (r robot) simulate(seconds int, areaW int, areaH int) robot {
	robot := r
	newPos := r.position.Add(r.velocity.Multiply(seconds))
	newPos = Vector{X: wrapValue(newPos.X, areaW), Y: wrapValue(newPos.Y, areaH)}
	robot.position = newPos
	return robot
}

func wrapValue(value, max int) int {
	return ((value % max) + max) % max
}

func SolvePart1(input string) int {
	return SolvePart1WithParams(input, 101, 103)
}

func SolvePart2(input string) int {
	robots := parseRobots(input)

	for i := 0; i < 100000; i++ {
		for j, robot := range robots {
			robots[j] = robot.simulate(1, 101, 103)
		}
		if mightBeChristmasTree(robots) {
			print(robots)
			return i + 1
		}
	}

	return -1
}

func print(robots []robot) {
	lines := make([][]rune, 103)
	for y := range lines {
		lines[y] = []rune(strings.Repeat(" ", 101))
	}

	for _, robot := range robots {
		lines[robot.position.Y][robot.position.X] = '#'
	}

	for _, line := range lines {
		fmt.Println(string(line))
	}
}

func mightBeChristmasTree(robots []robot) bool {

	posMap := make(map[Vector]bool)
	for _, r := range robots {
		if _, exists := posMap[r.position]; exists {
			return false
		}
		posMap[r.position] = true
	}
	return true

}

func SolvePart1WithParams(input string, areaW int, areaH int) int {
	robots := parseRobots(input)

	var simulated []robot
	for _, robot := range robots {
		simulated = append(simulated, robot.simulate(100, areaW, areaH))
	}

	safetyFactor := 1
	for _, quadrant := range splitToQuadrants(simulated, areaW, areaH) {
		safetyFactor *= len(quadrant)
	}
	return safetyFactor
}

func splitToQuadrants(robots []robot, areaW int, areaH int) [4][]robot {
	middleX := areaW / 2
	middleY := areaH / 2
	var topLeft []robot
	var topRight []robot
	var bottomLeft []robot
	var bottomRight []robot
	for _, robot := range robots {
		if robot.position.X < middleX {
			if robot.position.Y < middleY {
				topLeft = append(topLeft, robot)
			} else if robot.position.Y > middleY {
				bottomLeft = append(bottomLeft, robot)
			}
		} else if robot.position.X > middleX {
			if robot.position.Y < middleY {
				topRight = append(topRight, robot)
			} else if robot.position.Y > middleY {
				bottomRight = append(bottomRight, robot)
			}
		}
	}
	return [4][]robot{
		topLeft, topRight, bottomLeft, bottomRight,
	}

}

func parseRobots(input string) []robot {
	var robots []robot
	for _, line := range strings.Split(input, "\n") {
		var r robot
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &r.position.X, &r.position.Y, &r.velocity.X, &r.velocity.Y)
		robots = append(robots, r)
	}
	return robots
}
