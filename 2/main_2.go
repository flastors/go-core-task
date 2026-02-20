package main

import (
	"fmt"
	"math/rand"
)

func main() {
	originalSlice := make([]int, 10)
	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100)
	}
	fmt.Println("Оригинальный слайс: ", originalSlice)
	fmt.Println("Слайс только четных чисел: ", sliceExample(originalSlice))
	addNum := 999
	fmt.Printf("Слайс с добавлением числа %d: %v\n", addNum, addElement(originalSlice, addNum))
	copiedSlice := copySlice(originalSlice)
	fmt.Println("Копия слайса: ", copiedSlice)
	for i := range copiedSlice {
		copiedSlice[i] = 0
	}
	fmt.Println("Затер копию слайса нулями: ", copiedSlice)
	fmt.Println("При этом оригинальный слайс: ", originalSlice)
	indexToRemove := 5
	fmt.Printf("Удаляю элемент по индексу %d: %v\n", indexToRemove, removeElement(originalSlice, indexToRemove))
}

func sliceExample(s []int) []int {
	result := make([]int, 0, len(s))
	for _, i := range s {
		if i%2 == 0 {
			result = append(result, i)
		}
	}
	return result
}

func addElement(s []int, num int) []int {
	newSlice := make([]int, len(s)+1)
	copy(newSlice, s)
	newSlice[len(s)] = num
	return newSlice
}

func copySlice(slice []int) []int {
	if slice == nil {
		return nil
	}
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	return newSlice
}

func removeElement(s []int, index int) []int {
	if s == nil {
		return nil
	}
	newSlice := make([]int, len(s))
	copy(newSlice, s)
	if len(newSlice) <= index {
		if len(newSlice) < index {
			return newSlice
		}
		return newSlice[:index]
	}
	newSlice = append(newSlice[:index], newSlice[index+1:]...)
	return newSlice
}
