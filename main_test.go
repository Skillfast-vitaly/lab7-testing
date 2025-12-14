package main

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{5, 3, 8, 1}, []int{1, 3, 5, 8}},
		{[]int{9, 8, 7, 6}, []int{6, 7, 8, 9}},
		{[]int{100}, []int{100}},
		{[]int{-3, 0, -5}, []int{-5, -3, 0}},
		{[]int{}, []int{}},
	}

	for _, test := range tests {
		bubbleSort(test.input)
		for i := range test.input {
			if test.input[i] != test.expected[i] {
				t.Errorf("Got %v, expected %v", test.input, test.expected)
				break
			}
		}
	}
}

func TestPalindrome(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"racecar", true},
		{"hello", false},
		{"A", true},
		{"", true},
		{"12321", true},
	}

	for _, test := range tests {
		result := palindrome(test.input)
		if result != test.expected {
			t.Errorf("Input: %s, Got %t, Expected %t", test.input, result, test.expected)
		}
	}
}

func TestFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 1},
		{5, 120},
		{7, 5040},
		{10, 3628800},
	}

	for _, test := range tests {
		result := factorial(test.input)
		if result != test.expected {
			t.Errorf("Input: %d, Got %d, Expected %d", test.input, result, test.expected)
		}
	}
}

func TestFibonacci(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 0},
		{1, 1},
		{5, 5},
		{7, 13},
		{10, 55},
	}

	for _, test := range tests {
		result := fibonacci(test.input)
		if result != test.expected {
			t.Errorf("Input: %d, Got %d, Expected %d", test.input, result, test.expected)
		}
	}
}

func TestFindSubString(t *testing.T) {
	tests := []struct {
		str      string
		substr   string
		expected bool
	}{
		{"hello world", "world", true},
		{"golang", "lang", true},
		{"test", "xyz", false},
		{"", "empty", false},
		{"case", "ca", true},
	}

	for _, test := range tests {
		result := findSubString(test.str, test.substr)
		if result != test.expected {
			t.Errorf("Str: %q, Substr: %q, Got %t, Expected %t", test.str, test.substr, result, test.expected)
		}
	}
}

func TestNumberPrime(t *testing.T) {
	tests := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{4, false},
		{17, true},
		{1, false},
		{7919, true},
	}

	for _, test := range tests {
		result := numberPrime(test.input)
		if result != test.expected {
			t.Errorf("Input: %d, Got %t, Expected %t", test.input, result, test.expected)
		}
	}
}

func TestReverseMath32(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{123, 321},
		{-456, -654},
		{0, 0},
		{1534236469, 0},
		{100, 1},
	}

	for _, test := range tests {
		result := reverseMath32(test.input)
		if result != test.expected {
			t.Errorf("Input: %d, Got %d, Expected %d", test.input, result, test.expected)
		}
	}
}

func TestIntToRim(t *testing.T) {
	tests := []struct {
		input    int
		expected string
	}{
		{1, "I"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}

	for _, test := range tests {
		result := intToRim(test.input)
		if result != test.expected {
			t.Errorf("Input: %d, Got %q, Expected %q", test.input, result, test.expected)
		}
	}
}
