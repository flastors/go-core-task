package main

import "testing"

func TestPipeline(t *testing.T) {
	tests := []struct {
		name     string
		input    []uint8
		expected []float64
	}{
		{
			name:     "Empty input",
			input:    []uint8{},
			expected: []float64{},
		},
		{
			name:     "One input",
			input:    []uint8{3},
			expected: []float64{27},
		},
		{
			name:     "Multiple inputs",
			input:    []uint8{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			expected: []float64{1, 8, 27, 64, 125, 216, 343, 512, 729, 1000},
		},
		{
			name:     "Max uint",
			input:    []uint8{255},
			expected: []float64{16581375},
		},
		{
			name:     "Zero input",
			input:    []uint8{0},
			expected: []float64{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			in := make(chan uint8)
			out := make(chan float64)
			go pipeline(in, out)
			go func() {
				for _, i := range tt.input {
					in <- i
				}
				close(in)
			}()
			results := make([]float64, 0, len(tt.expected))
			for i := range out {
				results = append(results, i)
			}
			if len(results) != len(tt.expected) {
				t.Errorf("Invalid length got %v, want %v", results, tt.expected)
			}
			for i, exp := range tt.expected {
				if results[i] != exp {
					t.Errorf("Invalid results got %v, want %v", results, tt.expected)
				}
			}
		})
	}
}

func TestClosingChannel(t *testing.T) {
	in := make(chan uint8)
	out := make(chan float64)
	go pipeline(in, out)
	go func() {
		in <- 5
		close(in)
	}()
	_ = <-out
	for _ = range out {

	}
}
