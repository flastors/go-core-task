package main

import (
	"context"
	"testing"
)

func TestRandomGenerator(t *testing.T) {
	seed := 123
	min := 1
	max := 100
	ctx, cancel := context.WithCancel(context.Background())
	count := 25
	expected := []int{36, 50, 4, 36, 40, 40, 76, 16, 22, 12, 88, 3, 83, 16, 49, 2, 4, 38, 73, 19, 55, 95, 73, 58, 14}
	defer cancel()
	nums := RandomGenerator(ctx, int64(seed), min, max)
	result := make([]int, count)
	for i := 0; i < count; i++ {
		result[i] = <-nums
	}
	cancel()
	if _, ok := <-nums; ok {
		t.Errorf("Random generator cancel failed")
	}
	for i, v := range result {
		if expected[i] != v {
			t.Errorf("Got %v, Expected %v", result, expected)
		}
	}
}
