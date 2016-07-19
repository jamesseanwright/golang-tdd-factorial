package main

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	
	var tests = []struct{
		N int
		Fac int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{6, 720},
		{10, 3628800},
	}
	
	for _, test := range tests {
		f := GetFactorial(test.N)
		if f != test.Fac {
			t.Errorf("Factorial miscalculated - got %v, expected %v", f, test.Fac)
		}
	}
}
