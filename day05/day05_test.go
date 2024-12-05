package day05

import (
	"aoc2024/utils"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	input := utils.ReadInput("input_test.txt")
	expected := 143
	result := SolvePart1(string(input))
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	input := utils.ReadInput("input_test.txt")
	expected := 123
	result := SolvePart2(string(input))
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
