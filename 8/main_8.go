package main

import "fmt"

func main() {
	wg := WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("Я горутина номер ", i)
		}(i)
	}
	wg.Wait()
	fmt.Println("Все горутины выполнились")
}

type WaitGroup struct {
	addSemaphore  chan struct{}
	doneSemaphore chan struct{}
	count         int
}

func (wg *WaitGroup) ensureInit() {
	if wg.addSemaphore != nil {
		return
	}
	wg.addSemaphore = make(chan struct{}, 1)
	wg.doneSemaphore = make(chan struct{})
}

func (wg *WaitGroup) Add(delta int) {
	wg.ensureInit()
	wg.addSemaphore <- struct{}{}
	wg.count += delta

	if wg.count < 0 {
		panic("WaitGroup: отрицательное количество")
	}

	if wg.count == 0 {
		close(wg.doneSemaphore)
		wg.doneSemaphore = make(chan struct{})
	}
	<-wg.addSemaphore
}

func (wg *WaitGroup) Done() {
	wg.Add(-1)
}

func (wg *WaitGroup) Wait() {
	wg.ensureInit()
	wg.addSemaphore <- struct{}{}
	if wg.count == 0 {
		<-wg.addSemaphore
		return
	}
	<-wg.addSemaphore
	<-wg.doneSemaphore
}
