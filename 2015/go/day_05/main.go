package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	inputFile  = "input_05.txt"
	stringsSep = "\n"
)

var badCombinations = []string{
	"ab",
	"cd",
	"pq",
	"xy",
}

var vowels = "aeiou"

func hasBadCombinations(s string) bool {
	for _, combination := range badCombinations {
		if strings.Contains(s, combination) {
			return true
		}
	}

	return false
}

func isNice(s string) bool {
	if hasBadCombinations(s) {
		return false
	}

	var vowelCount int
	var hasDoubleLetter bool
	previous := s[0]
	for i := 0; i < len(s); i++ {
		if strings.Contains(vowels, string(s[i])) {
			vowelCount += 1
		}
		if i > 0 && s[i] == previous {
			hasDoubleLetter = true
		}
		previous = s[i]
	}

	return vowelCount > 2 && hasDoubleLetter
}

func countNiceStrings(s string) int {
	var niceStringCount int
	strs := strings.Split(s, stringsSep)

	for _, str := range strs {
		if isNice(str) {
			niceStringCount += 1
		}
	}

	return niceStringCount
}

func containsPairTwice(s string) bool {
	if len(s) < 3 {
		return false
	}

	letters := strings.Split(s, "")

	for i := 0; i < len(letters)-3; i++ {
		next := i + 1
		pair := letters[i] + letters[next]
		if strings.Contains(s[i+2:], pair) {
			return true
		}
	}

	return false
}

func hasLetterBetweenSameLetters(s string) bool {
	if len(s) < 3 {
		return false
	}

	letters := strings.Split(s, "")
	for i := 1; i < len(letters)-1; i++ {
		if letters[i-1] == letters[i+1] {
			return true
		}
	}

	return false
}

func isNiceNew(s string) bool {
	return containsPairTwice(s) && hasLetterBetweenSameLetters(s)
}

func countNiceStringsNew(s string) int {
	var niceStringCount int
	strs := strings.Split(s, stringsSep)

	for _, str := range strs {
		if isNiceNew(str) {
			niceStringCount += 1
		}
	}

	return niceStringCount
}

func main() {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)
	fmt.Printf("There are %d nice strings in Santa's text file\n", countNiceStrings(input))

	fmt.Printf("There are %d nice strings in Santa's text file with a new model\n", countNiceStringsNew(input))
}
