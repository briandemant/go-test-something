//go:build unit
// +build unit

package main

import (
	"testing"
)

func TestAdd2(t *testing.T) {
	var examples = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{22, 24},
		{99998, 100000},
	}

	for _, example := range examples {
		actual := Add2(example.input)
		if actual != example.expected {
			t.Errorf("Add2(%v) expected %v got %v", example.input, example.expected, actual)
		}
	}
}

func TestAdd4(t *testing.T) {
	var examples = []struct {
		input    int
		expected int
	}{
		{2, 6},
		{-1, 3},
		{22, 26},
		{99998, 100002},
	}

	for _, example := range examples {
		actual := Add4(example.input)
		if actual != example.expected {
			t.Errorf("Add2(%v) expected %v got %v", example.input, example.expected, actual)
		}
	}
}

type AddResult struct {
	x        int
	y        int
	expected int
}

var addExamples = []AddResult{
	{1, 1, 2},
	{1, 11, 12},
	{11, 11, 22},
}

func TestAdd(t *testing.T) {
	for _, example := range addExamples {
		actual := Add(example.x, example.y)
		if actual != example.expected {
			t.Errorf("Add(%v,%v) expected %v got %v", example.x, example.y, example.expected, actual)
		}
	}
}
