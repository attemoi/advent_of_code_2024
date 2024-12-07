package day07

import (
	"aoc2024/utils"
	"strconv"
	"strings"
)

type equation struct {
	result  int
	numbers []int
}

type operator func(int, int) int

func SolvePart1(input string) int {
	equations := parseEquations(input)
	operators := []operator{multiply, add}
	possibleEquations := determinePossible(equations, operators)
	return sumOfResults(possibleEquations)
}

func SolvePart2(input string) int {
	equations := parseEquations(input)
	operators := []operator{multiply, add, concat}
	possibleEquations := determinePossible(equations, operators)
	return sumOfResults(possibleEquations)
}

func determinePossible(equations []equation, operators []operator) []equation {
	var possible []equation
	for _, eq := range equations {
		if isPossible(eq, operators) {
			possible = append(possible, eq)
		}
	}
	return possible
}

func isPossible(eq equation, operators []operator) bool {

	if len(eq.numbers) == 1 {
		return eq.numbers[0] == eq.result
	} else if eq.numbers[0] > eq.result {
		// None of the operators can decrease the value and we are already over
		return false
	}

	for _, op := range operators {
		newEq := eq
		newEq.numbers = append(
			[]int{op(eq.numbers[0], eq.numbers[1])},
			eq.numbers[2:]...,
		)
		if isPossible(newEq, operators) {
			return true
		}
	}
	return false
}

func multiply(a, b int) int {
	return a * b
}

func add(a, b int) int {
	return a + b
}

func concat(a, b int) int {
	num, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return num
}

func sumOfResults(equations []equation) int {
	sum := 0
	for _, eq := range equations {
		sum += eq.result
	}
	return sum
}

func parseEquations(input string) []equation {
	var equations []equation
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ":")
		result, _ := strconv.Atoi(parts[0])
		numbers, _ := utils.ToInts(strings.Fields(parts[1]))
		equations = append(equations, equation{result: result, numbers: numbers})
	}
	return equations
}
