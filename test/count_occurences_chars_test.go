package test

import (
	"projects/exercises"
	"reflect"
	"testing"
)

func TestOccurence(t *testing.T) {
	got := exercises.NumJewelsInStones("aA","aAAbbCC")
	want := 3

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot %v\nwant:%v", got, want)
	}

}
