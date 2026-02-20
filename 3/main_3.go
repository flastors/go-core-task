package main

import "fmt"

func main() {
	mapochka := NewStringIntMap()
	mapochka.Add("first", 1)
	mapochka.Add("second", 2)
	mapochka.Add("third", 3)
	mapochka.Remove("first")
	fmt.Println("Существует ли значение first: ", mapochka.Exists("first"))
	if v, ok := mapochka.Get("second"); ok {
		fmt.Println("Second = ", v)
	}
	copyMapochka := mapochka.Copy()
	fmt.Println(copyMapochka)

}

type StringIntMap struct {
	value map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{make(map[string]int)}
}

func (s *StringIntMap) Add(key string, value int) {
	s.value[key] = value
}

func (s *StringIntMap) Remove(key string) {
	delete(s.value, key)
}

func (s *StringIntMap) Exists(key string) bool {
	_, ok := s.value[key]
	return ok
}

func (s *StringIntMap) Get(key string) (int, bool) {
	v, ok := s.value[key]
	return v, ok
}

func (s *StringIntMap) Copy() map[string]int {
	m := make(map[string]int)
	for k, v := range s.value {
		m[k] = v
	}
	return m
}
