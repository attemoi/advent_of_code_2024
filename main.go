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
	"aoc2024/day10"
	"aoc2024/utils"
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <day>")
		return
	}

	var part1Solver, part2Solver func(input string) int

	day := os.Args[1]
	var input string
	switch day {
	case "00":
		part1Solver = day00.SolvePart1
		part2Solver = day00.SolvePart2
	case "01":
		part1Solver = day01.SolvePart1
		part2Solver = day01.SolvePart2
	case "02":
		part1Solver = day02.SolvePart1
		part2Solver = day02.SolvePart2
	case "03":
		part1Solver = day03.SolvePart1
		part2Solver = day03.SolvePart2
	case "04":
		part1Solver = day04.SolvePart1
		part2Solver = day04.SolvePart2
	case "05":
		part1Solver = day05.SolvePart1
		part2Solver = day05.SolvePart2
	case "06":
		part1Solver = day06.SolvePart1
		part2Solver = day06.SolvePart2
	case "07":
		part1Solver = day07.SolvePart1
		part2Solver = day07.SolvePart2
	case "08":
		part1Solver = day08.SolvePart1
		part2Solver = day08.SolvePart2
	case "09":
		part1Solver = day09.SolvePart1
		part2Solver = day09.SolvePart2
	case "10":
		part1Solver = day10.SolvePart1
		part2Solver = day10.SolvePart2
	default:
		fmt.Printf("Day %s not solved yet!\n", day)
		return
	}

	input = utils.ReadInput("day" + day + "/input.txt")
	start := time.Now()
	fmt.Println("Part 1:", part1Solver(input))
	fmt.Println("Part 2:", part2Solver(input))
	fmt.Println("Took", time.Since(start))

}
