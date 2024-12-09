package day09

import (
	"aoc2024/utils"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	input := utils.ReadInput("input_test.txt")
	expected := 1928
	result := SolvePart1(string(input))
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	input := utils.ReadInput("input_test.txt")
	expected := 2858
	result := SolvePart2(string(input))
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
