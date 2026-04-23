package normal

import "testing"

// TestSum demonstrates basic unit testing with arrange/act/assert pattern.
func TestSum(t *testing.T) {
	// Arrange: Set up test data
	want := 8
	input1 := 3
	input2 := 5

	// Act: Execute the function under test
	got := Sum(input1, input2)

	// Assert: Verify the result
	if got != want {
		t.Errorf("Test failed! Expected: %d, Got: %d", want, got)
	}
}
