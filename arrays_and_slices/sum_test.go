package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15
	if got != want {
		t.Errorf("Got %d, Want %d", got, want)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{3, 4, 5})
	want := []int{6, 12}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, Want %v", got, want)
	}
}

func TestSumAllTail(t *testing.T) {
	got := SumAllTail([]int{1, 2, 3}, []int{3, 4, 5}, []int{})
	want := []int{5, 9, 0}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, Want %v", got, want)
	}
}
