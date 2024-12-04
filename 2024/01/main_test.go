package main

import (
	_ "embed"
	"reflect"
	"testing"
)

var (
	exampleList1      = []int{3, 4, 2, 1, 3, 3}
	exampleList2      = []int{4, 3, 5, 3, 9, 3}
	exampleListOutput = 11
)

//go:embed test_input.txt
var testInput string

func Test_InputParse(t *testing.T) {
	tests := []struct {
		name  string
		input string
		listA []int
		listB []int
	}{
		{
			name:  "Parse Input String Test",
			input: testInput,
			listA: exampleList1,
			listB: exampleList2,
		},
	}

	for _, tt := range tests {
		resultListA, resultListB, error := ParseInputLists(tt.input)
		if error != nil {
			t.Errorf("ParseInputLists(%v) returned an error: %v", tt.input, error)
			continue
		}

		if !reflect.DeepEqual(resultListA, tt.listA) {
			t.Errorf(
				"ParseInputLists(%v) = %v; want list A to be %v",
				tt.input,
				resultListA,
				tt.listA,
			)
		}

		if !reflect.DeepEqual(resultListB, tt.listB) {
			t.Errorf(
				"ParseInputLists(%v) = %v; want list B to be %v",
				tt.input,
				resultListB,
				tt.listB,
			)
		}
	}
}

func Test_TotalDif(t *testing.T) {
	tests := []struct {
		name     string
		listA    []int
		listB    []int
		expected int
	}{
		{
			name:     "Find Total Difference Test",
			listA:    exampleList1,
			listB:    exampleList2,
			expected: exampleListOutput,
		},
	}

	for _, tt := range tests {
		result := FindTotalDif(tt.listA, tt.listB)
		if result != tt.expected {
			t.Errorf("TotalDif(%d, %d) = %d; want %d", tt.listA, tt.listB, result, tt.expected)
		}
	}
}
