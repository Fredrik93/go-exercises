package test

import (
	"projects/exercises"
	"reflect"
	"testing"
)

func TestGoodPairs1(t *testing.T) {
	got := exercises.NumOfGoodPairs( [] int {1,2,3,1,1,3})
	want := 4

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot %v\nwant:%v", got, want)
	}

}
func TestGoodPairs2(t *testing.T) {
	got := exercises.NumOfGoodPairs( [] int {1,1,1,1})
	want := 6

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot %v\nwant:%v", got, want)
	}

}
func TestGoodPairs3(t *testing.T) {
	got := exercises.NumOfGoodPairs( [] int {1,2,3})
	want := 0

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot %v\nwant:%v", got, want)
	}

}