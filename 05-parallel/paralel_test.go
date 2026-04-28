package parallel

import (
	"fmt"
	"testing"
)

// testCases holds a list of test scenarios for the Sum function.
var testCases = []struct {
	x    int
	y    int
	want int
}{
	{2, 2, 4},
	{5, 3, 8},
	{8, 4, 12},
	{12, 5, 17},
}

// TestSumParallel tests the Sum function with parallel execution.
// Each subtest runs in parallel to demonstrate concurrent test execution.
func TestSumParallel(t *testing.T) {
	t.Parallel()
	for _, tc := range testCases {
		// Rebind tc to avoid issues with loop variable capture in parallel tests.
		tc := tc
		testName := fmt.Sprintf("%d+%d", tc.x, tc.y)
		t.Run(testName, func(t *testing.T) {
			t.Parallel()
			got, _ := Sum(tc.x, tc.y)
			if got != tc.want {
				t.Errorf("Sum(%d, %d) = %d, want %d", tc.x, tc.y, got, tc.want)
			}
		})
	}
}
