package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	input    = "input_01.txt"
	inputSep = "   "
)

func ReadInputFile(fn string) string {
	rawInput, err := os.ReadFile(fn)
	if err != nil {
		log.Print(err)
	}
	return string(rawInput)
}

func GetLists(input string) ([]int, []int) {
	pairs := strings.Split(input, "\n")

	leftList := make([]int, 0, len(pairs))
	rightList := make([]int, 0, len(pairs))

	for _, pair := range pairs {
		strNums := strings.Split(pair, inputSep)
		left, _ := strconv.Atoi(strNums[0])
		leftList = append(leftList, left)

		right, _ := strconv.Atoi(strNums[1])
		rightList = append(rightList, right)

	}

	return leftList, rightList
}

func CalcDiff(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)
	var diff int
	for i := 0; i < len(left); i++ {
		diff += int(math.Abs(float64(left[i] - right[i])))
	}

	return diff
}

func CalcSimilarity(left, right []int) int {
	rightFrequency := map[int]int{}
	for _, digit := range right {
		rightFrequency[digit] += 1
	}

	var similarity int
	for _, digit := range left {
		if score, ok := rightFrequency[digit]; ok {
			similarity += digit * score
		}
	}

	return similarity
}

func main() {
	input := ReadInputFile(input)
	l, r := GetLists(input)
	diff := CalcDiff(l, r)
	fmt.Printf("List difference: %d\n", diff)
	similarity := CalcSimilarity(l, r)
	fmt.Printf("Similarity score: %d\n", similarity)
}
