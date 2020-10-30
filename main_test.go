package main

/**
1. Module is for testing utils
*/

import "testing"

func TestCalculate(t *testing.T) {
	if Calculate(1, 2) != 3 {
		t.Error("Expected 1 + 2 to be 3")
	}
}
