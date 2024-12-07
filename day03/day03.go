package day03

import (
	"aoc2024/utils"
	"regexp"
	"strings"
)

type instruction struct {
	operation string
	arguments []int
}

func SolvePart1(input string) int {
	instructions := parseInstructions(input)
	return processOnlyMul(instructions)
}

func SolvePart2(input string) int {
	instructions := parseInstructions(input)
	return process(instructions)
}

func parseInstructions(input string) []instruction {
	mulPattern := `mul\(\d+,\d+\)`
	doPattern := `do\(\)`
	dontPattern := `don't\(\)`

	rawInstructions := regexp.
		MustCompile(mulPattern+"|"+doPattern+"|"+dontPattern).
		FindAllString(input, -1)
	return toInstructions(rawInstructions)
}

func toInstructions(rawInstructions []string) []instruction {
	instructions := make([]instruction, len(rawInstructions))
	for i, rawValue := range rawInstructions {
		instructions[i] = parseInstruction(rawValue)
	}
	return instructions
}

func parseInstruction(rawValue string) instruction {
	matches := regexp.MustCompile(`(.*)\((.*)\)`).FindStringSubmatch(rawValue)
	operation := matches[1]
	arguments := matches[2]
	return instruction{operation: operation, arguments: parseArguments(arguments)}
}

func parseArguments(rawValue string) []int {
	arguments, _ := utils.ToInts(strings.Split(rawValue, ","))
	return arguments
}

func processOnlyMul(instructions []instruction) int {
	sum := 0
	for _, instruction := range instructions {
		if instruction.operation == "mul" {
			sum += instruction.arguments[0] * instruction.arguments[1]
		}
	}
	return sum
}

func process(instructions []instruction) int {
	sum := 0
	do := true
	for _, instruction := range instructions {
		if instruction.operation == "do" {
			do = true
		} else if instruction.operation == "don't" {
			do = false
		} else if do {
			if instruction.operation == "mul" {
				sum += instruction.arguments[0] * instruction.arguments[1]
			}
		}
	}
	return sum
}
