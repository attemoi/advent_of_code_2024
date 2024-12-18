package day15

import (
	"aoc2024/utils"
	"testing"
)

func TestSolvePart1SmallInput(t *testing.T) {
	input := utils.ReadInput("input_test_small.txt")
	expected := 2028
	result := SolvePart1(string(input))
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSolvePart1LargeInput(t *testing.T) {
	input := utils.ReadInput("input_test_large.txt")
	expected := 10092
	result := SolvePart1(string(input))
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	input := utils.ReadInput("input_test_large.txt")
	expected := 9021
	result := SolvePart2(string(input))
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
