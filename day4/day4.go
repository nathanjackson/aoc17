package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

func NoDuplicateWords(candidate string) bool {
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

func NoAnagrams(candidate string) bool {
	words := strings.Fields(candidate)
	// find the lengths of each word
	lengthMap := map[int][]string{}
	for _, word := range words {
		wl := len(word)
		if _, ok := lengthMap[wl]; !ok {
			lengthMap[wl] = []string{}
		}
		lengthMap[wl] = append(lengthMap[wl], word)
	}
	// process lists where the word list length is > 1
	for length, list := range lengthMap {
		entryHistograms := []map[rune]int{}
		if length > 1 {
			for _, word := range list {
				hist := Histogram(word)
				for _, knownHist := range entryHistograms {
					if reflect.DeepEqual(hist, knownHist) {
						return false
					}
				}
				entryHistograms = append(entryHistograms, hist)

			}
		}
	}
	return true
}

func Histogram(in string) map[rune]int {
	result := map[rune]int{}
	for _, r := range in {
		if _, ok := result[r]; !ok {
			result[r] = 0
		}
		result[r]++
	}
	return result
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

	noDuplicatesCount := 0
	noAnagramsCount := 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			fmt.Printf("encountered an error while reading: %v", err)
			os.Exit(1)
		}

		if NoDuplicateWords(scanner.Text()) {
			noDuplicatesCount++
		}

		if NoAnagrams(scanner.Text()) {
			noAnagramsCount++
		}
	}

	fmt.Printf("passphrases without duplicate words: %v\n", noDuplicatesCount)
	fmt.Printf("passphrases without anagrams: %v\n", noAnagramsCount)
}
