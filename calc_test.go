package math 

import "testing"

func TestAddFive(t *testing.T) {
	got := AddFive(3)
	want := 8
	if got != want {
		t.Errorf("AddFive(3) = %d; want %d", got, want)
	}
}