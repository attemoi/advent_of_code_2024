package main

import (
	"aoc2024/day00"
	"aoc2024/day01"
	"aoc2024/day02"
	"aoc2024/day03"
	"aoc2024/day04"
	"aoc2024/day05"
	"aoc2024/day06"
	"aoc2024/day07"
	"aoc2024/day08"
	"aoc2024/day09"
	"aoc2024/utils"
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
		input = utils.ReadInput("day00/input.txt")
		fmt.Println("Part 1:", day00.SolvePart1(input))
		fmt.Println("Part 2:", day00.SolvePart2(input))
	case "01":
		input = utils.ReadInput("day01/input.txt")
		fmt.Println("Part 1:", day01.SolvePart1(input))
		fmt.Println("Part 2:", day01.SolvePart2(input))
	case "02":
		input = utils.ReadInput("day02/input.txt")
		fmt.Println("Part 1:", day02.SolvePart1(input))
		fmt.Println("Part 2:", day02.SolvePart2(input))
	case "03":
		input = utils.ReadInput("day03/input.txt")
		fmt.Println("Part 1:", day03.SolvePart1(input))
		fmt.Println("Part 2:", day03.SolvePart2(input))
	case "04":
		input = utils.ReadInput("day04/input.txt")
		fmt.Println("Part 1:", day04.SolvePart1(input))
		fmt.Println("Part 2:", day04.SolvePart2(input))
	case "05":
		input = utils.ReadInput("day05/input.txt")
		fmt.Println("Part 1:", day05.SolvePart1(input))
		fmt.Println("Part 2:", day05.SolvePart2(input))
	case "06":
		input = utils.ReadInput("day06/input.txt")
		fmt.Println("Part 1:", day06.SolvePart1(input))
		fmt.Println("Part 2:", day06.SolvePart2(input))
	case "07":
		input = utils.ReadInput("day07/input.txt")
		fmt.Println("Part 1:", day07.SolvePart1(input))
		fmt.Println("Part 2:", day07.SolvePart2(input))
	case "08":
		input = utils.ReadInput("day08/input.txt")
		fmt.Println("Part 1:", day08.SolvePart1(input))
		fmt.Println("Part 2:", day08.SolvePart2(input))
	case "09":
		input = utils.ReadInput("day09/input.txt")
		fmt.Println("Part 1:", day09.SolvePart1(input))
		fmt.Println("Part 2:", day09.SolvePart2(input))
	default:
		fmt.Printf("Day %s not solved yet!\n", day)
	}
}
