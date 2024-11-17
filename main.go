package main

import (
	"aoc2024/day00"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <day>")
		return
	}

	day := os.Args[1]
	var input string
	switch day {
	case "00":
		input = readInput("day00/input.txt")
		fmt.Println("Part 1:", day00.SolvePart1(input))
		fmt.Println("Part 2:", day00.SolvePart2(input))
	default:
		fmt.Printf("Day %s not solved yet!\n", day)
	}
}

func readInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}
