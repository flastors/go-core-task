package main

import "testing"

func TestDivideSlices(t *testing.T) {
	tests := []struct {
		name     string
		inputA   []string
		inputB   []string
		expected []string
	}{
		{
			name:     "Basic Case",
			inputA:   []string{"banana", "cherry", "polygon", "apple"},
			inputB:   []string{"banana", "cherry", "apple"},
			expected: []string{"polygon"},
		},
		{
			name:     "Empty inputs",
			inputA:   []string{},
			inputB:   []string{},
			expected: []string{},
		},
		{
			name:     "Empty input A",
			inputA:   []string{},
			inputB:   []string{"banana", "cherry", "polygon", "apple"},
			expected: []string{},
		},
		{
			name:     "Empty input B",
			inputA:   []string{"banana", "cherry", "polygon", "apple"},
			inputB:   []string{},
			expected: []string{"banana", "cherry", "polygon", "apple"},
		},
		{
			name:     "Equal inputs",
			inputA:   []string{"banana", "cherry", "polygon", "apple"},
			inputB:   []string{"banana", "cherry", "polygon", "apple"},
			expected: []string{},
		},
		{
			name:     "Duplicate in input A",
			inputA:   []string{"banana", "cherry", "polygon", "apple", "banana", "cherry", "polygon", "apple"},
			inputB:   []string{"banana", "cherry", "apple"},
			expected: []string{"polygon"},
		},
		{
			name:     "Duplicate in input B",
			inputA:   []string{"banana", "cherry", "polygon", "apple"},
			inputB:   []string{"banana", "cherry", "apple", "banana", "cherry", "apple"},
			expected: []string{"polygon"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DivideSlices(tt.inputA, tt.inputB)
			if len(result) != len(tt.expected) {
				t.Errorf("expected %v, got %v", len(tt.expected), len(result))
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}
