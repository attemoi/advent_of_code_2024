package utils

import (
	"fmt"
	"os"
	"strconv"
)

func ReadInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func ToInts(strings []string) ([]int, error) {
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
