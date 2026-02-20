package main

import "testing"

func TestTypeOf(t *testing.T) {
	tests := []struct {
		name  string
		input any
		want  string
	}{
		{
			name:  "String Type",
			input: "typeString",
			want:  "string",
		},
		{
			name:  "Integer Type",
			input: 123,
			want:  "int",
		},
		{
			name:  "Uint Type",
			input: uint(123),
			want:  "uint",
		},
		{
			name:  "Slice Type",
			input: []int{1, 2, 3},
			want:  "[]int",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := typeOf(tt.input); got != tt.want {
				t.Errorf("typeOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHashSha256WithSalt(t *testing.T) {
	tests := []struct {
		name  string
		input string
		salt  string
		want  string
	}{
		{
			name:  "Basic Case",
			input: "stringToBeHashed",
			salt:  "someSalt",
			want:  "796E164DD33E8061B9206D5193B2B40F2C62E7CABF6789627F18403E77E92E58",
		},
		{
			name:  "Empty input",
			input: "",
			salt:  "someSalt",
			want:  "40FCAC02279CF685E942CDDF58C6B61265AFF3F90C58AE79B1029A0CC92D78F2",
		},
		{
			name:  "Empty Salt",
			input: "stringToBeHashed",
			salt:  "",
			want:  "E2BFA75CF177478E7C13E9A1B3467266B2D62F242421E7A62E8A050EDE959EB0",
		},
		{
			name:  "Empty input and salt",
			input: "",
			salt:  "",
			want:  "E3B0C44298FC1C149AFBF4C8996FB92427AE41E4649B934CA495991B7852B855",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hashSha256WithSalt(tt.input, tt.salt); got != tt.want {
				t.Errorf("hashSha256WithSalt() = %v, want %v", got, tt.want)
			}
		})
	}
}
