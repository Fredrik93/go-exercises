package test

import (
	"projects/exercises"
	"testing"
)

func TestClosestFriend(t *testing.T) {
	got := exercises.FindClosest(2, 7, 4)
	want := 1
	if got != want {
		t.Errorf("want %d, got %d", got, want)

	}

}
func TestClosestFriend1(t *testing.T) {
	got := exercises.FindClosest(2, 5, 6)
	want := 2
	if got != want {
		t.Errorf("want %d, got %d", got, want)
	}

}
func TestClosestFriend2(t *testing.T) {
	got := exercises.FindClosest(1, 5, 3)
	want := 0
	if got != want {
		t.Errorf("want %d, got %d", got, want)
	}

}
