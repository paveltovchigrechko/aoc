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
	input        = "input_02.txt"
	inputSep     = " "
	inputLineSep = "\n"
)

func readInputFile(fn string) string {
	rawInput, err := os.ReadFile(fn)
	if err != nil {
		log.Print(err)
	}
	return string(rawInput)
}

func getReports(s string) [][]int {
	lines := strings.Split(s, inputLineSep)
	reports := make([][]int, 0, len(lines))

	for _, line := range lines {
		digits := strings.Split(line, inputSep)
		report := make([]int, 0, len(line))
		for _, digit := range digits {
			level, _ := strconv.Atoi(digit)
			report = append(report, level)
		}
		reports = append(reports, report)
	}

	return reports
}

func isSafeReport(report []int) bool {
	for i := 0; i < len(report)-1; i++ {
		diff := math.Abs(float64(report[i] - report[i+1]))
		if diff > 3 || diff < 1 {
			return false
		}
	}

	isAscending := report[0] < report[1]
	if isAscending {
		return slices.IsSorted(report)
	} else {
		slices.Reverse(report)
		return slices.IsSorted(report)
	}
}

func calcSafeReports(reports [][]int) int {
	var safeReportsNum int
	for _, report := range reports {
		if isSafeReport(report) {
			safeReportsNum++
		}
	}

	return safeReportsNum
}

func main() {
	in := readInputFile(input)
	reports := getReports(in)
	safeReports := calcSafeReports(reports)
	fmt.Printf("Safe report number: %d\n", safeReports)
}
