package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed input.txt
var input string

func ParseMemoryInput(input string) []string {
	r, _ := regexp.Compile(`mul\(\d+,\d+\)`)
	return r.FindAllString(input, -1)
}

func EvaluateInputs(inputCommands []string) (int, error) {
	total := 0
	r, _ := regexp.Compile(`\d+`)

	for _, command := range inputCommands {
		inputs := r.FindAllString(command, -1)
		numA, _ := strconv.Atoi(inputs[0])
		numB, _ := strconv.Atoi(inputs[1])
		total += numA * numB
	}
	return total, nil
}

func ProcessMemoryInput(input string) (int, error) {
	inputCommands := ParseMemoryInput(input)
	finalResult, err := EvaluateInputs(inputCommands)
	if err != nil {
		return 0, err
	}
	return finalResult, nil
}

func main() {
	fmt.Println("Parsing Corrupted Input")
	result, err := ProcessMemoryInput(input)
	if err != nil {
		fmt.Printf("An error has occured! %v", err)
	}
	fmt.Println("The final result is: ", result)
}
