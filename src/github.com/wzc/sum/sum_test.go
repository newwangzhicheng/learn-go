package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	assetsString := func(want int, got int, numbers []int) {
		t.Helper()

		if got != want {
			t.Errorf("want %d but got %d, %v", want, got, numbers)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 3, 4, 5, 6}

		got := Sum(numbers)
		want := 19

		assetsString(want, got, numbers)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{2, 3})
	want := []int{3, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("want %v but got %v", want, got)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSum := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v but got %v", want, got)
		}
	}

	t.Run("make sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{2, 3, 4})
		want := []int{5, 7}
		checkSum(t, got, want)
	})

	t.Run("make sum of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{2, 3, 4})
		want := []int{0, 7}
		checkSum(t, got, want)
	})

}
