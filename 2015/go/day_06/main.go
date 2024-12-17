package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input_06.txt"
	lineSep   = "\n"
)

type Light struct {
	isOn bool
}

func (l *Light) toggle() {
	l.isOn = !l.isOn
}

func (l *Light) turnOn() {
	l.isOn = true
}

func (l *Light) turnOff() {
	l.isOn = false
}

type Grid struct {
	lights [][]Light
}

func NewGrid(length, width int) *Grid {
	grid := make([][]Light, length)
	rows := make([]Light, length*width)
	for i := range grid {
		grid[i], rows = rows[:width], rows[width:]
	}

	return &Grid{
		lights: grid,
	}
}

func (g Grid) countLightsOn() int {
	var lightsOn int
	for _, row := range g.lights {
		for _, light := range row {
			if light.isOn {
				lightsOn += 1
			}
		}
	}

	return lightsOn
}

func (g *Grid) executeInstruction(instruction *Instruction) {
	for i := instruction.startLength; i <= instruction.endLength; i++ { //rows
		for j := instruction.startWidth; j <= instruction.endWidth; j++ { // lights in a row
			switch instruction.operation {
			case "turn on":
				g.lights[i][j].turnOn()
			case "turn off":
				g.lights[i][j].turnOff()
			case "toggle":
				g.lights[i][j].toggle()
			default:
				log.Fatalf("Unknown operation %q\n", instruction.operation)
			}
		}
	}
}

type Instruction struct {
	operation   string
	startLength int
	startWidth  int
	endLength   int
	endWidth    int
}

func parseInstruction(s string) *Instruction {
	var instruction string
	var startLength, startWidth, endLength, endWidth int

	fields := strings.Fields(s)
	if len(fields) == 4 {
		instruction = fields[0]
		startLength, startWidth = parsePositions(fields[1])
		endLength, endWidth = parsePositions(fields[3])

	} else if len(fields) == 5 {
		instruction = strings.Join(fields[:2], " ")
		startLength, startWidth = parsePositions(fields[2])
		endLength, endWidth = parsePositions(fields[4])
	} else {
		log.Fatalf("Could not parse line %q\n", s)
	}

	return &Instruction{
		operation:   instruction,
		startLength: startLength,
		startWidth:  startWidth,
		endLength:   endLength,
		endWidth:    endWidth,
	}
}

func parsePositions(positions string) (int, int) {
	splitted := strings.Split(positions, ",")
	if len(splitted) != 2 {
		log.Fatalf("Expected %v to have 2 elements separated by comma\n", positions)
	}

	length, err := strconv.Atoi(splitted[0])
	if err != nil {
		log.Fatalf("Could not parse %s as integer\n", splitted[0])
	}

	width, err := strconv.Atoi(splitted[1])
	if err != nil {
		log.Fatalf("Could not parse %s as integer\n", splitted[1])
	}

	return length, width
}

func main() {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), lineSep)
	g := NewGrid(1000, 1000)

	for _, line := range lines {
		ins := parseInstruction(line)
		g.executeInstruction(ins)
	}

	fmt.Printf("There are %d lights lit on\n", g.countLightsOn())
}
