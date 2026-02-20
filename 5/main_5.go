package main

import "fmt"

// Еще бывают дубликаты, что в первом, что во втором слайсах
// Но в результате их быть не должно
func main() {
	a := []int{65, 3, 58, 678, 64, 3, 987, 2, 21}
	b := []int{64, 2, 3, 43, 33, 14, 28, 3}
	fmt.Println(crossValues(a, b))
}

func crossValues(a, b []int) ([]int, bool) {
	length := len(a)
	minus, maxus := a, b
	if length > len(b) {
		length = len(b)
		minus, maxus = maxus, minus
	}
	result := make([]int, 0, length)
	var duplicate bool
	for _, el := range minus {
		duplicate = false
		for _, resEl := range result {
			if el == resEl {
				duplicate = true
				break
			}
		}
		if duplicate {
			continue
		}
		for _, el2 := range maxus {
			if el == el2 {
				result = append(result, el)
				break
			}
		}
	}
	if len(result) != 0 {
		return result, true
	}
	return result, false

}
