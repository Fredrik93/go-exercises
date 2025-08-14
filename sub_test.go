package math

import "testing"

func TestSub(t *testing.T) {
	got := Sub(5, 3)
	want := 2
	if( got != want) {
		t.Errorf("Add(2, 3) = %d; want %d", got, want)
	}
	
}