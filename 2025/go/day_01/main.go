package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	input     = "input_01.txt"
	separator = "\n"
)

func ReadInputFile(fn string) string {
	rawInput, err := os.ReadFile(fn)
	if err != nil {
		log.Print(err)
	}
	return string(rawInput)
}

func RotateLeft(curPos, shiftVal int) int {
	newPos := (curPos - shiftVal) % 100
	if newPos < 0 {
		newPos = 100 + newPos
	}
	return newPos
}

func RotateRight(curPos, shiftVal int) int {
	newPos := (curPos + shiftVal) % 100
	return newPos
}

func ParseInstruction(s string) (string, int) {
	direction := string(s[0])
	shift, err := strconv.Atoi(s[1:])
	if err != nil {
		log.Printf("Could not parse value %d for instruction %s\n", shift, s)
		return direction, 0
	}

	return direction, shift
}

func main() {
	currentPosition := 50
	password := 0
	data := ReadInputFile(input)
	instructions := strings.Split(data, separator)

	for line, instruction := range instructions {
		direction, shift := ParseInstruction(instruction)
		switch direction {
		case "L":
			currentPosition = RotateLeft(currentPosition, shift)
		case "R":
			currentPosition = RotateRight(currentPosition, shift)
		default:
			log.Printf("Unknown instruction in line %d: %s\n", line, instruction)

		}

		if currentPosition == 0 {
			password += 1
		}
	}

	fmt.Printf("Password: %d\n", password)
}
