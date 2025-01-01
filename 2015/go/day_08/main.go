package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode/utf8"
)

const (
	inputFile = "input_08.txt"
)

func main() {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	lines := bytes.Split(content, []byte("\n"))

	var codeLength, memoryLength, encodedLength int
	for _, line := range lines {
		codeLength += utf8.RuneCount(line)

		inMemoryString, err := strconv.Unquote(string(line))
		if err != nil {
			log.Print(err)
			continue
		}

		memoryLength += len(inMemoryString)

		encoded := fmt.Sprintf("%q", string(line))
		encodedLength += len(encoded)
	}

	fmt.Printf("The difference between the number of characters of code for string literals minus the number of characters in memory for the values of the strings is %d\n", codeLength-memoryLength)
	fmt.Printf("The the total number of characters to represent the newly encoded strings minus the number of characters of code in each original string literal is %d\n", encodedLength-codeLength)
}
