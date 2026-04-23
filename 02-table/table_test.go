package table

import (
	"fmt"
	"testing"
)

// testCases holds a list of test scenarios for the Sum function.
// Each case contains:
// - a: first input number
// - b: second input number
// - want: expected result
var testCases = []struct {
	a    int
	b    int
	want int
}{
	{2, 2, 4},
	{5, 3, 8},
	{8, 4, 12},
	{12, 5, 17},
}

// TestSum tests the Sum function using a simple table-driven approach.
// It loops through all test cases and checks whether the actual result
// matches the expected value.
func TestSum(t *testing.T) {
	for _, row := range testCases {
		got := Sum(row.a, row.b)

		if got != row.want {
			t.Errorf("Sum(%d, %d) = %d, want %d", row.a, row.b, got, row.want)
		}
	}
}

// TestSumSubtest tests the Sum function using subtests.
// Subtests are useful because they give each case its own name,
// which makes failures easier to read and debug.
func TestSumSubtest(t *testing.T) {
	for _, row := range testCases {
		// Rebind row to avoid issues with loop variable capture in subtests.
		row := row

		// Create a readable name for each subtest.
		testName := fmt.Sprintf("%d+%d", row.a, row.b)

		t.Run(testName, func(t *testing.T) {
			got := Sum(row.a, row.b)

			if got != row.want {
				t.Errorf("Sum(%d, %d) = %d, want %d", row.a, row.b, got, row.want)
			}
		})
	}
}
