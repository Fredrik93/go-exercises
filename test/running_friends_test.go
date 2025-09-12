package test

import (
	"projects/exercises"
	"reflect"
	"testing"
)

func TestRunningFriends(t *testing.T) {
	got := exercises.RecoverOrder( [] int {3,1,2,5,4}, [] int {1,3,4})
	want := [] int {3,1,4}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot %v\nwant:%v", got, want)
	}

}
func TestRunningFriends2(t *testing.T) {
	got := exercises.RecoverOrder( [] int {1,4,5,3,2}, [] int {5,2})
	want := [] int {5,2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot %v\nwant:%v", got, want)
	}

}