package main

import (
	_ "embed"
	"testing"
)

//go:embed test_input.txt
var testInput string

func TestIsReportSafe(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Passing decreasing report",
			input:    "7 6 4 2 1",
			expected: true,
		},
		{
			name:     "Passing increasing report",
			input:    "1 3 6 7 9",
			expected: true,
		},
		{
			name:     "Failing increasing report",
			input:    "1 2 7 8 9",
			expected: false,
		},
		{
			name:     "Failing decreasing report",
			input:    "9 7 6 2 1",
			expected: false,
		},
		{
			name:     "Failing other report",
			input:    "8 6 4 4 1",
			expected: false,
		},
	}

	for _, tt := range tests {
		result, err := IsReportSafe(tt.input)
		if err != nil {
			t.Errorf("IsReportSafe(%v) returned an error: %v", tt.input, err)
		}
		if result != tt.expected {
			t.Errorf("IsReportSafe(%s) = %t; want %t", tt.input, result, tt.expected)
		}
	}
}

func TestCountSafeReports(t *testing.T) {
	input := ParseInputReports(testInput)
	expected := 2

	result := CountSafeReports(input)
	if result != expected {
		t.Errorf("CountSafeReports(%v) = %v; want %v", input, result, expected)
	}
}
