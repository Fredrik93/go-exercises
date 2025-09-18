package test

import (
	"projects/exercises"
	"testing"
)

func TestPalindrome1(t *testing.T) {
	got := exercises.IsPalindrome(121)
	want := true
	if got != want {
		t.Errorf("Expected %t, got %t", want, got)
	}
}

func TestPalindrome2(t *testing.T) {
	got := exercises.IsPalindrome(-121)
	want := false
	if got != want {
		t.Errorf("Expected %t, got %t", want, got)
	}
}

func TestPalindrome3(t *testing.T) {
	got := exercises.IsPalindrome(10)
	want := false
	if got != want {
		t.Errorf("Expected %t, got %t", want, got)
	}
}
