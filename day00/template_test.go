package day00

import (
	"os"
	"testing"
)

func TestSolvePart1(t *testing.T) {
	input, _ := os.ReadFile("input.txt")
	expected := 60
	result := SolvePart1(string(input))
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
