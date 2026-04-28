package subtest

import (
	"testing"
)

// TestSubtests demonstrates the use of subtests to organize test cases.
// Each subtest runs independently and can be run in isolation.
func TestSubtests(t *testing.T) {
	// Setup code can be placed here if needed.
	t.Run("A", func(t *testing.T) {
		t.Log("Test A completed")
	})

	t.Run("B", func(t *testing.T) {
		t.Log("Test B completed")
	})

	t.Run("C", func(t *testing.T) {
		t.Log("Test C completed")
	})
	// Teardown code can be placed here if needed.
}
