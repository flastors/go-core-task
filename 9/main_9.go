package main

import "fmt"

func main() {
	in, out := make(chan uint8), make(chan float64)
	inNums := []uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Вошло ", inNums)
	go pipeline(in, out)
	go func() {
		for _, i := range inNums {
			in <- i
		}
		close(in)
	}()
	result := make([]float64, 0, len(inNums))
	for o := range out {
		result = append(result, o)
	}
	fmt.Println("Вышло", result)
}

func pipeline(in <-chan uint8, out chan<- float64) {
	defer close(out)
	in2 := pipeUint8ToFloat64(in)
	out2 := pipeCubeFloat64(in2)
	for o := range out2 {
		out <- o
	}
}

func pipeUint8ToFloat64(in <-chan uint8) <-chan float64 {
	out := make(chan float64)
	go func() {
		defer close(out)
		for v := range in {
			out <- float64(v)
		}
	}()
	return out
}

func pipeCubeFloat64(in <-chan float64) <-chan float64 {
	out := make(chan float64)
	go func() {
		defer close(out)
		for v := range in {
			out <- v * v * v
		}
	}()
	return out
}
