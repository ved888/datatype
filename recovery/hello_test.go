package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add1(5, 9)
	expected := 14
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
