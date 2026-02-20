package main

import "testing"

func TestSliceExample(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Basic Case",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "Zero Even Numbers",
			input:    []int{1, 3, 5, 7, 9, 11},
			expected: []int{},
		},
		{
			name:     "Only Even Numbers",
			input:    []int{2, 4, 6, 8, 10},
			expected: []int{2, 4, 6, 8, 10},
		},
		{
			name:     "Nil Input",
			input:    nil,
			expected: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sliceExample(tt.input)
			if len(result) != len(tt.expected) {
				t.Errorf("Got: %v, Expected: %v", result, tt.expected)
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Got: %v, Expected: %v", result, tt.expected)
					return
				}
			}
		})
	}
}

func TestSliceExampleChange(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := sliceExample(input)
	for i := range result {
		result[i] = 0
	}
	for i := range input {
		if input[i] != expected[i] {
			t.Error("Changes in result slice affecting input slice")
		}
	}
}

func TestAddElement(t *testing.T) {
	tests := []struct {
		name          string
		inputSlice    []int
		inputElement  int
		expectedSlice []int
	}{
		{
			name:          "Basic Case",
			inputSlice:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			inputElement:  5,
			expectedSlice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 5},
		},
		{
			name:          "Zero Input Slice",
			inputSlice:    []int{},
			inputElement:  5,
			expectedSlice: []int{5},
		},
		{
			name:          "Nil Input",
			inputSlice:    nil,
			inputElement:  5,
			expectedSlice: []int{5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addElement(tt.inputSlice, tt.inputElement)
			if len(result) != len(tt.expectedSlice) {
				t.Errorf("Got: %v, Expected: %v", result, tt.expectedSlice)
				return
			}
			for i := range result {
				if result[i] != tt.expectedSlice[i] {
					t.Errorf("Got: %v, Expected: %v", result, tt.expectedSlice)
					return
				}
			}
		})
	}
}

func TestAddElementAffecting(t *testing.T) {
	inputSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	inputElement := 5
	expectedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := addElement(inputSlice, inputElement)
	for i := range result {
		result[i] = 0
	}
	for i := range inputSlice {
		if inputSlice[i] != expectedSlice[i] {
			t.Error("Changes in result slice affecting input slice")
		}
	}
}

func TestCopySlice(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Basic Case",
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:     "Zero Input Slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "Nil Input",
			input:    nil,
			expected: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := copySlice(tt.input)
			if result == nil {
				if tt.expected == nil {
					return
				}
				t.Errorf("Got: %v, Expected: %v", result, tt.expected)
				return
			}
			if len(result) != len(tt.expected) {
				t.Errorf("Got: %v, Expected: %v", result, tt.expected)
				return
			}
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("Got: %v, Expected: %v", result, tt.expected)
					return
				}
			}
		})
	}
}

func TestCopySliceAffecting(t *testing.T) {
	inputSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expectedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := copySlice(inputSlice)
	for i := range result {
		result[i] = 0
	}
	for i := range inputSlice {
		if inputSlice[i] != expectedSlice[i] {
			t.Error("Changes in result slice affecting input slice")
			return
		}
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name          string
		inputSlice    []int
		inputIndex    int
		expectedSlice []int
	}{
		{
			name:          "Basic Case",
			inputSlice:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			inputIndex:    5,
			expectedSlice: []int{1, 2, 3, 4, 5, 7, 8, 9, 10},
		},
		{
			name:          "Zero Input Slice",
			inputSlice:    []int{},
			inputIndex:    5,
			expectedSlice: []int{},
		},
		{
			name:          "Single Element",
			inputSlice:    []int{1},
			inputIndex:    0,
			expectedSlice: []int{},
		},
		{
			name:          "Remove First Element",
			inputSlice:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			inputIndex:    0,
			expectedSlice: []int{2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:          "Remove Last Element",
			inputSlice:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			inputIndex:    9,
			expectedSlice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			name:          "Index out of range",
			inputSlice:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			inputIndex:    15,
			expectedSlice: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
		{
			name:          "Nil Input",
			inputSlice:    nil,
			inputIndex:    0,
			expectedSlice: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeElement(tt.inputSlice, tt.inputIndex)
			if result == nil {
				if tt.expectedSlice == nil {
					return
				}
				t.Errorf("Got: %v, Expected: %v", result, tt.expectedSlice)
				return
			}
			if len(result) != len(tt.expectedSlice) {
				t.Errorf("Got: %v, Expected: %v", result, tt.expectedSlice)
				return
			}
			for i := range result {
				if result[i] != tt.expectedSlice[i] {
					t.Errorf("Got: %v, Expected: %v", result, tt.expectedSlice)
					return
				}
			}
		})
	}
}

func TestRemoveElementAffecting(t *testing.T) {
	inputSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	inputIndex := 5
	expectedSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := removeElement(inputSlice, inputIndex)
	for i := range result {
		result[i] = 0
	}
	for i := range inputSlice {
		if inputSlice[i] != expectedSlice[i] {
			t.Error("Changes in result slice affecting input slice")
			return
		}
	}
}
