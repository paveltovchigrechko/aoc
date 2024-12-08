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

func dumpLevel(report []int, level int) []int {
	var reducedReport []int
	if level == len(report)-1 {
		reducedReport = append(reducedReport, report[:level]...)
		return reducedReport
	}

	reducedReport = append(reducedReport, report[:level]...)
	reducedReport = append(reducedReport, report[level+1:]...)
	return reducedReport
}

func isSafeReportWithProblemDumpener(report []int) bool {
	return (levelDiffInRangeWithProblemDumpener(report) && isOrderedWithProblemDumpener(report))
}

func levelDiffInRangeWithProblemDumpener(report []int) bool {
	for i := len(report) - 1; i > 0; i-- {
		if !diffInRange(report[i], report[i-1]) {
			modifiedReport := dumpLevel(report, i)
			for j := len(modifiedReport) - 1; j > 0; j-- {
				if diffInRange(modifiedReport[j], modifiedReport[j-1]) {
					continue
				} else {
					newModifiedReport := dumpLevel(report, i-1)
					for k := len(newModifiedReport) - 1; k > 0; k-- {
						if !diffInRange(newModifiedReport[k], newModifiedReport[k-1]) {
							return false
						}
					}

					return true
				}
			}

			return true
		}
	}

	return true
}

func levelDiffInRange(report []int) bool {
	for i := len(report) - 1; i > 0; i-- {
		if !diffInRange(report[i], report[i-1]) {
			return false
		}
	}

	return true
}

func isOrderedWithProblemDumpener(report []int) bool {
	isAscending := report[0] < report[1]
	if isAscending {
		for i := len(report) - 1; i > 0; i-- {
			if report[i] > report[i-1] {
				continue
			}

			modifiedReport := dumpLevel(report, i)
			if slices.IsSorted(modifiedReport) && levelDiffInRange(modifiedReport) {
				continue
			}

			newModifiedReport := dumpLevel(report, i-1)
			if !slices.IsSorted(newModifiedReport) || !levelDiffInRange(newModifiedReport) {
				return false
			}
		}

		return true
	}

	for i := len(report) - 1; i > 0; i-- {
		if report[i] < report[i-1] {
			continue
		}

		modifiedReport := dumpLevel(report, i)
		slices.Reverse(modifiedReport)
		if slices.IsSorted(modifiedReport) {
			continue
		}

		newModifiedReport := dumpLevel(report, i-1)
		slices.Reverse(newModifiedReport)
		if !slices.IsSorted(newModifiedReport) {
			return false
		}
	}

	return true
}

func diffInRange(l, r int) bool {
	diff := math.Abs(float64(l - r))
	return diff <= 3 && diff >= 1
}

func calcSafeReportsWithProblemDumpener(reports [][]int) int {
	var safeReportsNum int
	for _, report := range reports {
		if isSafeReportWithProblemDumpener(report) {
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

	safeReportsWithProblemDumpener := calcSafeReportsWithProblemDumpener(reports)
	fmt.Printf("Safe reports with Problem Dumpener: %d\n", safeReportsWithProblemDumpener)
}
