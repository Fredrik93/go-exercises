package test

import (
	"projects/exercises"
	"reflect"
	"testing"
)


func TestShuffle1(t *testing.T) {
    got := exercises.Shuffle([]int{2, 5, 1, 3, 4, 7}, 3)
    want := []int{2, 3, 5, 4, 1, 7}
    if !reflect.DeepEqual(got, want) {
        t.Fatalf("want %v, got %v", want, got)
    }
}
func TestShuffle2(t *testing.T) {
    got := exercises.Shuffle([]int{1,2,3,4,4,3,2,1}, 4)
    want := []int{1,4,2,3,3,2,4,1}
    if !reflect.DeepEqual(got, want) {
        t.Fatalf("want %v, got %v", want, got)
    }
}
func TestShuffle3(t *testing.T) {
    got := exercises.Shuffle([]int{1,1,2,2}, 2)
    want := []int{1,2,1,2}
    if !reflect.DeepEqual(got, want) {
        t.Fatalf("want %v, got %v", want, got)
    }
}