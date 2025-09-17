package test

import (
	"projects/exercises"
	"testing"
)

func TestMinOperations(t *testing.T) {
	got := exercises.MinimumOperations([] int {1,2,3,4})
	want := 3
	if got != want {
		t.Errorf("want %d, got %d", got, want)

	}

}
func TestMinOperations1(t *testing.T) {
	got := exercises.MinimumOperations([] int {3,6,9})
	want := 0
	if got != want {
		t.Errorf("want %d, got %d", got, want)

	}

}

