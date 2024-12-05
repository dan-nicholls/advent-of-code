package main

import (
	_ "embed"
	"reflect"
	"testing"
)

//go:embed test_input.txt
var testInput string
var testOutput int = 161

func TestParseMemoryInput(t *testing.T) {
	input := testInput
	expected := []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}

	result := ParseMemoryInput(input)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("ParseMemoryInput(%v) = %v; want %v", input, result, expected)
	}
}

func TestEvaluateInputs(t *testing.T) {
	input := []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}
	expected := 161

	result, err := EvaluateInputs(input)
	if err != nil {
		t.Errorf("EvaluteInputs(%v) returned an error: %v", input, err)
	}
	if result != expected {
		t.Errorf("EvaluteInputs(%v) = %v; want %v", input, result, expected)
	}
}

func TestProcessMemoryInput(t *testing.T) {
	input := testInput
	expected := testOutput

	result, err := ProcessMemoryInput(testInput)
	if err != nil {
		t.Errorf("ProcessMemoryInput(%v) returned an error: %v", input, err)
	}
	if result != expected {
		t.Errorf("EvaluateInput(%v) = %v; want %v", input, result, expected)
	}
}
