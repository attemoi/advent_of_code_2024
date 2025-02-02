package day06

import (
	"aoc2024/utils"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	input := utils.ReadInput("input_test.txt")
	expected := 41
	result := SolvePart1(string(input))
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSolvePart1WithTripleTurn(t *testing.T) {
	input := utils.ReadInput("input_test2.txt")
	expected := 3
	result := SolvePart1(string(input))
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

func TestSolvePart2(t *testing.T) {
	input := utils.ReadInput("input_test.txt")
	expected := 6
	result := SolvePart2(string(input))
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
