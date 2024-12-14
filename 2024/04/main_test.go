package main

import (
	_ "embed"
	"testing"
)

//go:embed test_input.txt
var testInput string

var (
	testOutput  int = 18
	testOutput2 int = 9
)

func TestEvaluateWordSearch(t *testing.T) {
	input := testInput
	expected := testOutput

	result, err := EvaluateWordSearch(input, false)
	if err != nil {
		t.Errorf("EvaluateWordSearch(%v) encountered an error: %v", input, err)
	}
	if result != expected {
		t.Errorf("EvaluateWordSearch(%v) = %v; want %v", input, result, expected)
	}
}

func TestEvaluateCrossWordSearch(t *testing.T) {
	input := testInput
	expected := testOutput2

	result, err := EvaluateWordSearch(input, true)
	if err != nil {
		t.Errorf("EvaluateWordSearch(%v) encountered an error: %v", input, err)
	}
	if result != expected {
		t.Errorf("EvaluateWordSearch(%v) = %v; want %v", input, result, expected)
	}
}
