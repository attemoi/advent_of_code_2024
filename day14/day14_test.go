package day14

import (
	"aoc2024/utils"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	input := utils.ReadInput("input_test.txt")
	expected := 12
	result := SolvePart1WithParams(string(input), 11, 7)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
