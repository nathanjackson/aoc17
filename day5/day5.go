package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func EscapeMaze(instructions []int) int {
	pc := 0
	stepCount := 0
	for {
		jmp := instructions[pc]
		instructions[pc]++
		pc += jmp
		stepCount++
		if pc < 0 || pc >= len(instructions) {
			break
		}
	}
	return stepCount
}

func EscapeMazeP2(instructions []int) int {
	pc := 0
	stepCount := 0
	for {
		jmp := instructions[pc]
		if jmp >= 3 {
			instructions[pc]--
		} else {
			instructions[pc]++
		}
		pc += jmp
		stepCount++
		if pc < 0 || pc >= len(instructions) {
			break
		}
	}
	return stepCount
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: %v <input file>\n", os.Args[0])
	}

	inputFilePath := os.Args[1]
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		fmt.Printf("could not open file: %v\n", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	var instrs []int
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		err = scanner.Err()
		if err != nil {
			fmt.Printf("could not read file: %v\n", err)
			os.Exit(1)
		}

		instr, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("could not parse instruction: %v\n", err)
		}
		instrs = append(instrs, instr)
	}
	instrsP2 := make([]int, len(instrs))
	copy(instrsP2, instrs)

	stepCountP1 := EscapeMaze(instrs)
	stepCountP2 := EscapeMazeP2(instrsP2)
	fmt.Printf("part one: %v\n", stepCountP1)
	fmt.Printf("part two: %v\n", stepCountP2)
}
