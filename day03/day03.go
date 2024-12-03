package day03

import (
	"aoc2024/utils"
	"regexp"
	"strings"
)

type Instruction struct {
	Operation string
	Arguments []int
}

func SolvePart1(input string) int {
	instructions := parseInstructions(input)
	return processOnlyMul(instructions)
}

func SolvePart2(input string) int {
	instructions := parseInstructions(input)
	return process(instructions)
}

func parseInstructions(input string) []Instruction {
	mulPattern := `mul\(\d+,\d+\)`
	doPattern := `do\(\)`
	dontPattern := `don't\(\)`

	rawInstructions := regexp.
		MustCompile(mulPattern+"|"+doPattern+"|"+dontPattern).
		FindAllString(input, -1)
	return toInstructions(rawInstructions)
}

func toInstructions(rawInstructions []string) []Instruction {
	instructions := make([]Instruction, len(rawInstructions))
	for i, rawValue := range rawInstructions {
		instructions[i] = parseInstruction(rawValue)
	}
	return instructions
}

func parseInstruction(rawValue string) Instruction {
	matches := regexp.MustCompile(`(.*)\((.*)\)`).FindStringSubmatch(rawValue)
	operation := matches[1]
	arguments := matches[2]
	return Instruction{Operation: operation, Arguments: parseArguments(arguments)}
}

func parseArguments(rawValue string) []int {
	arguments, _ := utils.ToInts(strings.Split(rawValue, ","))
	return arguments
}

func processOnlyMul(instructions []Instruction) int {
	sum := 0
	for _, instruction := range instructions {
		if instruction.Operation == "mul" {
			sum += instruction.Arguments[0] * instruction.Arguments[1]
		}
	}
	return sum
}

func process(instructions []Instruction) int {
	sum := 0
	do := true
	for _, instruction := range instructions {
		if instruction.Operation == "do" {
			do = true
		} else if instruction.Operation == "don't" {
			do = false
		} else if do {
			if instruction.Operation == "mul" {
				sum += instruction.Arguments[0] * instruction.Arguments[1]
			}
		}
	}
	return sum
}
