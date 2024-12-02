package utils

import (
	"os"
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
