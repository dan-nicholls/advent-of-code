package main

import (
	_ "embed"
	"testing"
)

//go:embed test_input.txt
var testInput string
var testOutput int = 18

func TestEvaluateWordSearch(t *testing.T) {
	input := testInput
	expected := testOutput

	result, err := EvaluateWordSearch(input)
	if err != nil {
		t.Errorf("EvaluateWordSearch(%v) encountered an error: %v", input, err)
	}
	if result != expected {
		t.Errorf("EvaluateWordSearch(%v) = %v; want %v", input, result, expected)
	}
}
