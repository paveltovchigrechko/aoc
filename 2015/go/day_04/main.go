package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	inputFile = "input_04.txt"
	prefix_01 = "00000"
	prefix_02 = "000000"
)

func findMD5Complement(secret, prefix string) int {
	i := 1
	for {
		numStr := strconv.Itoa(i)
		bytes := []byte(secret + numStr)
		hashed := md5.Sum(bytes)
		hashedStr := hex.EncodeToString(hashed[:])
		if strings.HasPrefix(hashedStr, prefix) {
			return i
		}
		i++
	}
}

func main() {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	input := string(data)
	fmt.Printf("The lowest positive number to produce an MD5 hash with '%s' prefix is %d\n", prefix_01, findMD5Complement(input, prefix_01))
	fmt.Printf("The lowest positive number to produce an MD5 hash with '%s' prefix is %d\n", prefix_02, findMD5Complement(input, prefix_02))
}
