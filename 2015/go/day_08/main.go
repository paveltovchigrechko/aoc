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

	var codeLength, memoryLength int
	for _, line := range lines {
		codeLength += utf8.RuneCountInString(string(line))

		inMemoryString, err := strconv.Unquote(string(line))
		if err != nil {
			log.Print(err)
			continue
		}

		memoryLength += len(inMemoryString)
	}

	fmt.Printf("The difference between the number of characters of code for string literals minus the number of characters in memory for the values of the strings is %d\n", codeLength-memoryLength)
}
