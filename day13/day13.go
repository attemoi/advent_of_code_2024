package day13

import (
	"aoc2024/utils"
	"regexp"
	"strconv"
	"strings"
)

type Vector = utils.Vector

type machine struct {
	a     Vector
	b     Vector
	price Vector
}

func SolvePart1(input string) int {
	machines := parseMachines(input)
	tokens := 0
	for _, m := range machines {
		a, b := m.solve()
		tokens += 3*a + b
	}
	return tokens
}

func SolvePart2(input string) int {
	machines := parseMachines(input)
	tokens := 0
	for _, m := range machines {
		m.price = m.price.Add(Vector{X: 10000000000000, Y: 10000000000000})
		a, b := m.solve()
		tokens += 3*a + b
	}
	return tokens
}

func (m machine) solve() (a, b int) {
	a = (m.b.Y*m.price.X - m.b.X*m.price.Y) / (m.b.Y*m.a.X - m.b.X*m.a.Y)
	b = (m.price.X - m.a.X*a) / m.b.X

	if m.a.X*a+m.b.X*b != m.price.X || m.a.Y*a+m.b.Y*b != m.price.Y {
		return 0, 0
	}
	return a, b
}

func parseMachines(input string) []machine {
	var machines []machine
	for _, m := range strings.Split(input, "\n\n") {
		machines = append(machines, parseMachine(m))
	}
	return machines
}

func parseMachine(input string) machine {
	lines := strings.Split(input, "\n")
	a := parseButton(lines[0])
	b := parseButton(lines[1])
	price := parsePrice(lines[2])
	return machine{a: a, b: b, price: price}
}

func parseButton(input string) Vector {
	matches := regexp.MustCompile(`X\+(\d+), Y\+(\d+)`).FindStringSubmatch(input)
	x, _ := strconv.Atoi(matches[1])
	y, _ := strconv.Atoi(matches[2])
	return Vector{X: x, Y: y}
}

func parsePrice(input string) Vector {
	matches := regexp.MustCompile(`X=(\d+), Y=(\d+)`).FindStringSubmatch(input)
	x, _ := strconv.Atoi(matches[1])
	y, _ := strconv.Atoi(matches[2])
	return Vector{X: x, Y: y}
}
