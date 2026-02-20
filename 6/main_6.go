package main

import (
	"context"
	"fmt"
	"math/rand"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	randomNums := RandomGenerator(ctx, 423432, 1, 10)
	for i := 0; i < 10; i++ {
		fmt.Println(<-randomNums)
	}
	cancel()
}

func RandomGenerator(ctx context.Context, seed int64, min, max int) chan int {
	rnd := rand.New(rand.NewSource(seed))
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				select {
				case <-ctx.Done():
					return
				case out <- rnd.Intn(max-min+1) + min:
				}
			}
		}
	}()
	return out
}
