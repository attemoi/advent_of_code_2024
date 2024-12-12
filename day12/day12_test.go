package day12

import (
	"aoc2024/utils"
	"testing"
)

func TestSolvePart1WithSmallInput(t *testing.T) {
	input := utils.ReadInput("input_test_small.txt")
	expected := 140
	result := SolvePart1(string(input))
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSolvePart1WithLargerInput(t *testing.T) {
	input := utils.ReadInput("input_test_large.txt")
	expected := 1930
	result := SolvePart1(string(input))
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSolvePart2withSmallerInput(t *testing.T) {
	input := utils.ReadInput("input_test_e.txt")
	expected := 236
	result := SolvePart2(string(input))
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSolvePart2WithCornerInput(t *testing.T) {
	input := utils.ReadInput("input_test_corner.txt")
	expected := 368
	result := SolvePart2(string(input))
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSolvePart2WithLargerInput(t *testing.T) {
	input := utils.ReadInput("input_test_large.txt")
	expected := 1206
	result := SolvePart2(string(input))
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
