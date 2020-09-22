package main

import "testing"

func TestMin(t *testing.T) {
	expected := 2
	result := Min(2, 3)
	if result != expected {
		t.Errorf("Expected %d, but got %d\n", expected, result)
	}
}
