package day11

import (
	"aoc2024/utils"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	input := utils.ReadInput("input_test.txt")
	expected := 60
	result := SolvePart1(string(input))
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
