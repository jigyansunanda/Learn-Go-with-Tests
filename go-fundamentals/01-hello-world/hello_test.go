package helloworld

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	totalTests, passedTests := 0, 0

	// ---------------------- define runTest ----------------------
	runTest := func(name string, checker func(t *testing.T) bool) {
		totalTests++
		t.Run(name, func(t *testing.T) {
			if checker(t) {
				passedTests++
			}
		})
	}

	// ---------------------- run tests ---------------------------
	runTest("Empty string defaults to 'world'", func(t *testing.T) bool {
		got := Hello("")
		want := "Hello, World"
		if got != want {
			t.Errorf("got %q want %q", got, want)
			return false
		}
		return true
	})

	runTest("Saying hello to people", func(t *testing.T) bool {
		got := Hello("Abed")
		want := "Hello, Abed"
		if got != want {
			t.Errorf("got %q want %q", got, want)
			return false
		}
		return true
	})

	t.Cleanup(func() {
		fmt.Printf("✅ %d/%d test cases passed\n", passedTests, totalTests)
	})
}
