package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	inputFile = "input_02.txt"
	lineSep   = "\n"
	digitSep  = "x"
)

func parseInput(input string) [][]int {
	dimensions := [][]int{}
	lines := strings.Split(input, lineSep)

	for _, line := range lines {
		digits := strings.Split(line, digitSep)
		parsed := []int{}
		for _, digit := range digits {
			num, err := strconv.Atoi(digit)
			if err != nil {
				continue
			}
			parsed = append(parsed, num)
		}

		dimensions = append(dimensions, parsed)
	}

	return dimensions
}

func calcSurfaceArea(dimensions []int) int {
	if len(dimensions) != 3 {
		log.Printf("Array has %d dimensions, want 3\n", len(dimensions))
		return 0
	}

	area := 2*dimensions[0]*dimensions[1] + 2*dimensions[1]*dimensions[2] + 2*dimensions[0]*dimensions[2]
	return area
}

func calcMinSideArea(dimensions []int) int {
	if len(dimensions) != 3 {
		log.Printf("Array has %d dimensions, want 3\n", len(dimensions))
		return 0
	}

	minimalArea := dimensions[0] * dimensions[2]
	for i := 0; i < 2; i++ {
		area := dimensions[i] * dimensions[i+1]
		if area < minimalArea {
			minimalArea = area
		}
	}

	return minimalArea
}

func calcTotalArea(dimensions [][]int) int {
	var totalAreaFeet int
	for _, d := range dimensions {
		area := calcSurfaceArea(d)
		minArea := calcMinSideArea(d)
		totalAreaFeet += area + minArea
	}

	return totalAreaFeet
}

func calcShortestDistance(dimensions []int) int {
	if len(dimensions) != 3 {
		log.Printf("Array has %d dimensions, want 3\n", len(dimensions))
		return 0
	}

	slices.Sort(dimensions)
	return 2*dimensions[0] + 2*dimensions[1]
}

func calcRibbonBow(dimensions []int) int {
	if len(dimensions) != 3 {
		log.Printf("Array has %d dimensions, want 3\n", len(dimensions))
		return 0
	}

	return dimensions[0] * dimensions[1] * dimensions[2]
}

func calcWrappingPaper(input string) int {
	ds := parseInput(input)
	return calcTotalArea(ds)
}

func calcTotalRibbon(dimensions [][]int) int {
	var totalRibbonFeet int

	for _, present := range dimensions {
		wrappingRibbon := calcShortestDistance(present)
		bowRibbon := calcRibbonBow(present)
		totalRibbonFeet += wrappingRibbon + bowRibbon
	}

	return totalRibbonFeet
}

func calcRibbon(input string) int {
	ds := parseInput(input)
	return calcTotalRibbon(ds)
}

func main() {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)
	fmt.Printf("Elves need %d square feet of wrapping paper\n", calcWrappingPaper(input))

	fmt.Printf("Elves need %d feet of ribbon\n", calcRibbon(input))
}
