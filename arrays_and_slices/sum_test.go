package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("sum all elements", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}
		checkSums(t, got, want)
	})

	t.Run("sum all elements(tails) apart from first one(head)", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9, 2})
		want := []int{5, 11}
		checkSums(t, got, want)
	})
	t.Run("sum all tails but one is empty", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{})
		want := []int{5, 0}
		checkSums(t, got, want)
	})
}
