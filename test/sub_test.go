package test

import ("testing"
		"projects/exercises")

func TestSub(t *testing.T) {
	got := exercises.Sub(5, 3)
	want := 2
	if( got != want) {
		t.Errorf("Add(2, 3) = %d; want %d", got, want)
	}
	
}