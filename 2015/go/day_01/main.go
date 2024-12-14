package main

import (
	"fmt"
	"log"
	"os"
)

const (
	inputFile = "input_01.txt"
)

func calcFloor(s string) int {
	var result int
	for _, r := range s {
		switch string(r) {
		case "(":
			result += 1
		case ")":
			result -= 1
		default:
			continue
		}
	}
	return result
}

func findFirstBasementPosition(s string) int {
	var currentFloor int
	for i, r := range s {
		switch string(r) {
		case "(":
			currentFloor += 1
		case ")":
			currentFloor -= 1
		default:
		}

		if currentFloor == -1 {
			return i + 1
		}
	}

	return -1
}

func main() {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)

	puzzleAnswer := calcFloor(input)
	fmt.Printf("The instructions take Santa to the floor %d\n", puzzleAnswer)

	puzzleAnswer02 := findFirstBasementPosition(input)
	fmt.Printf("The position of the character that causes Santa to first enter the basement is %d\n", puzzleAnswer02)
}
