package day05

import (
	"aoc2024/utils"
	"slices"
	"strconv"
	"strings"
)

type Rule struct {
	X int
	Y int
}

func SolvePart1(input string) int {
	rules, pages := parseInput(input)
	filteredPages := pagesInRightOrder(pages, rules)
	return sumOfMiddlePageNumbers(filteredPages)
}

func SolvePart2(input string) int {
	rules, pages := parseInput(input)
	filteredPages := pagesInWrongOrder(pages, rules)
	orderPages(filteredPages, rules)
	return sumOfMiddlePageNumbers(filteredPages)
}

func orderPages(pages [][]int, rules []Rule) {
	for _, page := range pages {
		slices.SortFunc(page, ruleComparator(rules))
	}
}

func sumOfMiddlePageNumbers(pages [][]int) int {
	sum := 0
	for _, page := range pages {
		middlePageNum := page[len(page)/2]
		sum += middlePageNum
	}
	return sum
}

func pagesInRightOrder(pages [][]int, rules []Rule) [][]int {
	var filtered [][]int
	for _, page := range pages {
		if slices.IsSortedFunc(page, ruleComparator(rules)) {
			filtered = append(filtered, page)
		}
	}
	return filtered
}

func pagesInWrongOrder(pages [][]int, rules []Rule) [][]int {
	var filtered [][]int
	for _, page := range pages {
		if !slices.IsSortedFunc(page, ruleComparator(rules)) {
			filtered = append(filtered, page)
		}
	}
	return filtered
}

func ruleComparator(rules []Rule) func(a, b int) int {
	return func(a, b int) int {
		if slices.Contains(rules, Rule{X: a, Y: b}) {
			return -1
		}
		return 0
	}
}

func parseInput(input string) ([]Rule, [][]int) {
	parts := strings.SplitN(input, "\n\n", 2)
	return parseOrderingRules(parts[0]), parseUpdateNumbers(parts[1])
}

func parseOrderingRules(input string) []Rule {
	var rules []Rule
	for _, line := range strings.Split(input, "\n") {
		separatorIndex := strings.Index(line, "|")
		// Trust the input, ignore errors
		x, _ := strconv.Atoi(line[:separatorIndex])
		y, _ := strconv.Atoi(line[separatorIndex+1:])
		rules = append(rules, Rule{X: x, Y: y})
	}
	return rules
}

func parseUpdateNumbers(input string) [][]int {
	var updateNumbers [][]int
	for _, line := range strings.Split(input, "\n") {
		// Trust the input, ignore errors
		numbers, _ := utils.ToInts(strings.Split(line, ","))
		updateNumbers = append(updateNumbers, numbers)
	}
	return updateNumbers
}
