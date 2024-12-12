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
var enhancedPattern = regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\))|(do\(\))|(don\'t\(\))`)

func extractPairs(matches [][][]byte) [][]int {
	var pairs [][]int
	for _, match := range matches {
		pair := extractPair(match)
		pairs = append(pairs, pair)
	}

	return pairs
}

func extractPair(match [][]byte) []int {
	firstNum, _ := strconv.Atoi(string(match[1]))
	secondNum, _ := strconv.Atoi(string(match[2]))
	return []int{firstNum, secondNum}
}

func sumPairsProduct(pairs [][]int) int {
	var result int
	for _, pair := range pairs {
		product := pair[0] * pair[1]
		result += product
	}

	return result
}

func enhancedPairs(matches [][][]byte) int {
	doIsOn := true
	result := 0
	for _, match := range matches {
		switch string(match[0]) {
		case "do()":
			doIsOn = true
		case "don't()":
			doIsOn = false
		default:
			if doIsOn {
				firstNum, _ := strconv.Atoi(string(match[2]))
				secondNum, _ := strconv.Atoi(string(match[3]))
				result += firstNum * secondNum
			}
		}
	}
	return result
}

func main() {
	data, _ := os.ReadFile(input)

	matches := pattern.FindAllSubmatch(data, -1)

	pairs := extractPairs(matches)
	result := sumPairsProduct(pairs)

	fmt.Printf("The sum of multiplications: %d\n", result)

	matches = enhancedPattern.FindAllSubmatch(data, -1)
	result = enhancedPairs(matches)
	fmt.Printf("The sum of enhanced multiplications: %d\n", result)
}
