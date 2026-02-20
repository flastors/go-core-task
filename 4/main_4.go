package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "apple", "banana", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	fmt.Println(DivideSlices(slice1, slice2))
}

func DivideSlices(a []string, b []string) []string {
	result := make([]string, 0, len(a))
	var found bool
	for _, i := range a {
		found = false
		for _, k := range result {
			if i == k {
				found = true
				break
			}
		}
		if !found {
			for _, j := range b {
				if i == j {
					found = true
					break
				}
			}
			if !found {
				result = append(result, i)
			}
		}
	}
	return result
}
