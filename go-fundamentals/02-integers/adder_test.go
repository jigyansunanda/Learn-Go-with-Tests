package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	totalTests, passedTests := 0, 0

	// ---------------------- define runTest ----------------------
	runTest := func(name string, tester func(t *testing.T) bool) {
		totalTests++
		t.Run(name, func(t *testing.T) {
			if tester(t) {
				passedTests++
			}
		})
	}

	// ---------------------- run tests ---------------------------
	runTest("Testing with 1 element", func(t *testing.T) bool {
		nums := []int{9807}
		got := Add(nums...)
		want := 9807
		if got != want {
			t.Errorf("got %d want %d", got, want)
			return false
		}
		return true
	})

	runTest("Testing with 0 element", func(t *testing.T) bool {
		nums := []int{}
		got := Add(nums...)
		want := 0
		if got != want {
			t.Errorf("got %d want %d", got, want)
			return false
		}
		return true
	})

	runTest("Testing with 5 element", func(t *testing.T) bool {
		nums := []int{12, 23, 8, 97, 30, -5}
		got := Add(nums...)
		want := 165
		if got != want {
			t.Errorf("got %d want %d", got, want)
			return false
		}
		return true
	})

	// ----------------------- Clean up ----------------------------
	t.Cleanup(func() {
		fmt.Printf("✅ %d/%d test cases passed\n", passedTests, totalTests)
	})
}
