package test

import (
	"projects/exercises"
	"testing"
)

func TestAdd(t *testing.T) {
	got := exercises.Add(2, 3)
	want := 5
	if( got != want) {
		t.Errorf("Add(2, 3) = %d; want %d", got, want)
	}
	
}