package main

import (
	"testing"
	"time"
)

func TestFanInClose(t *testing.T) {
	ch := make(chan int)
	go func() {
		defer close(ch)
		ch <- 1
	}()
	<-ch
	mergedCh := fanIn(ch)
	select {
	case <-mergedCh:
	case <-time.After(100 * time.Millisecond):
		t.Error("Channel not closed after 100ms")
	}
}

func TestFanIn(t *testing.T) {
	tests := []struct {
		name  string
		input []chan int
		want  []int
	}{
		{
			name: "Basic Case",
			input: []chan int{
				make(chan int),
				make(chan int),
				make(chan int),
				make(chan int),
				make(chan int),
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name:  "Zero input channels",
			input: []chan int{},
			want:  []int{},
		},
		{
			name:  "Nil input channels",
			input: nil,
			want:  []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			for i, ch := range tt.input {
				go func(ch chan int, i int) {
					defer close(ch)
					ch <- i
				}(ch, i)
			}
			mergedCh := fanIn(tt.input...)
			result := make([]int, 0, len(tt.want))
			for v := range mergedCh {
				result = append(result, v)
			}
			if len(result) != len(tt.want) {
				t.Errorf("Got %v, want %v", result, tt.want)
				return
			}
			var found bool
			for _, v := range result {
				found = false
				for _, w := range tt.want {
					if v == w {
						found = true
					}
				}
				if !found {
					t.Errorf("Got %v, want %v", result, tt.want)
					return
				}
			}
		})
	}
}
