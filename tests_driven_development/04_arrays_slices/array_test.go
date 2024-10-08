package main

import (
	"slices"
	"testing"
)

// test for sum of elements of array
func TestSum(t *testing.T) {
	t.Run("sum of an array with variable length", func(t *testing.T) {
		// slice
		nums := []int{1, 2, 3}

		result := Sum(nums)
		expected := 6

		if result != expected {
			t.Errorf("your result is %d, but we expect %d with array: %v", result, expected, nums)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("your result is %d, but we expect %d", got, want)
	}
	// if !reflect.DeepEqual(got, want) {
	// 	t.Errorf("your result is %d, but we expect %d", got, want)
	// }
}

// test to check the sum of all tails of slices
func TestSumAllTails(t *testing.T) {

	// assigning a function to a variable.
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("calulate sums of the slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})
	t.Run("edge case - empty slice as input", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})

}
