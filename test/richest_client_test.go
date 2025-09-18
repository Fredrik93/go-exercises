package test

import (
	"projects/exercises"
	"reflect"
	"testing"
)


func TestRichestClient(t *testing.T) {
    got := exercises.MaximumWealth([][]int{{1,2,3},{3,2,1}});
    want := 6
    if !reflect.DeepEqual(got, want) {
        t.Fatalf("want %v, got %v", want, got)
    }
}
func TestRichestClient2(t *testing.T) {
    got := exercises.MaximumWealth([][]int{{1,5},{7,3},{3,5}});
    want := 10
    if !reflect.DeepEqual(got, want) {
        t.Fatalf("want %v, got %v", want, got)
    }
}

func TestRichestClient3(t *testing.T) {
    got := exercises.MaximumWealth([][]int{{2,8,7},{7,1,3},{1,9,5}});
    want := 17
    if !reflect.DeepEqual(got, want) {
        t.Fatalf("want %v, got %v", want, got)
    }
}