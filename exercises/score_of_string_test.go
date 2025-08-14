package exercises


import "testing"

func TestZaz(t *testing.T) {
	got := scoreOfString("zaz")
	want := 50
	if( got != want) {
		t.Errorf("Score of string  = %d; want %d", got, want)
	}
	
}