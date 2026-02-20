package main

import "testing"

func TestCrossValues(t *testing.T) {
	tests := []struct {
		name          string
		inputA        []int
		inputB        []int
		expectedSlice []int
		expectedBool  bool
	}{
		{
			name:          "Empty inputs",
			inputA:        []int{},
			inputB:        []int{},
			expectedSlice: []int{},
			expectedBool:  false,
		},
		{
			name:          "Empty one input",
			inputA:        []int{1, 2, 3},
			inputB:        []int{},
			expectedSlice: []int{},
			expectedBool:  false,
		},
		{
			name:          "Slice A bigger than Slice B",
			inputA:        []int{1, 2, 3, 4, 5, 6},
			inputB:        []int{4, 3},
			expectedSlice: []int{4, 3},
			expectedBool:  true,
		},
		{
			name:          "Slice B bigger than Slice A",
			inputA:        []int{1, 2, 3, 4, 5, 6},
			inputB:        []int{4, 3, 8, 9, 10, 11, 12},
			expectedSlice: []int{3, 4},
			expectedBool:  true,
		},
		{
			name:          "Slice A equal to Slice B",
			inputA:        []int{1, 2, 3, 4, 5, 6},
			inputB:        []int{1, 2, 3, 4, 5, 6},
			expectedSlice: []int{1, 2, 3, 4, 5, 6},
			expectedBool:  true,
		},
		{
			name:          "Zero cross values",
			inputA:        []int{1, 2, 3, 4, 5, 6},
			inputB:        []int{7, 8, 9, 10, 11, 12},
			expectedSlice: []int{},
			expectedBool:  false,
		},
		{
			name:          "Duplicate inputs",
			inputA:        []int{1, 2, 3, 4, 3, 5},
			inputB:        []int{1, 7, 3, 8, 9, 3},
			expectedSlice: []int{1, 3},
			expectedBool:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, ok := crossValues(tt.inputA, tt.inputB)
			if ok != tt.expectedBool {
				t.Errorf("ExpectedBool: got %v, want %v", ok, tt.expectedBool)
			}
			if len(output) != len(tt.expectedSlice) {
				t.Errorf("ExpectedSlice: got %v, want %v", len(output), len(tt.expectedSlice))
			} else {
				for i := range output {
					if output[i] != tt.expectedSlice[i] {
						t.Errorf("ExpectedSlice: got %v, want %v", output, tt.expectedSlice)
					}
				}
			}
		})
	}
}
