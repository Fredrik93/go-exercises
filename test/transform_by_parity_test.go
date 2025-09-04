package test

import (
	"projects/exercises"
	"reflect"
	"testing"
)

func TestTransformOk(t *testing.T) {
	got := exercises.TransformArray( [] int {4,3,2,1})
	want := [] int {0,0,1,1}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("\ngot %v\nwant:%v", got, want)
	}

}