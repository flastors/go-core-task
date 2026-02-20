package main

import "testing"

func TestStringIntMap_Add(t *testing.T) {
	m := NewStringIntMap()
	inputKey := "key"
	inputValue := 123
	m.Add(inputKey, inputValue)
	if val, ok := m.value[inputKey]; !ok || val != inputValue {
		t.Errorf("Got: %v, Expected: %v", val, inputValue)
	}
}

func TestStringIntMap_Remove(t *testing.T) {
	tests := []struct {
		name     string
		inputMap map[string]int
		inputKey string
		expected map[string]int
	}{
		{
			name: "Basic Case",
			inputMap: map[string]int{
				"key":   123,
				"test":  404,
				"test2": 200,
			},
			inputKey: "key",
			expected: map[string]int{
				"test":  404,
				"test2": 200,
			},
		},
		{
			name: "Not Existing Key",
			inputMap: map[string]int{
				"key":   123,
				"test":  404,
				"test2": 200,
			},
			inputKey: "key2",
			expected: map[string]int{
				"key":   123,
				"test":  404,
				"test2": 200,
			},
		},
		{
			name:     "Empty Map",
			inputMap: map[string]int{},
			inputKey: "key",
			expected: map[string]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewStringIntMap()
			m.value = tt.inputMap
			m.Remove(tt.inputKey)
			if len(m.value) != len(tt.expected) {
				t.Errorf("Got: %v, Expected: %v", tt.expected, tt.expected)
				return
			}
			for k, v := range m.value {
				if v != tt.expected[k] {
					t.Errorf("Got: %v, Expected: %v", tt.expected, tt.expected)
					return
				}
			}
		})
	}
}

func TestStringIntMap_Exists(t *testing.T) {
	tests := []struct {
		name     string
		inputMap map[string]int
		inputKey string
		expected bool
	}{
		{
			name: "Key Exists",
			inputMap: map[string]int{
				"key":   123,
				"test":  404,
				"test2": 200,
			},
			inputKey: "key",
			expected: true,
		},
		{
			name: "Key Not Exists",
			inputMap: map[string]int{
				"key":   123,
				"test":  404,
				"test2": 200,
			},
			inputKey: "key2",
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewStringIntMap()
			m.value = tt.inputMap
			result := m.Exists(tt.inputKey)
			if result != tt.expected {
				t.Errorf("Got: %v, Expected: %v", result, tt.expected)
			}
		})
	}
}

func TestStringIntMap_Get(t *testing.T) {
	tests := []struct {
		name           string
		inputMap       map[string]int
		inputKey       string
		expectedValue  int
		expectedExists bool
	}{
		{
			name: "Element Exists",
			inputMap: map[string]int{
				"key":   123,
				"test":  404,
				"test2": 200,
			},
			inputKey:       "key",
			expectedValue:  123,
			expectedExists: true,
		},
		{
			name: "Element Not Exists",
			inputMap: map[string]int{
				"key":   123,
				"test":  404,
				"test2": 200,
			},
			inputKey:       "key2",
			expectedValue:  0,
			expectedExists: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := NewStringIntMap()
			m.value = tt.inputMap
			val, ok := m.Get(tt.inputKey)
			if ok != tt.expectedExists {
				t.Errorf("expectedExists: Got: %v, Expected: %v", ok, tt.expectedExists)
				return
			}
			if val != tt.expectedValue {
				t.Errorf("expectedValue: Got: %v, Expected: %v", val, tt.expectedValue)
				return
			}
		})
	}
}

func TestStringIntMap_Copy(t *testing.T) {
	inputMap := map[string]int{
		"key":   123,
		"test":  404,
		"test2": 200,
	}
	expectedMap := map[string]int{
		"key":   123,
		"test":  404,
		"test2": 200,
	}
	m := NewStringIntMap()
	m.value = inputMap
	result := m.Copy()
	if len(result) != len(expectedMap) {
		t.Errorf("Got: %v, Expected: %v", result, expectedMap)
		return
	}
	for k, v := range result {
		if v != expectedMap[k] {
			t.Errorf("Got: %v, Expected: %v", result, expectedMap)
		}
	}
	for k := range result {
		result[k] = 0
	}
	for k, v := range expectedMap {
		if v != inputMap[k] {
			t.Errorf("Result map affecting input map")
		}
	}
}
