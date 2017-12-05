package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func IsValid(candidate string) bool {
	words := strings.Fields(candidate)
	wordSet := map[string]bool{}
	for _, word := range words {
		if _, hasWord := wordSet[word]; hasWord {
			return false
		}
		wordSet[word] = true
	}
	return true
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %s <input file>\n", os.Args[0])
		os.Exit(1)
	}

	inputFilePath := os.Args[1]
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Printf("could not open input file: %v\n", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	validCount := 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Printf("encountered an error while reading: %v", err)
			os.Exit(1)
		}

		if IsValid(scanner.Text()) {
			validCount++
		}
	}

	fmt.Printf("%v\n", validCount)
}
