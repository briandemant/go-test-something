package main

import "testing"

func TestAdd2(t *testing.T) {
	var tests = []struct {
		input int
		expected int
	}{
		{2,4},
		{-1,1},
		{22,24},
		{99998,100000},
	}

	for _, test := range tests {
		actual := add2(test.input)
		if actual != test.expected {
			t.Errorf("add2(%v) expected %v got %v",test.input,test.expected,actual)
		}
	}
}

func TestAdd4(t *testing.T) {
	var tests = []struct {
		input int
		expected int
	}{
		{2,6},
		{-1,3},
		{22,26},
		{99998,100002},
	}

	for _, test := range tests {
		actual := add4(test.input)
		if actual != test.expected {
			t.Errorf("add2(%v) expected %v got %v",test.input,test.expected,actual)
		}
	}
}