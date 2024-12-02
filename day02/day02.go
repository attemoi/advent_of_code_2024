package day02

import (
	"aoc2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func SolvePart1(input string) int {
	reports := parseReports(input)
	return countSafeReports(reports, isSafe)
}

func SolvePart2(input string) int {
	reports := parseReports(input)
	return countSafeReports(reports, isSafeWithDampener)
}

func parseReports(input string) (reports [][]int) {
	for _, line := range strings.Split(input, "\n") {
		// Trust input, ignore errors
		levels, _ := toInts(strings.Fields(line))
		reports = append(reports, levels)
	}
	return reports
}

func countSafeReports(reports [][]int, safetyCheck func([]int) bool) int {
	count := 0
	for _, report := range reports {
		if safetyCheck(report) {
			count++
		}
	}
	return count
}

func isSafe(report []int) bool {
	isAscending, isDescending := true, true
	for i := 0; i < len(report)-1; i++ {
		if report[i] < report[i+1] {
			isDescending = false
		}
		if report[i] > report[i+1] {
			isAscending = false
		}
		isUnordered := !isDescending && !isAscending
		diff := utils.Abs(report[i] - report[i+1])
		if isUnordered || diff == 0 || diff > 3 {
			return false
		}
	}
	return true
}

func isSafeWithDampener(report []int) bool {
	for i := range report {
		dampened := removeElementAtIndex(report, i)
		if isSafe(dampened) {
			return true
		}
	}
	return isSafe(report)
}

func removeElementAtIndex(s []int, index int) []int {
	ret := make([]int, 0, len(s)-1)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func toInts(strings []string) ([]int, error) {
	ints := make([]int, len(strings))
	for i, str := range strings {
		num, err := strconv.Atoi(str)
		if err != nil {
			return nil, fmt.Errorf("failed to convert %q to int at index %d: %w", str, i, err)
		}
		ints[i] = num
	}
	return ints, nil
}
