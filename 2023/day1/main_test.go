package main

import (
	"testing"
)

func TestFirstDigitOrNum(t *testing.T) {
	i, n := firstDigitOrNumber(nums, "2shhzch")
	if i != 0 {
		t.Errorf("Expected 0, got %v", n)
	}
	if n != 2 {
		t.Errorf("Expected 2, got %v", n)
	}

}

func TestFirstNum(t *testing.T) {
	i, n := firstNumber(nums, "2shhzch")
	if n != -1 {
		t.Errorf("Expected -1, got %v", n)
	}
	if i != -1 {
		t.Errorf("Expected -1, got %v", i)
	}

}

func TestFirstDigit(t *testing.T) {
	i, n := firstDigit("2shhzch")
	if n != 2 {
		t.Errorf("Expected 2, got %v", n)
	}
	if i != 0 {
		t.Errorf("Expected 0, got %v", i)
	}
}
