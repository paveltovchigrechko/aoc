package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input_07.txt"
	AND       = "AND"
	OR        = "OR"
	NOT       = "NOT"
	LSHIFT    = "LSHIFT"
	RSHIFT    = "RSHIFT"

	inputSep       = "\n"
	instructionSep = "->"
)

type Instruction struct {
	leftOperand     string
	rightOperand    string
	operator        string
	destination     string
	instructionsNum int
}

func parseLine(s string) []string {
	ins := strings.Split(s, instructionSep)
	params := strings.Fields(strings.Join(ins, " "))
	return params
}

func parseInstruction(instructions []string) *Instruction {
	var leftOperand, rightOperand, operator, destination string
	instructionsNum := len(instructions) - 1
	destination = instructions[instructionsNum] // last instruction is always a destination ID

	switch instructionsNum {
	// "123 -> aq" or "abc -> def"
	case 1:
		leftOperand = instructions[0]
	// "NOT 23 -> aq" or "NOT ret -> bef"
	case 2:
		operator = instructions[0]
		rightOperand = instructions[1]
	// "34 AND adf -> gjt"
	case 3:
		leftOperand = instructions[0]
		operator = instructions[1]
		rightOperand = instructions[2]
	default:
		log.Fatalf("Unsupported number of instruction: %d, want 1-3\n", len(instructions))
	}

	return &Instruction{
		leftOperand:     leftOperand,
		rightOperand:    rightOperand,
		operator:        operator,
		destination:     destination,
		instructionsNum: instructionsNum,
	}
}

func ExecuteInstruction(ins Instruction, wires map[string]uint16) {
	switch ins.operator {
	case NOT:
		executeNOTOperation(ins, wires)
	case AND:
		executeANDOperation(ins, wires)
	case OR:
		executeOROperation(ins, wires)
	case LSHIFT:
		executeLSHIFTOperation(ins, wires)
	case RSHIFT:
		executeRSHIFTOperation(ins, wires)
	default:
		executeAssignOperation(ins, wires)
	}
}

func executeNOTOperation(ins Instruction, wires map[string]uint16) {
	circuitValue := handleOperand(ins.rightOperand, wires)
	wires[ins.destination] = ^circuitValue
}

func executeANDOperation(ins Instruction, wires map[string]uint16) {
	leftValue := handleOperand(ins.leftOperand, wires)
	rightValue := handleOperand(ins.rightOperand, wires)

	wires[ins.destination] = leftValue & rightValue
}

func executeOROperation(ins Instruction, wires map[string]uint16) {
	leftValue := handleOperand(ins.leftOperand, wires)
	rightValue := handleOperand(ins.rightOperand, wires)

	wires[ins.destination] = leftValue | rightValue
}

func executeLSHIFTOperation(ins Instruction, wires map[string]uint16) {
	leftValue := handleOperand(ins.leftOperand, wires)
	rightValue := handleOperand(ins.rightOperand, wires)

	wires[ins.destination] = leftValue << rightValue
}

func executeRSHIFTOperation(ins Instruction, wires map[string]uint16) {
	leftValue := handleOperand(ins.leftOperand, wires)
	rightValue := handleOperand(ins.rightOperand, wires)

	wires[ins.destination] = leftValue >> rightValue
}

func executeAssignOperation(ins Instruction, wires map[string]uint16) {
	value := handleOperand(ins.leftOperand, wires)
	wires[ins.destination] = value
}

func handleOperand(op string, wires map[string]uint16) uint16 {
	if value, ok := wires[op]; ok {
		return value
	}

	value, err := strconv.Atoi(op)
	if err != nil {
		return 0
	}
	return uint16(value)
}

func parseInstructions(input []string) map[string]string {
	instructions := make(map[string]string)
	for _, line := range input {
		parts := strings.Split(line, " -> ")
		instructions[parts[1]] = parts[0]
	}
	return instructions
}

func evaluate(wire string, instructions map[string]string, cache map[string]uint16) uint16 {
	// Check if the wire is already evaluated
	if val, ok := cache[wire]; ok {
		return val
	}

	// Try to parse wire as a number
	if val, err := strconv.Atoi(wire); err == nil {
		return uint16(val)
	}

	// Get the instruction for this wire
	instruction := instructions[wire]
	parts := strings.Fields(instruction)

	var result uint16

	switch len(parts) {
	case 1: // Direct assignment
		result = evaluate(parts[0], instructions, cache)
	case 2: // NOT operation
		operand := evaluate(parts[1], instructions, cache)
		result = ^operand & 0xFFFF
	case 3: // Binary operations
		left := evaluate(parts[0], instructions, cache)
		right := evaluate(parts[2], instructions, cache)
		switch parts[1] {
		case "AND":
			result = left & right
		case "OR":
			result = left | right
		case "LSHIFT":
			shift, _ := strconv.Atoi(parts[2])
			result = left << shift
		case "RSHIFT":
			shift, _ := strconv.Atoi(parts[2])
			result = left >> shift
		}
	}

	// Cache the result
	cache[wire] = result
	return result
}

func main() {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(data), inputSep)
	// Parse instructions
	instructions := parseInstructions(input)

	// Cache for computed values
	cache := make(map[string]uint16)

	// Evaluate the signal for the desired wire
	originalSignal := evaluate("a", instructions, cache)
	fmt.Printf("Signal on wire a: %d\n", originalSignal)

	// Step 2: Override wire "b" with the signal from wire "a"
	instructions["b"] = fmt.Sprintf("%d", originalSignal)

	// Step 3: Reset cache and recompute
	cache = make(map[string]uint16)
	newSignal := evaluate("a", instructions, cache)
	fmt.Printf("New signal on wire a: %d\n", newSignal)
}
