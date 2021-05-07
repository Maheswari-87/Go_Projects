package main

import "testing"

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("Expected 2+2 is 4")
	}
}

func TestTableCalculate(t *testing.T) {
	var table = []struct {
		input    int
		expected int
	}{
		{2, 4},
		{-1, 1},
		{4, 6},
		{9999, 10001},
	}
	for _, test := range table {
		if output := Calculate(test.input); output != test.expected {
			t.Error("test failed: {} inputted, {} expected, received {}", test.input, test.expected, output)
		}
	}
}
