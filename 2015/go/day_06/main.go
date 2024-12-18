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

type Device interface {
	GetIndication() int
	Toggle()
	TurnOff()
	TurnOn()
}

type Light struct {
	isOn bool
}

func (l *Light) TurnOn() {
	l.isOn = true
}

func (l *Light) Toggle() {
	l.isOn = !l.isOn
}

func (l *Light) TurnOff() {
	l.isOn = false
}

func (l Light) GetIndication() int {
	if l.isOn {
		return 1
	}
	return 0
}

type LightNew struct {
	brightness int
}

func (l *LightNew) TurnOn() {
	l.brightness += 1
}

func (l *LightNew) TurnOff() {
	if l.brightness > 0 {
		l.brightness -= 1
	}
}

func (l *LightNew) Toggle() {
	l.brightness += 2
}

func (l LightNew) GetIndication() int {
	return l.brightness
}

type Grid struct {
	lights [][]Device
}

func CreateGridWithLights(length, width int) *Grid {
	grid := make([][]Device, 0)

	for i := 0; i <= width; i++ {
		rows := make([]Device, 0)
		for j := 0; j <= length; j++ {
			l := &Light{}
			rows = append(rows, l)
		}
		grid = append(grid, rows)
	}

	return &Grid{
		lights: grid,
	}
}

func CreateGridWithLightsNew(length, width int) *Grid {
	grid := make([][]Device, 0)

	for i := 0; i <= width; i++ {
		rows := make([]Device, 0)
		for j := 0; j <= length; j++ {
			l := &LightNew{}
			rows = append(rows, l)
		}
		grid = append(grid, rows)
	}

	return &Grid{
		lights: grid,
	}
}

func (g Grid) countIndications() int {
	var lightsIndications int
	for _, row := range g.lights {
		for _, light := range row {
			lightsIndications += light.GetIndication()
		}
	}

	return lightsIndications
}

func (g *Grid) executeInstruction(instruction *Instruction) {
	for i := instruction.startLength; i <= instruction.endLength; i++ { //rows
		for j := instruction.startWidth; j <= instruction.endWidth; j++ { // lights in a row
			switch instruction.operation {
			case "turn on":
				g.lights[i][j].TurnOn()
			case "turn off":
				g.lights[i][j].TurnOff()
			case "toggle":
				g.lights[i][j].Toggle()
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
	firstGrid := CreateGridWithLights(1000, 1000)
	secondGrid := CreateGridWithLightsNew(1000, 1000)

	for _, line := range lines {
		ins := parseInstruction(line)
		firstGrid.executeInstruction(ins)
		secondGrid.executeInstruction(ins)
	}

	fmt.Printf("There are %d lights lit on\n", firstGrid.countIndications())
	fmt.Printf("The total brightness of all lights combined is %d\n", secondGrid.countIndications())
}
