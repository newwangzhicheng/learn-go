package main

import "testing"

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
