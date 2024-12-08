package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	input = "input_03.txt"
)

var pattern = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func extractPairs(matches [][][]byte) [][]int {
	var pairs [][]int
	for _, match := range matches {
		firstNum, _ := strconv.Atoi(string(match[1]))
		secondNum, _ := strconv.Atoi(string(match[2]))
		pairs = append(pairs, []int{firstNum, secondNum})
	}

	return pairs
}

func sumPairsProduct(pairs [][]int) int {
	var result int
	for _, pair := range pairs {
		product := pair[0] * pair[1]
		result += product
	}

	return result
}

func main() {
	data, _ := os.ReadFile(input)

	matches := pattern.FindAllSubmatch(data, -1)

	pairs := extractPairs(matches)
	result := sumPairsProduct(pairs)

	fmt.Printf("The sum of multiplications: %d\n", result)
}
