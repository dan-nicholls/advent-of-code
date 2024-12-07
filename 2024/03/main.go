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
	r, err := regexp.Compile(`mul\(\d+,\d+\)|do\(\)|don\'t\(\)`)
	if err != nil {
		fmt.Println("Error in Reg: ", err)
	}
	return r.FindAllString(input, -1)
}

func EvaluateInputs(inputCommands []string) (int, error) {
	total := 0
	isEnabled := true
	commandReg, _ := regexp.Compile(`mul|don't|do`)
	numReg, _ := regexp.Compile(`\d+`)

	for _, command := range inputCommands {
		res := commandReg.FindString(command)
		switch res {
		case "mul":
			if !isEnabled {
				continue
			}

			inputs := numReg.FindAllString(command, -1)
			numA, _ := strconv.Atoi(inputs[0])
			numB, _ := strconv.Atoi(inputs[1])
			total += numA * numB
		case "don't":
			isEnabled = false
		case "do":
			isEnabled = true
		default:
			fmt.Println("INVALID COMMAND: ", command)
		}
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
