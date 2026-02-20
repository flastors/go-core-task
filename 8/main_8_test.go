package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestWaitGroupUsage(t *testing.T) {
	wg := WaitGroup{}
	wg.Add(1)
	var completed bool
	go func() {
		defer wg.Done()
		completed = true
	}()
	wg.Wait()
	if !completed {
		t.Error("Goroutine not complete")
	}
}

func TestWaitGroupConcurrency(t *testing.T) {
	wg := WaitGroup{}
	var counter int64 = 0
	var count int64 = 100
	var i int64
	for i = 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
		}()
	}
	wg.Wait()
	if counter != count {
		t.Errorf("Got %d, expected %d", counter, count)
	}
}

func TestWaitGroupNegative(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Panic expected")
		}
	}()
	wg := WaitGroup{}
	wg.Done()

}

func TestWaitGroupZeroCount(t *testing.T) {
	wg := WaitGroup{}
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Error("WaitGroup locked")
	}
}

func TestWaitGroupConcurrentInitialization(t *testing.T) {
	wg := WaitGroup{}

	for i := 0; i < 10; i++ {
		go wg.Add(1)
	}

	time.Sleep(50 * time.Millisecond)

	for i := 0; i < 10; i++ {
		go wg.Done()
	}

	wg.Wait()
}

func TestWaitGroupReuse(t *testing.T) {
	wg := WaitGroup{}
	counter := 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		counter++
	}()
	wg.Wait()
	wg.Add(1)
	go func() {
		defer wg.Done()
		counter++
	}()
	wg.Wait()
	if counter != 2 {
		t.Errorf("Got %d, expected 2", counter)
	}
}
