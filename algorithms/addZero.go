package algorithms

import (
	"fmt"
	"strings"
)

var zeros = strings.Repeat("0", 1000) // Use the maximum length here

func AddZero(strs [][]string) [][]string {
	for _, line := range strs {
		count := len(line) - 1
		for i := range line {
			line[i] += zeros[:count-i]
		}
	}
	fmt.Println(strs)
	return strs
}
