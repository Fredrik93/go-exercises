package test

import (
	"projects/exercises"
	"reflect"
	"testing"
)

func TestConcatenate123(t *testing.T) {
	got := exercises.GetConcatenation([]int{1, 2, 3})
	want := []int{1, 2, 3, 1, 2, 3}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot %v\nwant:%v", got, want)
	}

}
