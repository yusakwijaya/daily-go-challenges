package main

import (
	"testing"
)

func TestFizzBuzzValue(t *testing.T) {
	tests := map[int]string{
		1:  "1",
		3:  "Fizz",
		5:  "Buzz",
		15: "FizzBuzz",
		30: "FizzBuzz",
		13: "13",
	}

	for input, expected := range tests {
		got := fizzBuzzValue(input)
		if got != expected {
			t.Errorf("For %d, expected %s but got %s", input, expected, got)
		}
	}
}
