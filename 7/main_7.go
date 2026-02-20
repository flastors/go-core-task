package main

import (
	"fmt"
	"sync"
)

func main() {
	channels := make([]chan int, 5)
	for i := range channels {
		channels[i] = make(chan int)
		go func(c chan int, i int) {
			defer close(c)
			for j := 0; j <= i; j++ {
				c <- j
			}
		}(channels[i], i)
	}
	mergedCh := fanIn(channels...)
	for v := range mergedCh {
		fmt.Println(v)
	}
}

func fanIn(channels ...chan int) <-chan int {
	mergedChan := make(chan int)
	wg := &sync.WaitGroup{}

	worker := func(in <-chan int) {
		defer wg.Done()
		for n := range in {
			mergedChan <- n
		}
	}

	wg.Add(len(channels))

	for _, c := range channels {
		go worker(c)
	}

	go func() {
		wg.Wait()
		close(mergedChan)
	}()
	return mergedChan
}
