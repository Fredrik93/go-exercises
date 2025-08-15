package exercises


import "testing"

func Test10DivisibleBy3(t *testing.T) {
	got := DifferenceOfSum(10,3)
	want := 19
	if( got != want) {
		t.Errorf("Score of string  = %d; want %d", got, want)
	}
}
func Test5DivisibleBy6(t *testing.T) {
	got := DifferenceOfSum(5,6)
	want := 15
	if( got != want) {
		t.Errorf("Score of string  = %d; want %d", got, want)
	}
}